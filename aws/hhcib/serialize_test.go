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
		type DescribeInstances struct {
			InstanceIds       []string            `aws:"InstanceId"`
			Name              string              `aws:"Name"`
			Empty             string              `aws:"Empty"`
			Count             *IntValue           `aws:"Count"`
			Enabled           *BoolValue          `aws:"Enabled"`
			Rate              *FloatValue         `aws:"Rate"`
			NetworkInterfaces []*NetworkInterface `aws:"NetworkInterface"`
		}
		d := &DescribeInstances{
			InstanceIds: []string{"instance1"},
			Name:        "test",
			NetworkInterfaces: []*NetworkInterface{
				{Interface: "eth0"},
			},
		}
		v := url.Values{}
		urlValues(v, d, "")
		So(v.Get("Action"), ShouldEqual, "DescribeInstances")
		So(v.Get("Version"), ShouldEqual, "2013-10-15")
		So(v.Get("InstanceId.1"), ShouldEqual, "instance1")
		So(v.Get("Name"), ShouldEqual, "test")
		So(v.Get("NetworkInterface.1.Interface"), ShouldEqual, "eth0")

		v = url.Values{}
		d.Count = &IntValue{10}
		urlValues(v, d, "")
		So(v.Get("Count"), ShouldEqual, "10")

		d.Enabled = &BoolValue{true}
		urlValues(v, d, "")
		So(v.Get("Enabled"), ShouldEqual, "true")
	})
}
