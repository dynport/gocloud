package jiffybox

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributions(t *testing.T) {
	f := mustReadFixture(t, "distributions.json")
	assert.NotNil(t, f)

	rsp := &DistributionsResponse{}
	e := json.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, rsp)

	assert.Equal(t, len(rsp.DistributionsMap), 2)
	assert.Equal(t, len(rsp.Distributions()), 2)

	dist := rsp.Distributions()[0]
	assert.Equal(t, dist.Name, "CentOS 5.4")
	assert.Equal(t, dist.Key, "centos_5_4_32bit")
}
