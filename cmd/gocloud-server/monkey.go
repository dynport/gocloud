package main

import (
	"encoding/xml"
	"net/url"
	"strconv"
	"time"

	"github.com/dynport/gocloud/aws"
	"github.com/dynport/gocloud/aws/ec2"
)

type Snapshot struct {
	SnapshotId  string    `xml:"snapshotId"`  // snap-1a2b3c4d</snapshotId>
	VolumeId    string    `xml:"volumeId"`    // vol-1a2b3c4d</volumeId>
	Status      string    `xml:"status"`      // pending</status>
	StartTime   time.Time `xml:"startTime"`   // YYYY-MM-DDTHH:MM:SS.SSSZ</startTime>
	Progress    string    `xml:"progress"`    // 80%</progress>
	OwnerId     string    `xml:"ownerId"`     // 111122223333</ownerId>
	VolumeSize  int       `xml:"volumeSize"`  // 15</volumeSize>
	Description string    `xml:"description"` // Daily Backup</description>
}

type DescribeSnapshotsResponse struct {
	XMLName   xml.Name    `xml:"DescribeSnapshotsResponse"`
	Snapshots []*Snapshot `xml:"snapshotSet>item"`
}

type DescribeVolumesResponse struct {
	XMLName xml.Name  `xml:"DescribeVolumesResponse"`
	Volumes []*Volume `xml:"volumeSet>item"`
}

type VolumeAttachment struct {
	VolumeId            string    `xml:"volumeId"`            // vol-1a2b3c4d</volumeId>
	InstanceId          string    `xml:"instanceId"`          // i-1a2b3c4d</instanceId>
	Device              string    `xml:"device"`              // /dev/sdh</device>
	Status              string    `xml:"status"`              // attached</status>
	AttachTime          time.Time `xml:"attachTime"`          // YYYY-MM-DDTHH:MM:SS.SSSZ</attachTime>
	DeleteOnTermination bool      `xml:"deleteOnTermination"` // false</deleteOnTermination>
}

type Volume struct {
	VolumeId         string              `xml:"volumeId"`         // vol-1a2b3c4d</volumeId>
	Size             string              `xml:"size"`             // 80</size>
	SnapshotId       string              `xml:"snapshotId"`       // >
	AvailabilityZone string              `xml:"availabilityZone"` // us-east-1a</availabilityZone>
	Status           string              `xml:"status"`           // in-use</status>
	CreateTime       string              `xml:"createTime"`       // YYYY-MM-DDTHH:MM:SS.SSSZ</createTime>
	Attachments      []*VolumeAttachment `xml:"attachmentSet>item"`
	VolumeType       string              `xml:"volumeType"` // standard</volumeType>
}

func loadResource(action string, i interface{}, values url.Values) error {
	if values == nil {
		values = url.Values{}
	}
	values["Version"] = []string{ec2.API_VERSIONS_EC2}
	values["Action"] = []string{action}
	rsp, e := ec2Client.DoSignedRequest("GET", ec2Client.Endpoint(), values.Encode(), nil)
	if e != nil {
		return e
	}
	e = aws.ExtractError(rsp.Content)
	if e != nil {
		return e
	}
	e = xml.Unmarshal(rsp.Content, i)
	return e
}

type DescribeSnapshotsParameters struct {
	Owners []string
}

func describeSnapshots(params *DescribeSnapshotsParameters) ([]*Snapshot, error) {
	v := url.Values{}
	for i, o := range params.Owners {
		v.Add("Owner."+strconv.Itoa(i+1), o)
	}
	r := &DescribeSnapshotsResponse{}
	e := loadResource("DescribeSnapshots", r, v)
	return r.Snapshots, e
}

func describeVolumes() ([]*Volume, error) {
	r := &DescribeVolumesResponse{}
	e := loadResource("DescribeVolumes", r, nil)
	return r.Volumes, e
}

func describeVpcs() (VpcList, error) {
	r := &DescribeVpcsResponse{}
	e := loadResource("DescribeVpcs", r, nil)
	list := VpcList{}
	for _, v := range r.Vpcs {
		list = append(list, NewVpc(v))
	}
	return list, e
}

type Subnet struct {
	SubnetId                string     `xml:"subnetId"`                // subnet-86987ee3</subnetId>
	State                   string     `xml:"state"`                   // available</state>
	VpcId                   string     `xml:"vpcId"`                   // vpc-8e273bec</vpcId>
	CidrBlock               string     `xml:"cidrBlock"`               // 172.31.0.0/20</cidrBlock>
	AvailableIpAddressCount string     `xml:"availableIpAddressCount"` // 4091</availableIpAddressCount>
	AvailabilityZone        string     `xml:"availabilityZone"`        // eu-west-1c</availabilityZone>
	DefaultForAz            bool       `xml:"defaultForAz"`            // false</defaultForAz>
	MapPublicIpOnLaunch     bool       `xml:"mapPublicIpOnLaunch"`     // false</mapPublicIpOnLaunch>
	Tags                    []*ec2.Tag `xml:"tagSet>item"`
}

type SubnetList []*Subnet

func (list SubnetList) Query(q string) SubnetList {
	filtered := SubnetList{}
	for _, s := range list {
		if s.VpcId == q {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

type DescribeSubnetsResponse struct {
	XMLName xml.Name  `xml:"DescribeSubnetsResponse"`
	Subnets []*Subnet `xml:"subnetSet>item"`
}

func describeSubnets() (SubnetList, error) {
	values := url.Values{"Version": {ec2.API_VERSIONS_EC2}, "Action": {"DescribeSubnets"}}
	rsp, e := ec2Client.DoSignedRequest("GET", ec2Client.Endpoint(), values.Encode(), nil)
	if e != nil {
		return nil, e
	}
	e = aws.ExtractError(rsp.Content)
	if e != nil {
		return nil, e
	}
	r := &DescribeSubnetsResponse{}
	e = xml.Unmarshal(rsp.Content, r)
	if e != nil {
		return nil, e
	}
	return r.Subnets, nil
}
