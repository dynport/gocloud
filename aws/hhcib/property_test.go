package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type toTypeFieldTest struct {
	Name, Type, Description, Response string
}

func TestPropertiesToRequestTypeDefinition(t *testing.T) {
	Convey("toRequestTypeDefinition", t, func() {
		cases := []*toTypeFieldTest{
			{Name: "InstanceId.n", Type: "String", Response: "InstanceIds []string `aws:" + `"InstanceId"` + "`"},
			{Name: "MaxResults", Type: "Integer", Response: "MaxResults *IntValue `aws:" + `"MaxResults"` + "`"},
			{Name: "Storage.S3.Bucket", Type: "String", Response: "StorageS3Bucket string `aws:" + `"Storage.S3.Bucket"` + "`"},
		}
		for _, c := range cases {
			prop := &Property{Name: c.Name, Description: c.Description, Type: c.Type}
			So(prop.toRequestTypeDefinition(), ShouldEqual, c.Response)
		}
	})
}

func TestPropertiesToTypeDefinition(t *testing.T) {
	Convey("toTypeDefinition", t, func() {
		tests := []*toTypeFieldTest{
			{
				Name:        "requestId",
				Description: "The ID of the request.",
				Type:        "xsd:string",
				Response:    "RequestId string `xml:" + `"requestId,omitempty"` + "`",
			},
			{
				Description: "The root device name (for example, /dev/sda1).",
				Type:        "String",
				Name:        "rootDeviceName",
				Response:    "RootDeviceName string `xml:" + `"rootDeviceName,omitempty"` + "`",
			},
			{
				Name:        "noDevice",
				Description: "Include this empty element to suppress the specified device included in the block device mapping of the AMI.",
				Response:    "NoDevice struct{} `xml:" + `"noDevice,omitempty"` + "`",
			},
			{
				Name:        "attachTime",
				Description: "The time stamp when the attachment initiated.",
				Type:        "DateTime",
				Response:    "AttachTime time.Time `xml:" + `"attachTime,omitempty"` + "`",
			},
			{
				Name:        "pricingDetailsSet",
				Description: "The pricing details of the Reserved Instance offering wrapped in an item element.",
				Type:        "PricingDetailsSetItemType.",
				Response:    "PricingDetails []*PricingDetailsSetItem `xml:" + `"pricingDetailsSet>item,omitempty"` + "`",
			},
			{
				Name:        "groupSet.item",
				Description: "A security group.",
				Type:        "GroupItemType",
				Response:    "Groups []*GroupItem `xml:" + `"groupSet>item,omitempty"` + "`",
			},
			{
				Name:        "instanceCounts",
				Description: "The number of instances in this state.",
				Type:        "InstanceCountsSetType",
				Response:    "InstanceCounts []*InstanceCountsSet `xml:" + `"instanceCounts,omitempty"` + "`",
			},
			{
				Name:        "instanceState",
				Description: "The current state of the instance.",
				Type:        "InstanceStateType",
				Response:    "InstanceState *InstanceState `xml:" + `"instanceState,omitempty"` + "`",
			},
			{
				Name:        "networkInterfaceSet",
				Description: "[EC2-VPC] One or more network interfaces for the instance.",
				Type:        "InstanceNetworkInterfaceSetItemType",
				Response:    "NetworkInterfaces []*InstanceNetworkInterfaceSetItem `xml:" + `"networkInterfaceSet>item,omitempty"` + "`",
			},
			{
				Description: "Any block device mapping entries for the instance, each one wrapped in an item element.",
				Type:        "InstanceBlockDeviceMappingResponseItemType",
				Name:        "blockDeviceMapping",
				Response:    "BlockDeviceMappings []*InstanceBlockDeviceMappingResponseItem `xml:" + `"blockDeviceMapping>item,omitempty"` + "`",
			},
		}
		for _, t := range tests {
			prop := &Property{Name: t.Name, Description: t.Description, Type: t.Type}
			So(prop.toTypeDefinition().String(), ShouldEqual, t.Response)
		}
	})
}
