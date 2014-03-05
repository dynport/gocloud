package rds

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dynport/gocloud/aws"
	"github.com/dynport/gocloud/aws/ec2"
)

type Client struct {
	*aws.Client
	CustomRegion string
}

func (client *Client) Endpoint() string {
	prefix := "https://rds"
	if client.Client.Region != "" {
		prefix += "." + client.Client.Region
	}
	return prefix + ".amazonaws.com"
}

func NewFromEnv() *Client {
	return &Client{Client: aws.NewFromEnv()}
}

type DescribeDBInstances struct {
	DBInstanceIdentifier string
	Filters              []*ec2.Filter
	Marker               string
	MaxRecords           int
}

func newAction(name string) url.Values {
	return url.Values{"Action": {name}, "Version": {Version}}
}

func (d *DescribeDBInstances) Execute(client *Client) (*DescribeDBInstancesResponse, error) {
	v := newAction("DescribeDBInstances")
	e := loadValues(v, d)
	if e != nil {
		return nil, e
	}
	r := &DescribeDBInstancesResponse{}
	e = client.loadResource("GET", client.Endpoint()+"?"+v.Encode(), nil, r)
	return r, e
}

func (client *Client) loadResource(method string, url string, r io.Reader, i interface{}) error {
	req, e := http.NewRequest(method, url, r)
	if e != nil {
		return e
	}
	client.SignAwsRequestV2(req, time.Now())
	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()

	if rsp.Status[0] != '2' {
		b, _ := ioutil.ReadAll(rsp.Body)
		return fmt.Errorf("expected status 2xx, got %s (payload=%q", rsp.Status, string(b))
	}

	buf := &bytes.Buffer{}
	reader := io.TeeReader(rsp.Body, buf)
	e = xml.NewDecoder(reader).Decode(i)
	return e
}

type DescribeDBInstancesResponse struct {
	XMLName                   xml.Name                   `xml:"DescribeDBInstancesResponse"`
	DescribeDBInstancesResult *DescribeDBInstancesResult `xml:"DescribeDBInstancesResult"`
}

type DescribeDBInstancesResult struct {
	Instances []*DBInstance `xml:"DBInstances>DBInstance"`
}

type Endpoint struct {
	Port    string `xml:"Port"`
	Address string `xml:"Address"`
}

type VpcSecurityGroupMembership struct {
	Status             string `xml:"Status"`
	VpcSecurityGroupId string `xml:"VpcSecurityGroupId"`
}

type DBSecurityGroup struct {
	Status              string `xml:"Status"`
	DBSecurityGroupName string `xml:"DBSecurityGroupName"`
}

type DBInstance struct {
	LatestRestorableTime       string                        `xml:"LatestRestorableTime"`
	Engine                     string                        `xml:"Engine"`
	PendingModifiedValues      interface{}                   `xml:"PendingModifiedValues"`
	BackupRetentionPeriod      string                        `xml:"BackupRetentionPeriod"`
	MultiAZ                    bool                          `xml:"MultiAZ"`
	LicenseModel               string                        `xml:"LicenseModel"`
	DBInstanceStatus           string                        `xml:"DBInstanceStatus"`
	EngineVersion              string                        `xml:"EngineVersion"`
	Endpoint                   *Endpoint                     `xml:"Endpoint"`
	DBInstanceIdentifier       string                        `xml:"DBInstanceIdentifier"`
	VpcSecurityGroups          []*VpcSecurityGroupMembership `xml:"VpcSecurityGroups"`
	DBSecurityGroups           []*DBSecurityGroup            `xml:"DBSecurityGroups"`
	PreferredBackupWindow      string                        `xml:"PreferredBackupWindow"`
	AutoMinorVersionUpgrade    bool                          `xml:"AutoMinorVersionUpgrade"`
	PreferredMaintenanceWindow string                        `xml:"PreferredMaintenanceWindow"`
	AvailabilityZone           string                        `xml:"AvailabilityZone"`
	InstanceCreateTime         time.Time                     `xml:"InstanceCreateTime"`
	AllocatedStorage           int                           `xml:"AllocatedStorage"`
	DBInstanceClass            string                        `xml:"DBInstanceClass"`
	MasterUsername             string                        `xml:"MasterUsername"`
}
