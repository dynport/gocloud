package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCloudformation(t *testing.T) {
	Convey("cloudformation", t, func() {
		type TestCase struct {
			In  string
			Out string
		}
		cases := []*TestCase{
			{In: "AWS::AutoScaling::AutoScalingGroup", Out: "AutoScalingAutoScalingGroup"},
			{In: "CloudFront CacheBehavior Type", Out: "CloudFrontCacheBehaviorType"},
		}
		for _, tc := range cases {
			So(normalizeCloudformationType(tc.In), ShouldEqual, tc.Out)
		}
	})
}
