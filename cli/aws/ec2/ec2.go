package ec2

import (
	"fmt"
	"github.com/dynport/dgtk/cli"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/ec2"
	"github.com/dynport/gocloud/aws/pricing"
	"log"
	"sort"
	"strings"
	"time"
)

const HOURS_PER_MONTH = 365 * 24.0 / 12.0

func Register(router *cli.Router) {
	router.RegisterFunc("aws/ec2/instances/describe", DescribeInstances, "Describe ec2 instances")
	router.Register("aws/ec2/instances/run", &RunInstances{}, "Run ec2 instances")
	router.Register("aws/ec2/images/create", &CreateImage{}, "Create image from instance")
	router.Register("aws/ec2/instances/terminate", &TerminateInstances{}, "Terminate ec2 instances")
	router.Register("aws/ec2/tags/create", &CreateTags{}, "Create Tags")
	router.RegisterFunc("aws/ec2/tags/describe", DescribeTags, "Describe Tags")
	router.Register("aws/ec2/images/describe", &DescribeImages{}, "Describe ec2 Images")
	router.RegisterFunc("aws/ec2/key-pairs/describe", DescribeKeyPairs, "Describe key pairs")
	router.RegisterFunc("aws/ec2/addresses/describe", DescribeAddresses, "Describe Addresses")
	router.RegisterFunc("aws/ec2/security-groups/describe", DescribeSecurityGroups, "Describe Security Groups")
	router.RegisterFunc("aws/ec2/spot-price-history/describe", DescribeSpotPriceHistory, "Describe Spot Price History")
	router.Register("aws/ec2/prices", &Prices{}, "EC2 Prices")
}

func client() *ec2.Client {
	return ec2.NewFromEnv()
}

type Prices struct {
	Region string `cli:"type=opt short=r default=eu-ireland"`
	Heavy  bool   `cli:"type=opt long=heavy"`
}

