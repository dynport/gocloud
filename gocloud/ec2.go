package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/ec2"
	"sort"
	"strings"
	"time"
)

const (
	CLI_INSTANCE_TYPE  = "--instance-type"
	CLI_SSH_KEY        = "--ssh-key"
	CLI_SECURITY_GROUP = "--security-group"
	CLI_CANONICAL      = "--canonical"
	CLI_SELF           = "--self"
	CLI_UBUNTU         = "--ubuntu"
	CLI_UBUNTU_RARING  = "--raring"
	CLI_UBUNTU_SAUCY   = "--saucy"

	USAGE_IMAGE_ID            = "IMAGE"
	USAGE_KEY_NAME            = "KEY_NAME"
	USAGE_IMAGE_TYPE          = "IMAGE_TYPE"
	USAGE_TERMINATE_INSTANCES = "INSTANCE [INSTANCE [...]]"
	USAGE_CREATE_TAGS         = "RESOURCE KEY VALUE"

	USAGE_RUN_INSTANCE = "IMAGE"
)

func ec2Client() *ec2.Client {
	return ec2.NewFromEnv()
}

func init() {
	router.RegisterFunc("aws/ec2/instances/describe", ec2DescribeInstances, "Describe ec2 instances")
	router.Register("aws/ec2/instances/run", &ec2RunInstances{}, "Run ec2 instances")
	router.Register("aws/ec2/images/create", &ec2CreateImage{}, "Create image from instance")
	router.Register("aws/ec2/instances/terminate", &ec2TerminateInstances{}, "Terminate ec2 instances")
	router.Register("aws/ec2/tags/create", &ec2CreateTags{}, "Create Tags")
	router.RegisterFunc("aws/ec2/tags/describe", ec2DescribeTags, "Describe Tags")
	router.Register("aws/ec2/images/describe", &ec2DescribeImages{}, "Describe ec2 Images")

	router.RegisterFunc("aws/ec2/key-pairs/describe", ec2DescribeKeyPairs, "Describe key pairs")
	router.RegisterFunc("aws/ec2/addresses/describe", ec2DescribeAddresses, "Describe Addresses")
	router.RegisterFunc("aws/ec2/security-groups/describe", ec2DescribeSecurityGroups, "Describe Security Groups")
	router.RegisterFunc("aws/ec2/spot-price-history/describe", ec2DescribeSpotPriceHistory, "Describe Spot Price History")
}

func ec2DescribeSpotPriceHistory() error {
	filter := &ec2.SpotPriceFilter{
		InstanceTypes:       []string{"c1.medium"},
		ProductDescriptions: []string{ec2.DESC_LINUX_UNIX},
		StartTime:           time.Now().Add(-7 * 24 * time.Hour),
	}
	prices, e := ec2Client().DescribeSpotPriceHistory(filter)
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

func ec2DescribeSecurityGroups() error {
	groups, e := ec2Client().DescribeSecurityGroups()
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

func ec2DescribeAddresses() error {
	addresses, e := ec2Client().DescribeAddresses()
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

func ec2DescribeKeyPairs() error {
	pairs, e := ec2Client().DescribeKeyPairs()
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

const USAGE_CREATE_IMAGE = "INSTANCE"

type ec2CreateImage struct {
	ImageId string `cli:"type=arg required=true"`
}

func (a *ec2CreateImage) Run() error {
	logger.Infof("creating image of instance %s", a.ImageId)
	return fmt.Errorf("implement me")
}

type ec2RunInstances struct {
	InstanceType  string `cli:"type=opt short=t desc='Instance Type' required=true"`
	ImageId       string `cli:"type=opt short=i desc='Image Id' required=true"`
	KeyName       string `cli:"type=opt short=k desc='SSH Key' required=true"`
	SecurityGroup string `cli:"type=opt short=g desc='Security Group' required=true"`
}

func (a *ec2RunInstances) Run() error {
	config := &ec2.RunInstancesConfig{
		ImageId:        a.ImageId,
		KeyName:        a.KeyName,
		InstanceType:   a.ImageId,
		SecurityGroups: []string{a.SecurityGroup},
	}
	_, e := ec2Client().RunInstances(config)
	return e
}

type ec2TerminateInstances struct {
	InstanceIds []string `cli:"type=arg required=true"`
}

func (a *ec2TerminateInstances) Run() error {
	return ec2Client().TerminateInstances(a.InstanceIds)
}

type ec2CreateTags struct {
	ResourceId string `cli:"type=arg required=true"`
	Key        string `cli:"type=arg required=true"`
	Value      string `cli:"type=arg required=true"`
}

func (r *ec2CreateTags) Run() error {
	tags := map[string]string{
		r.Key: r.Value,
	}
	return ec2Client().CreateTags([]string{r.ResourceId}, tags)
}

func ec2DescribeTags() error {
	tags, e := ec2Client().DescribeTags()
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

func ec2DescribeInstances() error {
	logger.Info("describing ec2 instances")
	instances, e := ec2Client().DescribeInstances()
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

type ec2DescribeImages struct {
	Canonical bool `cli:"type=opt long=canonical"`
	Self      bool `cli:"type=opt long=self"`
	Ubuntu    bool `cli:"type=opt long=ubuntu"`
	Raring    bool `cli:"type=opt long=raring"`
	Saucy     bool `cli:"type=opt long=saucy"`
}

func (a *ec2DescribeImages) Run() error {
	logger.Info("describing images")
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

	images, e := ec2Client().DescribeImagesWithFilter(filter)
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
