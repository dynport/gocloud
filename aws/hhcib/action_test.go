package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAction(t *testing.T) {
	Convey("Action", t, func() {
		So(1, ShouldEqual, 1)
		a := &Action{
			Name: "DescribeInstances",
			RequestParameters: RequestParameters{
				{
					Name: "InstanceId.n",
					Type: "String",
				},
				{
					Name: "Filter.n.Name",
					Type: "String",
				},
				{
					Name: "Filter.n.Value.m",
					Type: "String",
				},
			},
		}
		typ := a.RequestType()
		So(typ, ShouldNotBeNil)
		So(typ.Name, ShouldEqual, "DescribeInstances")
		So(len(typ.Fields), ShouldEqual, 2)

		So(typ.Fields[0].Name, ShouldEqual, "InstanceIds")
		So(typ.Fields[0].Type, ShouldEqual, "[]string")
		So(typ.Fields[0].Comments["aws"], ShouldEqual, "InstanceId")

		So(typ.Fields[1].Name, ShouldEqual, "Filters")
		So(typ.Fields[1].Type, ShouldEqual, "[]*Filter")
		So(typ.Fields[1].Comments["aws"], ShouldEqual, "Filter")
	})
}
