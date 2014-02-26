package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertTypes(t *testing.T) {
	Convey("Convert Types", t, func() {
		type testCase struct {
			Raw        string
			Normalized string
		}
		tests := []*testCase{
			{"Integer (16-bit unsigned)", "int"},
			{"String", "string"},
			{"xsd:string", "string"},
			{"Double", "float64"},
			{"dateTime", "time.Time"},
		}
		for _, tc := range tests {
			So(convertType(tc.Raw), ShouldEqual, tc.Normalized)
		}
	})
}

func TestTypes(t *testing.T) {
	Convey("Types", t, func() {
		type testCase struct {
			Name       string
			Type       string
			Properties []*Property
			Matches    []string
		}
		cases := []*testCase{
			{
				Name: "ReservationInfoType",
				Type: "ReservationInfo",
				Properties: []*Property{
					{Name: "reservationId", Type: "String"},
					{Name: "ownerId", Type: "String"},
					{Name: "groupSet", Type: "GroupItemType"},
				},
				Matches: []string{
					"type ReservationInfo struct {\n\tReservationId string `xml:" + `"reservationId,omitempty"` + "`\n",
				},
			},
			{
				Name: "InstanceCountsSetItemType",
				Type: "InstanceCountsSetItem",
			},
		}
		for _, c := range cases {
			fields := []*TypeField{}
			for _, p := range c.Properties {
				fields = append(fields, p.toTypeDefinition())
			}
			typ := &Type{
				Name:   c.Name,
				Fields: fields,
			}
			s := typ.String()
			So(s, ShouldStartWith, "type "+c.Type+" struct {\n")
			So(s, ShouldEndWith, "\n}\n")
			So(typ.Type(), ShouldEqual, c.Type)
			for _, m := range c.Matches {
				So(s, ShouldContainSubstring, m)
			}
		}
	})
}
