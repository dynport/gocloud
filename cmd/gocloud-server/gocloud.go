package main

import (
	"encoding/xml"

	"github.com/dynport/gocloud/aws/ec2"
)

type DescribeVpcsResponse struct {
	XMLName xml.Name  `xml:"DescribeVpcsResponse"`
	Vpcs    []*Ec2Vpc `xml:"vpcSet>item"`
}

type Ec2Vpc struct {
	VpcId           string     `xml:"vpcId"`
	State           string     `xml:"state"`
	CidrBlock       string     `xml:"cidrBlock"`
	DhcpOptions     string     `xml:"dhcpOptions"`
	InstanceTenancy string     `xml:"instanceTenancy"`
	IsDefault       string     `xml:"isDefault"`
	Tags            []*ec2.Tag `xml:"tagSet>item"`
}
