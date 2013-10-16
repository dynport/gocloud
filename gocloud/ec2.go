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
	CLI_INSTANCE_TYPE         = "--instance-type"
	CLI_SSH_KEY               = "--ssh-key"
	CLI_SECURITY_GROUP        = "--security-group"
	USAGE_IMAGE_ID            = "IMAGE"
	USAGE_KEY_NAME            = "KEY_NAME"
	USAGE_IMAGE_TYPE          = "IMAGE_TYPE"
	USAGE_TERMINATE_INSTANCES = "INSTANCE [INSTANCE [...]]"
	USAGE_CREATE_TAGS         = "RESOURCE KEY VALUE"

	USAGE_RUN_INSTANCE = "IMAGE"
)

var ec2Client = ec2.NewFromEnv()

func init() {
	router.Register("aws/ec2/instances/describe", &gocli.Action{
		Handler: ec2DescribeInstances, Description: "Describe ec2 instances",
	})

	args := gocli.NewArgs(nil)
	args.RegisterString(CLI_INSTANCE_TYPE, "instance_type", false, "c1.medium", "Instance Type")
	args.RegisterString(CLI_SSH_KEY, "ssh_key", true, "", "SSH Key")
	args.RegisterString(CLI_SECURITY_GROUP, "security_group", true, "", "Security Group")

	router.Register("aws/ec2/instances/run", &gocli.Action{
		Handler: ec2RunInstances, Description: "Run ec2 instances", Args: args,
	})

	router.Register("aws/ec2/images/create", &gocli.Action{
		Handler: ec2CreateImage, Description: "Create image from instance",
	})

	router.Register("aws/ec2/instances/terminate", &gocli.Action{
		Handler: ec2TerminateInstances, Description: "Terminate ec2 instances", Usage: USAGE_TERMINATE_INSTANCES,
	})

	router.Register("aws/ec2/tags/create", &gocli.Action{
		Handler: ec2CreateTags, Description: "Create Tags", Usage: USAGE_CREATE_TAGS,
	})

	router.Register("aws/ec2/tags/describe", &gocli.Action{
		Handler: ec2DescribeTags, Description: "Describe Tags",
	})

	router.Register("aws/ec2/images/describe", &gocli.Action{
		Handler: ec2DescribeImages, Description: "Describe ec2 images",
	})

	router.Register("aws/ec2/key-pairs/describe", &gocli.Action{
		Handler: ec2DescribeKeyPairs, Description: "Describe key pairs",
	})

	router.Register("aws/ec2/addresses/describe", &gocli.Action{
		Handler: ec2DescribeAddresses, Description: "Describe Addresses",
	})

	router.Register("aws/ec2/security-groups/describe", &gocli.Action{
		Handler: ec2DescribeSecurityGroups, Description: "Describe Security Groups",
	})

	router.Register("aws/ec2/spot-price-history/describe", &gocli.Action{
		Handler: ec2DescribeSpotPriceHistory, Description: "Describe Spot Price History",
	})
}

func ec2DescribeSpotPriceHistory(args *gocli.Args) error {
	filter := &ec2.SpotPriceFilter{
		InstanceTypes:       []string{"c1.medium"},
		ProductDescriptions: []string{ec2.DESC_LINUX_UNIX},
		StartTime:           time.Now().Add(-7 * 24 * time.Hour),
	}
	prices, e := ec2Client.DescribeSpotPriceHistory(filter)
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

func ec2DescribeSecurityGroups(args *gocli.Args) error {
	groups, e := ec2Client.DescribeSecurityGroups()
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

func ec2DescribeAddresses(args *gocli.Args) error {
	addresses, e := ec2Client.DescribeAddresses()
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

func ec2DescribeKeyPairs(args *gocli.Args) error {
	pairs, e := ec2Client.DescribeKeyPairs()
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

func ec2CreateImage(args *gocli.Args) error {
	if len(args.Args) != 0 {
		return fmt.Errorf(USAGE_CREATE_IMAGE)
	}
	instanceId := args.Args[0]
	logger.Infof("creating image of instance %s", instanceId)
	return nil
}

func ec2RunInstances(args *gocli.Args) error {
	if len(args.Args) == 0 {
		return fmt.Errorf(USAGE_IMAGE_ID)
	}
	imageId := args.Args[0]
	imageType := args.MustGetString(CLI_INSTANCE_TYPE)
	keyName := args.MustGetString(CLI_SSH_KEY)

	config := &ec2.RunInstancesConfig{
		ImageId:      imageId,
		KeyName:      keyName,
		InstanceType: imageType,
	}
	_, e := ec2Client.RunInstances(config)
	return e
}

func ec2TerminateInstances(args *gocli.Args) error {
	if len(args.Args) < 1 {
		return fmt.Errorf(USAGE_TERMINATE_INSTANCES)
	}
	return ec2Client.TerminateInstances(args.Args)
}

func ec2CreateTags(args *gocli.Args) error {
	if len(args.Args) != 3 {
		return fmt.Errorf(USAGE_CREATE_TAGS)
	}
	resourceId := args.Args[0]
	tags := map[string]string{
		args.Args[1]: args.Args[2],
	}
	return ec2Client.CreateTags([]string{resourceId}, tags)
}

func ec2DescribeTags(args *gocli.Args) error {
	if len(args.Args) < 0 {
		return fmt.Errorf("instance...")
	}
	tags, e := ec2Client.DescribeTags()
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

func ec2DescribeInstances(args *gocli.Args) error {
	logger.Info("describing ec2 instances")
	instances, e := ec2Client.DescribeInstances()
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

func ec2DescribeImages(args *gocli.Args) error {
	logger.Info("describing images")
	images, e := ec2Client.DescribeImagesWithFilter(&ec2.ImageFilter{Owner: ec2.CANONICAL_OWNER_ID, Name: ec2.UBUNTU_RARING_PREFIX})
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
