package main

import (
	"strings"

	"github.com/dynport/gocloud/aws/cloudformation"
	"github.com/dynport/gocloud/aws/ec2"
)

func linkForResource(s string) string {
	switch strings.Split(s, "-")[0] {
	case "vpc":
		return "/vpcs/" + s
	case "sg":
		return "/sgs/" + s
	case "subnet":
		return "/subnets/" + s
	case "i":
		return "/instances/" + s
	default:
		return ""
	}
}

func NewStackResource(raw *cloudformation.StackResource) *StackResource {
	return &StackResource{PhysicalResourceId: raw.PhysicalResourceId, LogicalResourceId: raw.LogicalResourceId, Link: linkForResource(raw.PhysicalResourceId)}
}

func extractTag(tags []*ec2.Tag, key string) string {
	for _, t := range tags {
		if t.Key == key {
			return t.Value
		}
	}
	return ""
}

func NewInstance(raw *ec2.Instance) *Instance {
	groups := []string{}
	for _, g := range raw.SecurityGroups {
		groups = append(groups, g.GroupId)
	}
	return &Instance{
		Id:               raw.InstanceId,
		Name:             raw.Name(),
		SubnetId:         raw.SubnetId,
		SecurityGroups:   groups,
		PrivateIpAddress: raw.PrivateIpAddress,
		Original:         raw,
		Type:             raw.InstanceType,
		AvailabilityZone: raw.PlacementAvailabilityZone,
		PublicIpAddress:  raw.IpAddress,
		State:            raw.InstanceStateName,
		VpcId:            raw.VpcId,
	}
}

func NewVpc(raw *Ec2Vpc) *Vpc {
	return &Vpc{
		Id: raw.VpcId, Name: extractTag(raw.Tags, "Name"),
		State:     raw.State,
		CidrBlock: raw.CidrBlock,
	}
}
