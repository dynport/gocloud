package route53

import (
	"encoding/xml"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSerialize(t *testing.T) {
	req := NewChangeResourceRecordSets(
		&ChangeBatch{
			Comment: "This is a comment",
			Changes: []*Change{
				{Action: "CREATE"},
			},
		},
	)
	Convey("Serialize", t, func() {
		b, e := xml.Marshal(req)
		So(e, ShouldBeNil)
		So(b, ShouldNotBeNil)
		s := string(b)
		So(s, ShouldContainSubstring, "xmlns")
		So(s, ShouldContainSubstring, "<Changes>")
		So(s, ShouldContainSubstring, "<ChangeBatch>")
		So(s, ShouldContainSubstring, "<Change>")
	})
}

func mustReadFixture(name string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + name)
	if e != nil {
		panic(e.Error())
	}
	return b
}

func TestParseHostedZones(t *testing.T) {
	f := mustReadFixture("list_hosted_zones.xml")
	rsp := &ListHostedZonesResponse{}
	xml.Unmarshal(f, rsp)
	assert.NotNil(t, rsp)
	assert.Equal(t, len(rsp.HostedZones), 1)

	assert.Equal(t, rsp.MaxItems, 1)
	assert.Equal(t, rsp.IsTruncated, true)

	zone := rsp.HostedZones[0]
	assert.Equal(t, zone.Id, "/hostedzone/Z111111QQQQQQQ")
	assert.Equal(t, zone.Name, "example2.com.")
	assert.Equal(t, zone.CallerReference, "MyUniqueIdentifier2")
	assert.Equal(t, zone.ResourceRecordSetCount, 42)
}

func TestParseGetHostedZOneResponse(t *testing.T) {
	f := mustReadFixture("get_hosted_zone_response.xml")
	rsp := &GetHostedZoneResponse{}
	e := xml.Unmarshal(f, rsp)
	zone := rsp.HostedZone

	assert.Nil(t, e)
	assert.NotNil(t, zone)
	assert.Equal(t, zone.Id, "/hostedzone/Z1PA6795UKMFR9")
	assert.Equal(t, zone.Name, "example.com.")

	nameServers := rsp.NameServers
	assert.Equal(t, len(nameServers), 4)
	assert.Equal(t, nameServers[0], "ns-2048.awsdns-64.com")
}

func TestListResourceRecordSets(t *testing.T) {
	f := mustReadFixture("list_resource_record_sets.xml")
	rsp := &ListResourceRecordSetsResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, rsp)
	assert.Equal(t, len(rsp.ResourceRecordSets), 1)

	rrs := rsp.ResourceRecordSets[0]
	assert.Equal(t, rrs.Name, "example.com.")
	assert.Equal(t, rrs.Type, "NS")
	assert.Equal(t, rrs.TTL, 172800)

	assert.Equal(t, len(rrs.ResourceRecords), 4)

	resourceRecord := rrs.ResourceRecords[0]
	assert.NotNil(t, resourceRecord)
	assert.Equal(t, resourceRecord.Value, "ns-2048.awsdns-64.com.")
}
