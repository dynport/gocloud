package main

import (
	"io/ioutil"
	"testing"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	. "github.com/smartystreets/goconvey/convey"
)

func mustRead(t *testing.T, p string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + p)
	if e != nil {
		t.Fatal(e.Error())
	}
	return b
}

func mapProperties(props []*Property) map[string]*Property {
	m := map[string]*Property{}
	for _, p := range props {
		m[p.Name] = p
	}
	return m
}

func mustParseDoc(t *testing.T, p string) xml.Node {
	doc, e := gokogiri.ParseHtml(mustRead(t, p))
	if e != nil {
		t.Fatal(e.Error())
	}
	return doc
}

func TestParseDocuNode(t *testing.T) {
	Convey("parseDocuNode", t, func() {
		Convey("parse ApiReference-ItemType-StateReasonType.html", func() {
			doc := mustParseDoc(t, "ApiReference-ItemType-StateReasonType.html")
			So(doc, ShouldNotBeNil)
			a, e := parseDocuNode(doc)
			So(e, ShouldBeNil)
			So(a, ShouldNotBeNil)
			last := a.AllProperties["Contents"][1]
			So(last.Name, ShouldEqual, "message")
			So(last.Type, ShouldEqual, "String")
		})
		Convey("NetworkInterfaceAttachment", func() {
			doc := mustParseDoc(t, "ApiReference-ItemType-InstanceNetworkInterfaceAttachmentType.html")
			So(doc, ShouldNotBeNil)
			a, e := parseDocuNode(doc)
			So(e, ShouldBeNil)
			So(a, ShouldNotBeNil)
			last := a.AllProperties["Contents"][3]
			So(last.Name, ShouldEqual, "attachTime")
			So(last.Type, ShouldEqual, "DateTime")
		})
	})

}

func TestParseTypes(t *testing.T) {
	Convey("Parse EC2 types", t, func() {
		doc, e := gokogiri.ParseHtml(mustRead(t, "API-ItemTypes.html"))
		So(e, ShouldBeNil)
		links, e := extractLinks(doc)
		So(e, ShouldBeNil)
		So(links, ShouldNotBeNil)
		So(len(links), ShouldEqual, 127)

		link := links[0]
		So(link.Name, ShouldEqual, "AccountAttributeSetItemType")
		So(link.Url, ShouldEqual, "http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AccountAttributeSetItemType.html")
	})
}

func TestParseType(t *testing.T) {
	Convey("Parse EC2 type", t, func() {
		doc, e := parseDocu(mustRead(t, "ApiReference-ItemType-AccountAttributeSetItemType.html"))
		So(e, ShouldBeNil)
		So(doc, ShouldNotBeNil)

		contents := doc.AllProperties["Contents"]

		So(len(contents), ShouldEqual, 2)

		c := contents[0]
		So(c.Name, ShouldEqual, "attributeName")
		So(c.Type, ShouldEqual, "String")
		So(c.Description, ShouldEqual, "The name of the attribute.")

		c2 := contents[1]
		So(c2.Name, ShouldEqual, "attributeValueSet")
		So(c2.Type, ShouldEqual, "AccountAttributeValueSetItemType")
	})
}

func TestParseReportInstanceStatus(t *testing.T) {
	Convey("Parse ReportInstanceStatus", t, func() {
		action, e := parseDocu(mustRead(t, "ApiReference-query-ReportInstanceStatus.html"))
		So(e, ShouldBeNil)
		So(action, ShouldNotBeNil)
		So(len(action.RequestParameters()), ShouldEqual, 6)
		cnt := 0
		for _, param := range action.RequestParameters() {
			switch param.Name {
			case "Status":
				So(param.Required, ShouldEqual, true)
				So(param.Type, ShouldEqual, "String")
				cnt++
			}
		}
		So(cnt, ShouldEqual, 1)
	})
}

func TestParseDocu(t *testing.T) {
	Convey("Parse AWS Docu", t, func() {
		Convey("Extract actions", func() {
			actions, e := parseActions(mustParseDoc(t, "query-apis.html"))
			So(e, ShouldBeNil)
			So(len(actions), ShouldEqual, 148)
			So(actions[0].DocumentationUrl, ShouldEqual, "http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AllocateAddress.html")
		})

		Convey("Parse DescribeSnapshots", func() {
			action, e := parseDocu(mustRead(t, "ApiReference-query-DescribeSnapshots.html"))
			So(e, ShouldBeNil)
			So(action, ShouldNotBeNil)
			So(len(action.RequestParameters()), ShouldEqual, 5)
		})

		Convey("Parse RunInstance", func() {
			action, e := parseDocu(mustRead(t, "ApiReference-query-RunInstances.html"))
			So(e, ShouldBeNil)
			So(action, ShouldNotBeNil)
			params := action.RequestParameters()
			So(len(params), ShouldEqual, 41)

			first := params[0]
			So(first.Name, ShouldEqual, "ImageId")

			checked := 0
			for _, param := range params {
				switch param.Name {
				case "ImageId":
					So(param.Required, ShouldEqual, true)
					So(param.Type, ShouldEqual, "String")
					checked++
				case "MaxCount":
					So(param.Required, ShouldEqual, true)
					checked++
				case "InstanceType":
					So(param.Type, ShouldEqual, "String")
					So(len(param.ValidValues), ShouldEqual, 29)
					So(param.ValidValues[1], ShouldEqual, "m1.medium")
					// implement me!
					//So(param.Default, ShouldEqual, "m1.small")
					checked++
				}
			}
			So(checked, ShouldEqual, 3)
		})

		Convey("Parse S3", func() {
			action, e := parseDocu(mustRead(t, "aws-properties-s3-bucket.html"))
			So(e, ShouldBeNil)
			properties := action.Properties()
			So(properties, ShouldNotBeNil)
			So(len(properties), ShouldEqual, 4)
			cnt := 0
			for _, prop := range properties {
				switch prop.Name {
				case "AccessControl":
					So(prop.Required, ShouldEqual, false)
					So(prop.Type, ShouldEqual, "String")
					So(len(prop.ValidValues), ShouldEqual, 7)
					So(prop.ValidValues[1], ShouldEqual, "PublicRead")
					cnt++
				}
			}
			So(cnt, ShouldEqual, 1)
		})

		Convey("Parse EC2 Instance", func() {
			action, e := parseDocu(mustRead(t, "aws-properties-ec2-instance.html"))
			So(e, ShouldBeNil)
			properties := action.Properties()
			So(properties, ShouldNotBeNil)
			So(len(properties), ShouldEqual, 22)

			propMap := mapProperties(properties)

			So(propMap["KeyName"].Description, ShouldEqual, "Provides the name of the Amazon EC2 key pair.")
			So(propMap["Tags"].Required, ShouldEqual, false)
			So(propMap["ImageId"].Required, ShouldEqual, true)
			So(propMap["InstanceType"].Type, ShouldEqual, "String")
			So(propMap["InstanceType"].List, ShouldEqual, false)
			So(propMap["BlockDeviceMappings"].Type, ShouldEqual, "EC2 Block Device Mapping")
			So(propMap["BlockDeviceMappings"].List, ShouldEqual, true)
		})
	})
}
