package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNetworkInterfaces(t *testing.T) {
	Convey("TestNetworkInterfaces", t, func() {
		p := RequestParameters{
			{Name: "NetworkInterface.n.NetworkInterfaceId", Type: "String"},
		}
		So(p, ShouldNotBeNil)
		fields := p.fields()
		So(len(fields), ShouldEqual, 1)
		So(fields[0].Type, ShouldEqual, "[]*RequestNetworkInterface")
		So(fields[0].CustomType, ShouldNotBeNil)
		So(fields[0].CustomType.Name, ShouldEqual, "RequestNetworkInterface")
	})
}

func TestAuthorizeSecurityGroupEgress(t *testing.T) {
	Convey("AuthorizeSecurityGroupEgress", t, func() {
		p := RequestParameters{
			{Name: "IpPermissions.n.ToPort", Type: "Integer"},
		}
		So(p, ShouldNotBeNil)

		fields := p.fields()

		So(fields[0].Name, ShouldEqual, "IpPermissions")
		So(fields[0].Type, ShouldEqual, "[]*IpPermissions")
		So(fields[0].Comments["aws"], ShouldEqual, "IpPermissions")
		So(fields[0].CustomType.Name, ShouldEqual, "IpPermissions")

		t.Log(fields[0].CustomType.Fields)
		So(len(fields[0].CustomType.Fields), ShouldEqual, 1)
		So(fields[0].CustomType.Fields[0].Type, ShouldEqual, "int")
	})
}

func TestNestedRequestParameters(t *testing.T) {
	Convey("Nested request Params", t, func() {
		p := RequestParameters{
			{Name: "IpPermissions.n.Groups.m.GroupId", Type: "String"},
			{Name: "IpPermissions.n.Groups.m.UserId", Type: "String"},
		}
		So(p, ShouldNotBeNil)

		fields := p.fields()

		So(fields[0].Name, ShouldEqual, "IpPermissions")
		So(fields[0].Type, ShouldEqual, "[]*IpPermissions")
		So(fields[0].Comments["aws"], ShouldEqual, "IpPermissions")
		So(fields[0].CustomType.Name, ShouldEqual, "IpPermissions")

		t.Log(fields[0].CustomType.Fields)
		So(len(fields[0].CustomType.Fields), ShouldEqual, 1)
	})
}

func TestRequestParameters(t *testing.T) {
	Convey("RequestParameters", t, func() {
		p := RequestParameters{
			{Name: "InstanceId.n", Type: "String"},
			{Name: "Filter.n.Name", Type: "String"},
			{Name: "Filter.n.Value.m", Type: "String"},
			{Name: "LaunchPermission.Add.n.UserId", Type: "String"},
			{Name: "LaunchPermission.Add.n.Group", Type: "String"},
		}
		So(p, ShouldNotBeNil)
		fields := p.fields()
		So(len(fields), ShouldEqual, 3)

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
	})
}
