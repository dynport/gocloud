package sts

import (
	"encoding/xml"

	"os"
	"testing"
)

func TestGetSessionTokenParse(t *testing.T) {
	f, err := os.Open("fixtures/get_session_token_response.xml")
	if err != nil {
		t.Fatal("error opening fixture", err)
	}
	defer f.Close()
	var rsp *GetSessionTokenResponse
	err = xml.NewDecoder(f).Decode(&rsp)
	if err != nil {
		t.Fatal("error decoding response", err)
	}
	creds := rsp.Credentials
	tests := []struct {
		Name     string
		Value    interface{}
		Expected interface{}
	}{
		{"Expiration", creds.Expiration.Format("2006-01-02"), "2011-07-11"},
		{"AccessKeyID", creds.AccessKeyID, "AKIAIOSFODNN7EXAMPLE"},
	}

	for _, tst := range tests {
		if tst.Expected != tst.Value {
			t.Errorf("expected %s to be %#v, was %#v", tst.Name, tst.Expected, tst.Value)
		}
	}
}
