package jiffybox

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustReadFixture(t *testing.T, name string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + name)
	if e != nil {
		t.Fatal("fixture " + name + " does not exist")
	}
	return b
}

func TestJiffyBoxes(t *testing.T) {
	f := mustReadFixture(t, "jiffyBoxes.json")
	assert.NotNil(t, f)

	rsp := &JiffyBoxesResponse{}
	e := json.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.Equal(t, len(rsp.Messages), 0)

	assert.Equal(t, len(rsp.Servers()), 1)

	server := rsp.Server()
	assert.Equal(t, server.Id, 12345)
	assert.Equal(t, server.Name, "Test")
	assert.Equal(t, len(server.Ips), 2)
	assert.Equal(t, server.Ips["public"], []string{"188.93.14.176"})
	assert.Equal(t, server.Status, "READY")

	plan := server.Plan
	assert.Equal(t, plan.Id, 22)
	assert.Equal(t, plan.Name, "CloudLevel 3")
	assert.Equal(t, plan.RamInMB, 8192)

	assert.Equal(t, server.Metadata, map[string]string{
		"createdby": "JiffyBoxTeam",
	})

	ap := server.ActiveProfile
	assert.Equal(t, ap.Name, "Standard")
	assert.Equal(t, ap.Created, 1234567890)

	assert.Equal(t, len(ap.DisksHash), 2)
	assert.Equal(t, len(ap.Disks()), 2)

	disk := ap.DisksHash["xvda"]

	assert.Equal(t, disk.Name, "CentOS 5.4")
	assert.Equal(t, disk.SizeInMB, 81408)
}

func TestUnmarshalling(t *testing.T) {
	f := mustReadFixture(t, "error_creating_response.json")
	rsp := &ErrorResponse{}
	e := json.Unmarshal(f, rsp)
	assert.Nil(t, e)
	t.Log(rsp.Result)

	f = mustReadFixture(t, "no_module_response.json")
	rsp = &ErrorResponse{}
	e = json.Unmarshal(f, rsp)
	assert.Nil(t, e)
	t.Log(rsp.Result)

}
