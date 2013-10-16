package ec2

import (
	"encoding/xml"
)

type DescribeSecurityGroupsResponse struct {
	SecurityGroups []*SecurityGroup `xml:"securityGroupInfo>item"`
}

type IpPermission struct {
	IpProtocol string           `xml:"ipProtocol"`  // tcp</ipProtocol>
	FromPort   int              `xml:"fromPort"`    // 80</fromPort>
	ToPort     int              `xml:"toPort"`      // 80</toPort>
	Groups     []*SecurityGroup `xml:"groups>item"` //
	IpRanges   []string         `xml:"ipRanges>item>cidrIp"`
}

type SecurityGroup struct {
	OwnerId          string          `xml:"ownerId"`          // 111122223333</ownerId>
	GroupId          string          `xml:"groupId"`          // sg-1a2b3c4d</groupId>
	GroupName        string          `xml:"groupName"`        // WebServers</groupName>
	GroupDescription string          `xml:"groupDescription"` // Web Servers</groupDescription>
	VpcId            string          `xml:"vpcId/"`           //
	IpPermissions    []*IpPermission `xml:"ipPermissions>item"`
}

func (client *Client) DescribeSecurityGroups() (groups []*SecurityGroup, e error) {
	raw, e := client.DoSignedRequest("GET", ENDPOINT, queryForAction("DescribeSecurityGroups"), nil)
	if e != nil {
		return groups, e
	}
	rsp := &DescribeSecurityGroupsResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return groups, e
	}
	return rsp.SecurityGroups, nil
}
