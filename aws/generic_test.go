package aws

import (
	"testing"

	"github.com/dynport/dgtk/expect"
)

type dummyAction struct {
	Version         Version `aws:"2011-06-15"`
	Action          Action  `aws:"GetSessionToken"`
	DurationSeconds int     `aws:"DurationSeconds"`
	SerialNumber    string  `aws:"SerialNumber"`
	TokenCode       string  `aws:"TokenCode"`
}

func TestUrlForAction(t *testing.T) {
	expect := expect.New(t)
	pr := &dummyAction{DurationSeconds: 3601}
	expect(pr).ToNotBeNil()

	params, err := ParamsForAction(pr)
	expect(err).ToBeNil()
	expect(params).ToNotBeNil()
	expect(params["DurationSeconds"]).ToContain("3601")
	expect(params["Version"]).ToContain("2011-06-15")
	expect(params["Action"]).ToContain("GetSessionToken")

}
