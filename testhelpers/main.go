package testhelpers

import (
	"io/ioutil"
	"testing"
)

func MustReadFixture(t *testing.T, name string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + name)
	if e != nil {
		t.Fatal("fixture " + name + " does not exist")
	}
	return b
}
