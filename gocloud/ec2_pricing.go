package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/pricing"
	"sort"
)

const (
	CLI_REGION = "-r"
	CLI_HEAVY  = "--heavy"
)

func init() {
	args := gocli.NewArgs(nil)
	args.RegisterString(CLI_REGION, "region", false, "eu-ireland", "AWS Region")
	args.RegisterBool(CLI_HEAVY, "heavy", false, false, "Use prices for reserved instances, heavy")
	router.Register("aws/ec2/prices", &ec2Prices{}, "EC2 Prices")
}

var regionMapping = map[string]string{
	"eu-ireland": "eu-west-1",
	"eu-west":    "eu-west-1",
	"apac-tokyo": "ap-northeast-1",
	"apac-sin":   "ap-southeast-1",
	"apac-syd":   "ap-southeast-2",
}

func normalizeRegion(raw string) string {
	if v, ok := regionMapping[raw]; ok {
		return v
	}
	return raw
}

type ec2Prices struct {
	Region string `cli:"type=opt short=r default=eu-ireland"`
	Heavy  bool   `cli:"type=opt long=heavy"`
}

func (a *ec2Prices) Run() error {
	configs := pricing.InstanceTypeConfigs
	var pr *pricing.Pricing
	var e error
	regionName := a.Region
	typ := "od"
	if a.Heavy {
		regionName = normalizeRegion(regionName)
		typ = "ri-heavy"
		pr, e = pricing.LinuxReservedHeavy()
	} else {
		pr, e = pricing.LinuxOnDemand()
	}
	if e != nil {
		return e
	}
	logger.Debugf("using region %q", regionName)
	priceMapping := map[string]pricing.PriceList{}
	region := pr.FindRegion(regionName)
	if region == nil {
		return fmt.Errorf("could not find prices for reagion %q. Known regions are %v", regionName, pr.RegionNames())
	}
	for _, t := range region.InstanceTypes {
		for _, size := range t.Sizes {
			priceMapping[size.Size] = size.ValueColumns.Prices()
		}
	}
	table := gocli.NewTable()
	table.Add("Type", "Cores", "ECUs", "GB RAM", "Region", "Type", "$/Hour", "$/Month", "$/Core", "$/GB")
	for _, config := range configs {
		cols := []interface{}{
			config.Name, config.Cpus, config.ECUs, config.Memory,
		}
		if prices, ok := priceMapping[config.Name]; ok {
			cols = append(cols, normalizeRegion(regionName), typ)
			if len(prices) > 0 {
				sort.Sort(prices)
				price := prices[0].TotalPerHour()
				perMonth := price * HOURS_PER_MONTH
				perCore := perMonth / float64(config.Cpus)
				perGb := perMonth / config.Memory
				cols = append(cols, fmt.Sprintf("%.03f", price), monthlyPrice(perMonth), monthlyPrice(perCore), monthlyPrice(perGb))
			}
		}
		table.Add(cols...)
	}
	fmt.Println(table)
	return nil
}

func monthlyPrice(price float64) string {
	return fmt.Sprintf("%.02f", price)
}
