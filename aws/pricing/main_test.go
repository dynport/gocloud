package pricing

import (
	"io/ioutil"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	_ "launchpad.net/xmlpath"
)

func mustReadFile(t *testing.T, path string) []byte {
	b, e := ioutil.ReadFile(path)
	if e != nil {
		t.Fatal(e.Error())
	}
	return b
}

func TestLoadPricing(t *testing.T) {
	b := mustReadFile(t, "fixtures/linux-od.json")
	Convey("Instance configs", t, func() {
		configs, e := AllInstanceTypeConfigs()
		So(e, ShouldBeNil)
		for _, c := range configs {
			switch c.Name {
			case "m3.large":
				So(c.Cpus, ShouldEqual, 2)
			}
		}
	})
	Convey("marshall pricing", t, func() {
		pricing, e := LoadPricing(b)
		So(e, ShouldBeNil)
		So(pricing, ShouldNotBeNil)
		So(pricing.Config, ShouldNotBeNil)
		So(len(pricing.Config.Regions), ShouldEqual, 8)

		region := pricing.Config.Regions[0]
		So(region.Region, ShouldEqual, "us-east")
		So(len(region.InstanceTypes), ShouldEqual, 9)

		it := region.InstanceTypes[0]
		So(it.Type, ShouldEqual, "generalCurrentGen")
		So(len(it.Sizes), ShouldEqual, 2)

		size := it.Sizes[0]
		So(size.Size, ShouldEqual, "m3.xlarge")
		So(len(size.ValueColumns), ShouldEqual, 1)
		vc := size.ValueColumns[0]
		So(vc.Name, ShouldEqual, "linux")
		So(vc.Prices["USD"], ShouldEqual, "0.450")
	})
}

// {
//       "name": "yrTerm1",
//       "prices": {
//         "USD": "338"
//       }
//     },
//     {
//       "name": "yrTerm1Hourly",
//       "rate": "perhr",
//       "prices": {
//         "USD": "0.028"
//       }
//     },
//     {
//       "name": "yrTerm3",
//       "prices": {
//         "USD": "514"
//       }
//     },
//     {
//       "name": "yrTerm3Hourly",
//       "rate": "perhr",
//       "prices": {
//         "USD": "0.023"
//       }
//     }

func TestValueColumnes(t *testing.T) {
	Convey("Value Columns", t, func() {
		Convey("on demand instances", func() {
			vcs := ValueColumns{
				{Name: "linux", Prices: map[string]string{"USD": "0.450"}},
			}
			So(vcs, ShouldNotBeNil)
			So(len(vcs.Prices()), ShouldEqual, 1)

			price := vcs.Prices()[0]
			So(price.PerHour, ShouldEqual, 0.45)
			So(price.TotalPerHour(), ShouldEqual, 0.45)
			So(price.Upfront, ShouldEqual, 0)

		})
		Convey("reserved instances", func() {
			vcs := ValueColumns{
				{Name: "yrTerm1Hourly", Rate: "perhr", Prices: map[string]string{"USD": "0.028"}},
				{Name: "yrTerm1", Prices: map[string]string{"USD": "338"}},
				{Name: "yrTerm3Hourly", Rate: "perhr", Prices: map[string]string{"USD": "0.023"}},
				{Name: "yrTerm3", Prices: map[string]string{"USD": "514"}},
			}
			So(vcs, ShouldNotBeNil)
			So(len(vcs.Prices()), ShouldEqual, 2)

			price := vcs.Prices()[0]
			So(price.Upfront, ShouldEqual, 338)
			So(price.PerHour, ShouldEqual, 0.028)
			So(price.TotalPerHour(), ShouldBeBetween, 0.06658, 0.06659)

			price = vcs.Prices()[1]
			So(price.Upfront, ShouldEqual, 514)
			So(price.PerHour, ShouldEqual, 0.023)
			So(price.TotalPerHour(), ShouldBeBetween, 0.042558, 0.042559)
		})
	})
}
