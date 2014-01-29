package pricing

import (
	"encoding/json"
	"strconv"
)

type Pricing struct {
	Version float64        `json:"vers"`
	Config  *PricingConfig `json:"config"`
}

func (pricing *Pricing) FindRegion(region string) *Region {
	for _, r := range pricing.Config.Regions {
		if r.Region == region {
			return r
		}
	}
	return nil
}

func (pricing *Pricing) RegionNames() []string {
	regions := []string{}
	for _, r := range pricing.Config.Regions {
		regions = append(regions, r.Region)
	}
	return regions
}

type PricingConfig struct {
	Rate         string    `json:"rate"`
	ValueColumns []string  `json:"valueColumns"`
	Currencies   []string  `json:"currencies"`
	Regions      []*Region `json:"regions"`
}

type Region struct {
	Region        string          `json:"region"`
	InstanceTypes []*InstanceType `json:"instanceTypes"`
	Types         []*Type         `json:"types"`
}

type Type struct {
	Values []*Value `json:"values"`
}

type Value struct {
	Prices Prices `json:"prices"`
	Rate   string `json:"rate"`
}

type Prices map[string]string

func (prices Prices) USD() (float64, bool) {
	p, ok := prices["USD"]
	if !ok {
		return 0, ok
	}
	price, e := strconv.ParseFloat(p, 64)
	if e != nil {
		return 0, false
	}
	return price, true
}

type InstanceType struct {
	Type  string  `json:"type"`
	Sizes []*Size `json:"sizes"`
}

type Size struct {
	Size         string       `json:"size"`
	ValueColumns ValueColumns `json:"valueColumns"`
}

func LoadPricing(b []byte) (p *Pricing, e error) {
	pricing := &Pricing{}
	e = json.Unmarshal(b, pricing)
	return pricing, e
}

func LinuxOnDemand() (p *Pricing, e error) {
	return loadPricesFor("linux-od.json")
}

func LinuxReservedHeavy() (p *Pricing, e error) {
	return loadPricesFor("linux-ri-heavy.json")
}

func loadPricesFor(t string) (p *Pricing, e error) {
	b, e := readAsset(t)
	if e != nil {
		return nil, e
	}
	return LoadPricing(b)
}

type InstanceTypeConfigs []*InstanceTypeConfig

func (config InstanceTypeConfigs) Len() int {
	return len(config)
}

func (config InstanceTypeConfigs) Swap(a, b int) {
	config[a], config[b] = config[b], config[a]
}

var sortOrder = map[string]int{
	"General purpose":   1,
	"Compute optimized": 2,
	"GPU instances":     3,
	"Memory optimized":  4,
	"Storage optimized": 5,
	"Micro Instances":   6,
}

func (config InstanceTypeConfigs) Less(a, b int) bool {
	instanceA := config[a]
	instanceB := config[b]
	famA, okA := sortOrder[instanceA.Family]
	famB, okB := sortOrder[instanceB.Family]
	if okA && okB {
		if famA == famB {
			return instanceA.Cpus < instanceB.Cpus
		} else {
			return famA < famB
		}
	} else if okA {
		return true
	} else if okB {
		return false
	} else {
		return instanceA.Cpus < instanceB.Cpus
	}
}

func AllInstanceTypeConfigs() (configs InstanceTypeConfigs, e error) {
	b, e := readAsset("instance_types.json")
	if e != nil {
		return nil, e
	}
	e = json.Unmarshal(b, &configs)
	return configs, e
}