func (a *Prices) Run() error {
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
func monthlyPrice(price float64) string {
	return fmt.Sprintf("%.02f", price)
}

type DescribeImages struct {
	Canonical bool `cli:"type=opt long=canonical"`
	Self      bool `cli:"type=opt long=self"`
	Ubuntu    bool `cli:"type=opt long=ubuntu"`
	Raring    bool `cli:"type=opt long=raring"`
	Saucy     bool `cli:"type=opt long=saucy"`
}

func (a *DescribeImages) Run() error {
	log.Println("describing images")
	filter := &ec2.ImageFilter{}
	if a.Canonical {
		filter.Owner = ec2.CANONICAL_OWNER_ID
	} else if a.Self {
		filter.Owner = ec2.SELF_OWNER_ID
	}

	if a.Ubuntu {
		filter.Name = ec2.UBUNTU_PREFIX
	} else if a.Raring {
		filter.Name = ec2.UBUNTU_RARING_PREFIX
	} else if a.Saucy {
		filter.Name = ec2.UBUNTU_SAUCY_PREFIX
	}

	images, e := client().DescribeImagesWithFilter(filter)
	if e != nil {
		return e
	}
	sort.Sort(images)
	table := gocli.NewTable()
	for _, image := range images {
		table.Add(image.ImageId, image.Name, image.ImageState)
	}
	fmt.Println(table)
	return nil
}

func DescribeSpotPriceHistory() error {
	filter := &ec2.SpotPriceFilter{
		InstanceTypes:       []string{"c1.medium"},
		ProductDescriptions: []string{ec2.DESC_LINUX_UNIX},
		StartTime:           time.Now().Add(-7 * 24 * time.Hour),
	}
	prices, e := client().DescribeSpotPriceHistory(filter)
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, price := range prices {
		table.Add(price.InstanceType, price.ProductDescription, price.SpotPrice, price.Timestamp, price.AvailabilityZone)
	}
	fmt.Println(table)
	return nil
}

func DescribeSecurityGroups() error {
	groups, e := client().DescribeSecurityGroups()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, group := range groups {
		table.Add(group.GroupId, "", "", group.GroupName)
		for _, perm := range group.IpPermissions {
			ports := ""
			if perm.FromPort > 0 {
				if perm.FromPort == perm.ToPort {
					ports += fmt.Sprintf("%d", perm.FromPort)
				} else {
					ports += fmt.Sprintf("%d-%d", perm.FromPort, perm.ToPort)
				}
			}
			groups := []string{}
			for _, group := range perm.Groups {
				groups = append(groups, group.GroupId)
			}
			if len(groups) > 0 {
				table.Add("", perm.IpProtocol, ports, strings.Join(groups, ","))
			}

			if len(perm.IpRanges) > 0 {
				table.Add("", perm.IpProtocol, ports, strings.Join(perm.IpRanges, ","))
			}
		}
	}
	fmt.Print(table)
	return nil
}

func DescribeAddresses() error {
	addresses, e := client().DescribeAddresses()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, address := range addresses {
		table.Add(address.PublicIp, address.PrivateIpAddress, address.Domain, address.InstanceId)
	}
	fmt.Print(table)
	return nil
}

func DescribeKeyPairs() error {
	pairs, e := client().DescribeKeyPairs()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, pair := range pairs {
		table.Add(pair.KeyName, pair.KeyFingerprint)
	}
	fmt.Print(table)
	return nil
}

type CreateImage struct {
	ImageId string `cli:"type=arg required=true"`
}

func (a *CreateImage) Run() error {
	log.Printf("creating image of instance %s", a.ImageId)
	return fmt.Errorf("implement me")
}

type RunInstances struct {
	InstanceType     string `cli:"type=opt short=t desc='Instance Type' required=true"`
	ImageId          string `cli:"type=opt short=i desc='Image Id' required=true"`
	KeyName          string `cli:"type=opt short=k desc='SSH Key' required=true"`
	SecurityGroup    string `cli:"type=opt short=g desc='Security Group'"`
	SubnetId         string `cli:"type=opt long=subnet-id desc='Subnet Id'"`
	PublicIp         bool   `cli:"type=opt long=public-ip desc='Assign Public IP'"`
	AvailabilityZone string `cli:"type=opt long=availability-zone desc='Availability Zone'"`
	Name             string `cli:"type=opt long=name desc='Name of the Instanze'"`
}

func (a *RunInstances) Run() error {
	config := &ec2.RunInstancesConfig{
		ImageId:          a.ImageId,
		KeyName:          a.KeyName,
		InstanceType:     a.InstanceType,
		AvailabilityZone: a.AvailabilityZone,
		SubnetId:         a.SubnetId,
	}
	if a.PublicIp {
		nic := &ec2.CreateNetworkInterface{
			DeviceIndex: 0, AssociatePublicIpAddress: true, SubnetId: a.SubnetId,
		}
		if a.SecurityGroup != "" {
			nic.SecurityGroupIds = []string{a.SecurityGroup}
		}
		config.NetworkInterfaces = []*ec2.CreateNetworkInterface{nic}
	} else {
		if a.SecurityGroup != "" {
			config.SecurityGroups = []string{a.SecurityGroup}
		}
	}
	list, e := client().RunInstances(config)
	ids := []string{}
	for _, i := range list {
		ids = append(ids, i.InstanceId)
	}
	if a.Name != "" {
		log.Printf("tagging %v with %q", ids, a.Name)
		e := client().CreateTags(ids, map[string]string{"Name": a.Name})
		if e != nil {
			log.Printf("ERROR: " + e.Error())
		}
	}
	log.Printf("started instances %v", ids)
	return e
}

type TerminateInstances struct {
	InstanceIds []string `cli:"type=arg required=true"`
}

func (a *TerminateInstances) Run() error {
	return client().TerminateInstances(a.InstanceIds)
}

type CreateTags struct {
	ResourceId string `cli:"type=arg required=true"`
	Key        string `cli:"type=arg required=true"`
	Value      string `cli:"type=arg required=true"`
}

func (r *CreateTags) Run() error {
	tags := map[string]string{
		r.Key: r.Value,
	}
	return client().CreateTags([]string{r.ResourceId}, tags)
}

func DescribeTags() error {
	tags, e := client().DescribeTags()
	if e != nil {
		return e
	}
	sort.Sort(tags)
	table := gocli.NewTable()
	for _, tag := range tags {
		table.Add(tag.ResourceType, tag.ResourceId, tag.Key, tag.Value)
	}
	fmt.Println(table)
	return nil
}

func DescribeInstances() error {
	log.Print("describing ec2 instances")
	instances, e := client().DescribeInstances()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("id", "image", "name", "state", "type", "private_ip", "ip", "az", "launched")
	for _, i := range instances {
		sgs := []string{}
		if i.InstanceStateName != "running" {
			continue
		}
		for _, g := range i.SecurityGroups {
			sgs = append(sgs, g.GroupId)
		}
		table.Add(
			i.InstanceId,
			i.ImageId,
			i.Name(),
			i.InstanceStateName,
			i.InstanceType,
			i.PrivateIpAddress,
			i.IpAddress,
			i.PlacementAvailabilityZone,
			i.LaunchTime.Format("2006-01-02T15:04"),
		)
	}
	fmt.Println(table)

	return nil
}
