package main

import (
	"net/url"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSerialize(t *testing.T) {
	Convey("Test serialize", t, func() {
		type NetworkInterface struct {
			Interface string `aws:"Interface"`
		}
		type DescribeInstance struct {
			InstanceIds       []string            `aws:"InstanceId"`
			Name              string              `aws:"Name"`
			Empty             string              `aws:"Empty"`
			NetworkInterfaces []*NetworkInterface `aws:"NetworkInterface"`
		}
		d := &DescribeInstance{
			InstanceIds: []string{"instance1"},
			Name:        "test",
			NetworkInterfaces: []*NetworkInterface{
				{Interface: "eth0"},
			},
		}
		v := url.Values{}
		urlValues(v, d, "")
		So(len(v), ShouldEqual, 3)
		So(v.Get("InstanceId.1"), ShouldEqual, "instance1")
		So(v.Get("Name"), ShouldEqual, "test")
		So(v.Get("NetworkInterface.1.Interface"), ShouldEqual, "eth0")
	})
}
