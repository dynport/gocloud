package elb

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func mustReadFixture(name string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + name)
	if e != nil {
		panic(e.Error())
	}
	return b
}

func TestMarshalling(t *testing.T) {
	f := mustReadFixture("describe_load_balancers.xml")
	rsp := &DescribeLoadBalancersResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	lbs := rsp.LoadBalancers
	assert.Equal(t, len(lbs), 1)
	lb := lbs[0]
	assert.Equal(t, lb.LoadBalancerName, "MyLoadBalancer")
	assert.Equal(t, lb.CreatedTime.Unix(), 1369430131)
	assert.Equal(t, lb.CanonicalHostedZoneName, "MyLoadBalancer-123456789.us-east-1.elb.amazonaws.com")
	assert.Equal(t, len(lb.AvailabilityZones), 1)
	assert.Equal(t, lb.AvailabilityZones[0], "us-east-1a")
	assert.Equal(t, len(lb.Subnets), 0)
	assert.Equal(t, lb.HealthCheckTarget, "HTTP:80/")
	assert.Equal(t, lb.HealthCheckInterval, 90)
	assert.Equal(t, len(lb.Listeners), 1)
	assert.Equal(t, lb.SourceSecurityGroupOwnerAlias, "amazon-elb")
	listener := lb.Listeners[0]
	assert.Equal(t, listener.Protocol, "HTTP")

	assert.Equal(t, len(lb.Instances), 1)
	assert.Equal(t, lb.Instances[0], "i-e4cbe38d")
}

func TestMarshalInstanceHealth(t *testing.T) {
	f := mustReadFixture("describe_instances_health.xml")
	rsp := &DescribeInstanceHealthResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Equal(t, len(rsp.InstanceStates), 1)
	assert.Nil(t, e)
	assert.NotNil(t, rsp)

	state := rsp.InstanceStates[0]
	assert.Equal(t, state.Description, "Instance registration is still in progress.")
	assert.Equal(t, state.InstanceId, "i-315b7e51")
}
