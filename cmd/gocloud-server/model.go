package main

import (
	"strconv"

	"github.com/dynport/gocloud/aws/ec2"
)

type Instance struct {
	Id               string
	Name             string
	Type             string
	SecurityGroups   []string
	PrivateIpAddress string
	SubnetId         string
	PublicIpAddress  string
	State            string
	AvailabilityZone string
	VpcId            string
	Original         interface{}
}

func (i *Instance) HasSecurityGroup(id string) bool {
	for _, g := range i.SecurityGroups {
		if g == id {
			return true
		}
	}
	return false
}

type Image struct {
	Id   string
	Name string
}

type Stack struct {
	Name     string
	Status   string
	Original interface{}
}

type StackResource struct {
	Link               string
	PhysicalResourceId string
	LogicalResourceId  string
}

type SecurityGroup struct {
	Id            string
	Name          string
	Description   string
	VpcId         string
	IpPermissions []*IpPermission
	Original      interface{}
}

type IpPermission struct {
	Protocol string
	Ports    string
	Groups   []string
	IpRanges []string
}

type Vpc struct {
	Id        string
	Name      string
	CidrBlock string
	State     string
}

func NewSecurityGroup(raw *ec2.SecurityGroup) *SecurityGroup {
	g := &SecurityGroup{
		Id: raw.GroupId, Name: raw.GroupName, Description: raw.GroupDescription,
		VpcId: raw.VpcId, Original: raw,
	}
	for _, p := range raw.IpPermissions {
		protocol := p.IpProtocol
		if protocol == "-1" {
			protocol = "all"
		}
		port := "all"
		if p.FromPort != 0 {
			port = strconv.Itoa(p.FromPort)
			if p.ToPort != p.FromPort {
				port += "-" + strconv.Itoa(p.ToPort)
			}
		}
		groups := []string{}
		for _, srcGroup := range p.Groups {
			groups = append(groups, srcGroup.GroupId)
		}
		g.IpPermissions = append(g.IpPermissions, &IpPermission{
			Protocol: protocol,
			Ports:    port,
			IpRanges: p.IpRanges,
			Groups:   groups,
		})
	}
	return g
}
