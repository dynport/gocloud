package ec2

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustReadFixture(name string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + name)
	if e != nil {
		panic(e.Error())
	}
	return b
}

func TestMarshallRunInstanceResponse(t *testing.T) {
	rsp := &RunInstancesResponse{}
	e := xml.Unmarshal(mustReadFixture("run_instances_response.xml"), rsp)
	assert.Nil(t, e)
	i := rsp.Instances[0]
	assert.Equal(t, i.InstanceId, "i-1122")
}

func TestMarshalling(t *testing.T) {
	f := mustReadFixture("describe_images.xml")
	rsp := &DescribeImagesResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.Equal(t, rsp.RequestId, "59dbff89-35bd-4eac-99ed-be587EXAMPLE")
	assert.NotNil(t, f)
	assert.Equal(t, len(rsp.Images), 1)

	img := rsp.Images[0]

	assert.Equal(t, img.ImageId, "ami-1a2b3c4d")
}

func TestDescribeInstances(t *testing.T) {
	f := mustReadFixture("describe_instances.xml")
	assert.NotNil(t, f)
}

func TestMarshalTags(t *testing.T) {
	f := mustReadFixture("describe_tags.xml")
	rsp := &DescribeTagsResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.Equal(t, len(rsp.Tags), 6)
	tag := rsp.Tags[1]
	assert.Equal(t, tag.Key, "stack")
	assert.Equal(t, tag.Value, "Production")

	tag = &Tag{Key: "Name", Value: "staging"}
	b, e := xml.Marshal(tag)
	s := string(b)
	assert.Contains(t, s, "staging")
	assert.NotContains(t, s, "resourceId")
}

func TestDescribeKeyPair(t *testing.T) {
	f := mustReadFixture("describe_key_pairs.xml")
	rsp := &DescribeKeyPairsResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, rsp)
	assert.Equal(t, len(rsp.KeyPairs), 1)

	pair := rsp.KeyPairs[0]
	assert.Equal(t, pair.KeyName, "my-key-pair")
	assert.Equal(t, pair.KeyFingerprint, "1f:51:ae:28:bf:89:e9:d8:1f:25:5d:37:2d:7d:b8:ca:9f:f5:f1:6f")
}

func TestAddresses(t *testing.T) {
	f := mustReadFixture("describe_addresses.xml")
	assert.NotNil(t, f)
	rsp := &DescribeAddressesResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, rsp)
	assert.Equal(t, len(rsp.Addresses), 1)
	a := rsp.Addresses[0]
	assert.Equal(t, a.PublicIp, "203.0.113.41")
}

func TestSecurityGroups(t *testing.T) {
	f := mustReadFixture("describe_security_groups.xml")
	assert.NotNil(t, f)
	rsp := &DescribeSecurityGroupsResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, rsp)
	assert.Equal(t, len(rsp.SecurityGroups), 2)

	group := rsp.SecurityGroups[0]
	assert.Equal(t, group.GroupId, "sg-1a2b3c4d")
	assert.Equal(t, len(group.IpPermissions), 1)

	perm := group.IpPermissions[0]
	assert.Equal(t, perm.IpProtocol, "tcp")
	assert.Equal(t, perm.FromPort, 80)

	assert.Equal(t, len(perm.IpRanges), 1)
	assert.Equal(t, perm.IpRanges[0], "0.0.0.0/0")

	add2 := rsp.SecurityGroups[1]
	assert.NotNil(t, add2)
	assert.Equal(t, len(add2.IpPermissions[0].Groups), 1)
}
