package main

import (
	"github.com/dynport/gocloud/aws/pricing"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func mustRead(t *testing.T, file string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + file)
	if e != nil {
		t.Fatal(e.Error())
	}
	return b
}

func pickInstance(instances []*pricing.InstanceTypeConfig, name string) *pricing.InstanceTypeConfig {
	for _, i := range instances {
		if i.Name == name {
			return i
		}
	}
	return nil
}

func TestParseInstanceTypes(t *testing.T) {
	Convey("Parse instance types", t, func() {
		b := mustRead(t, "instance_types.html")
		types, e := parseInstanceTypes(b)
		So(e, ShouldBeNil)
		So(types, ShouldNotBeNil)
		So(len(types), ShouldEqual, 29)

		micro := pickInstance(types, "t1.micro")
		So(micro, ShouldNotBeNil)
		So(micro.Family, ShouldEqual, "Micro instances")
		So(micro.Arch, ShouldEqual, "32-bit or 64-bit")
		So(micro.Storage, ShouldEqual, "EBS only")

		c3large := pickInstance(types, "c3.large")
		So(c3large, ShouldNotBeNil)
		So(c3large.Family, ShouldEqual, "Compute optimized")
		So(c3large.Arch, ShouldEqual, "64-bit")
		So(c3large.Memory, ShouldEqual, 3.75)
		So(c3large.Storage, ShouldEqual, "2 x 16 SSD")
		So(c3large.EbsOptimizable, ShouldEqual, false)
		So(c3large.NetworkPerformance, ShouldEqual, "Moderate")

		m1xlarge := pickInstance(types, "m1.xlarge")
		So(m1xlarge, ShouldNotBeNil)
		So(m1xlarge.Family, ShouldEqual, "General purpose")
		So(m1xlarge.Arch, ShouldEqual, "64-bit")
		So(m1xlarge.Memory, ShouldEqual, 15)
		So(m1xlarge.Storage, ShouldEqual, "4 x 420")
		So(m1xlarge.EbsOptimizable, ShouldEqual, true)
		So(m1xlarge.NetworkPerformance, ShouldEqual, "High")

		m3medium := pickInstance(types, "m3.medium")
		So(m3medium, ShouldNotBeNil)
		So(m3medium.Cpus, ShouldEqual, 1)
		So(m3medium.ECUs, ShouldEqual, 3)
	})
}
