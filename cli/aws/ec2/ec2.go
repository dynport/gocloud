package ec2

import (
	"fmt"
	"github.com/dynport/dgtk/cli"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/ec2"
	"log"
	"sort"
	"strings"
	"time"
)

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
}

func client() *ec2.Client {
	return ec2.NewFromEnv()
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
	InstanceType  string `cli:"type=opt short=t desc='Instance Type' required=true"`
	ImageId       string `cli:"type=opt short=i desc='Image Id' required=true"`
	KeyName       string `cli:"type=opt short=k desc='SSH Key' required=true"`
	SecurityGroup string `cli:"type=opt short=g desc='Security Group' required=true"`
}

func (a *RunInstances) Run() error {
	config := &ec2.RunInstancesConfig{
		ImageId:        a.ImageId,
		KeyName:        a.KeyName,
		InstanceType:   a.ImageId,
		SecurityGroups: []string{a.SecurityGroup},
	}
	_, e := client().RunInstances(config)
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
	table.Add("id", "image", "name", "key", "state", "type", "private_ip", "ip", "launched", "groups")
	for _, i := range instances {
		sgs := []string{}
		for _, g := range i.SecurityGroups {
			sgs = append(sgs, g.GroupId)
		}
		table.Add(
			i.InstanceId,
			i.ImageId,
			i.Name(),
			i.KeyName,
			i.InstanceStateName,
			i.InstanceType,
			i.PrivateIpAddress,
			i.IpAddress,
			i.LaunchTime.Format("2006-01-02T15:04:05"),
			strings.Join(sgs, ","),
		)
	}
	fmt.Println(table)

	return nil
}
