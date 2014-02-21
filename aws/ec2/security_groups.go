package ec2

import (
	"encoding/xml"
)

type DescribeSecurityGroupsResponse struct {
	SecurityGroups []*SecurityGroup `xml:"securityGroupInfo>item"`
}

type IpPermission struct {
	IpProtocol string           `xml:"ipProtocol,omitempty"`  // tcp</ipProtocol>
	FromPort   int              `xml:"fromPort,omitempty"`    // 80</fromPort>
	ToPort     int              `xml:"toPort,omitempty"`      // 80</toPort>
	Groups     []*SecurityGroup `xml:"groups>item,omitempty"` //
	IpRanges   []string         `xml:"ipRanges>item>cidrIp"`
}

type SecurityGroup struct {
	OwnerId          string          `xml:"ownerId,omitempty"`          // 111122223333</ownerId>
	GroupId          string          `xml:"groupId,omitempty"`          // sg-1a2b3c4d</groupId>
	GroupName        string          `xml:"groupName,omitempty"`        // WebServers</groupName>
	GroupDescription string          `xml:"groupDescription,omitempty"` // Web Servers</groupDescription>
	VpcId            string          `xml:"vpcId,omitempty"`            //
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
