package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCastInputs(t *testing.T) {
	Convey("Cast Inputs", t, func() {
		inputs := Input{
			"DryRun": []interface{}{":bool"},
			"InstanceId": []interface{}{
				map[string][]interface{}{":list": []interface{}{":string"}},
				map[string]string{":rename": "InstanceIds"},
			},
		}
		fields := inputs.Fields()
		So(len(fields), ShouldEqual, 2)
	})
}

func TestYaml(t *testing.T) {
	Convey("Test YAML", t, func() {
		b := mustRead(t, "EC2-2013-10-15.yml")
		doc := &YamlDoc{}
		So(doc.Load(b), ShouldBeNil)
		So(doc.ApiVersion, ShouldEqual, "2013-10-15")
		So(len(doc.Operations), ShouldEqual, 148)

		var op *Operation

		for _, c := range doc.Operations {
			if c.Name == "DescribeInstances" {
				op = c
				break
			}
		}
		So(op, ShouldNotBeNil)
		fields := op.YamlTypes()
		So(len(fields), ShouldEqual, 5)

		if len(fields) != 5 {
			t.Fatalf("types must == 5 but are %d", len(fields))
		}

		So(fields[0].Name, ShouldEqual, "DryRun")
		So(fields[0].Type, ShouldEqual, "boolean")
		So(fields[0].List, ShouldEqual, false)

		So(fields[1].Name, ShouldEqual, "InstanceId")
		So(fields[1].Type, ShouldEqual, "string")
		So(fields[1].List, ShouldEqual, true)

		So(fields[2].Name, ShouldEqual, "Filter")
		So(fields[2].Struct, ShouldEqual, true)
		So(fields[2].List, ShouldEqual, true)
		So(len(fields[2].StructFields), ShouldEqual, 2)

		So(fields[2].StructFields[0].Name, ShouldEqual, "Name")
		So(fields[2].StructFields[0].Type, ShouldEqual, "string")
		So(fields[2].StructFields[0].List, ShouldEqual, false)

		So(fields[2].StructFields[1].Name, ShouldEqual, "Value")
		So(fields[2].StructFields[1].Type, ShouldEqual, "string")
		So(fields[2].StructFields[1].List, ShouldEqual, true)
	})
}
