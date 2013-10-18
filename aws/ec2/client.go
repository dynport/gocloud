package ec2

import (
	"encoding/xml"
	"fmt"
	"github.com/dynport/gocloud/aws"
	"log"
	"net/url"
	"strconv"
)

func NewFromEnv() *Client {
	return &Client{
		aws.NewFromEnv(),
	}
}

type Client struct {
	*aws.Client
}

const (
	API_VERSIONS_EC2     = "2013-08-15"
	CANONICAL_OWNER_ID   = "099720109477"
	SELF_OWNER_ID        = "self"
	UBUNTU_PREFIX        = "ubuntu/images/ubuntu-*"
	UBUNTU_RARING_PREFIX = "ubuntu/images/ubuntu-raring*"
	UBUNTU_SAUCY_PREFIX  = "ubuntu/images/ubuntu-saucy*"
	ENDPOINT             = "https://eu-west-1.ec2.amazonaws.com"
)

type ImageFilter struct {
	Owner string
	Name  string
}

type ImageList []*Image

type InstanceList []*Instance

func (list ImageList) Len() int {
	return len(list)
}

func (list ImageList) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func (list ImageList) Less(a, b int) bool {
	return list[a].Name > list[b].Name
}

type RunInstancesConfig struct {
	ImageId          string
	MinCount         int
	MaxCount         int
	InstanceType     string
	AvailabilityZone string
	KeyName          string
	SecurityGroups   []string
}

func queryForAction(action string) string {
	values := &url.Values{}
	values.Add("Version", API_VERSIONS_EC2)
	values.Add("Action", action)
	return values.Encode()
}

func (client *Client) DescribeTags() (tags TagList, e error) {
	query := queryForAction("DescribeTags")
	raw, e := client.DoSignedRequest("GET", ENDPOINT, query, nil)
	if e != nil {
		return tags, e
	}
	rsp := &DescribeTagsResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return tags, e
	}
	return rsp.Tags, e
}

func (client *Client) CreateTags(resourceIds []string, tags map[string]string) error {
	values := &url.Values{}
	for i, id := range resourceIds {
		values.Add("ResourceId."+strconv.Itoa(i), id)
	}
	tagsCount := 1
	for k, v := range tags {
		prefix := fmt.Sprintf("Tag.%d.", tagsCount)
		values.Add(prefix+"Key", k)
		values.Add(prefix+"Value", v)
		tagsCount++
	}
	query := queryForAction("CreateTags") + "&" + values.Encode()
	_, e := client.DoSignedRequest("POST", ENDPOINT, query, nil)
	if e != nil {
		return e
	}
	return nil
}

func (client *Client) TerminateInstances(ids []string) error {
	query := queryForAction("TerminateInstances")
	for i, id := range ids {
		query += fmt.Sprintf("&InstanceId.%d=%s", i, id)
	}
	rsp, e := client.DoSignedRequest("DELETE", ENDPOINT, query, nil)
	if e != nil {
		return e
	}
	log.Println(string(rsp.Content))
	log.Printf("terminates instances: %d", rsp.StatusCode)
	return nil
}

type RunInstancesResponse struct {
	RequestId     string      `xml:"requestId"`
	ReservationId string      `xml:"reservationId"`
	OwnerId       string      `xml:"ownerId"`
	Instances     []*Instance `xml:"instancesSet>item"`
}

func (client *Client) RunInstances(config *RunInstancesConfig) (list InstanceList, e error) {
	if config.MinCount == 0 {
		config.MinCount = 1
	}

	if config.MaxCount == 0 {
		config.MaxCount = 1
	}

	if config.ImageId == "" {
		return list, fmt.Errorf("ImageId must be provided")
	}

	values := &url.Values{}
	values.Add("MinCount", strconv.Itoa(config.MinCount))
	values.Add("MaxCount", strconv.Itoa(config.MaxCount))
	values.Add("ImageId", config.ImageId)

	if config.InstanceType != "" {
		values.Add("InstanceType", config.InstanceType)
	}

	if config.KeyName != "" {
		values.Add("KeyName", config.KeyName)
	}

	if config.AvailabilityZone != "" {
		values.Add("Placement.AvailabilityZone", config.AvailabilityZone)
	}

	for i, sg := range config.SecurityGroups {
		values.Add("SecurityGroup."+strconv.Itoa(i+1), sg)
	}

	query := queryForAction("RunInstances") + "&" + values.Encode()

	raw, e := client.DoSignedRequest("POST", ENDPOINT, query, nil)
	if e != nil {
		return list, e
	}
	rsp := &RunInstancesResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return list, e
	}
	return InstanceList(rsp.Instances), nil
}

func (client *Client) DescribeInstances() (instances []*Instance, e error) {
	raw, e := client.DoSignedRequest("GET", ENDPOINT, "Version="+API_VERSIONS_EC2+"&Action=DescribeInstances", nil)
	if e != nil {
		return instances, e
	}
	rsp := &DescribeInstancesResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return instances, e
	}
	return rsp.Instances(), nil
}
