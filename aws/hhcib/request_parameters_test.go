package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRequestParameters(t *testing.T) {
	Convey("RequestParameters", t, func() {
		p := RequestParameters{
			{Name: "InstanceId.n", Type: "String"},
			{Name: "Filter.n.Name", Type: "String"},
			{Name: "Filter.n.Value.m", Type: "String"},
			{Name: "LaunchPermission.Add.n.UserId", Type: "String"},
			{Name: "LaunchPermission.Add.n.Group", Type: "String"},
			{Name: "IpPermissions.n.Groups.m.GroupId", Type: "String"},
			{Name: "IpPermissions.n.Groups.m.UserId", Type: "String"},
		}
		So(p, ShouldNotBeNil)
		fields := p.fields()
		So(len(fields), ShouldEqual, 4)

		So(fields[0].Name, ShouldEqual, "InstanceIds")
		So(fields[0].Type, ShouldEqual, "[]string")
		So(fields[0].Comments["aws"], ShouldEqual, "InstanceId")

		So(fields[1].Name, ShouldEqual, "Filters")
		So(fields[1].Type, ShouldEqual, "[]*Filter")
		So(fields[1].Comments["aws"], ShouldEqual, "Filter")

		So(fields[2].Name, ShouldEqual, "LaunchPermissionAdds")
		So(fields[2].Type, ShouldEqual, "[]*LaunchPermissionAdd")
		So(fields[2].Comments["aws"], ShouldEqual, "LaunchPermission.Add")
		So(fields[2].CustomType.Name, ShouldEqual, "LaunchPermissionAdd")

		So(fields[3].Name, ShouldEqual, "IpPermissions")
		So(fields[3].Type, ShouldEqual, "[]*IpPermissions")
		So(fields[3].Comments["aws"], ShouldEqual, "IpPermissions")
		So(fields[3].CustomType.Name, ShouldEqual, "IpPermissions")

		So(len(fields[3].CustomType.Fields), ShouldEqual, 1)
		So(1, ShouldEqual, 2)
	})
}
