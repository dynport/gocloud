package sts

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dynport/gocloud/aws"

	"time"
)

type GetSessionToken struct {
	Version         aws.Version `aws:"2011-06-15"`
	Action          aws.Action  `aws:"GetSessionToken"`
	DurationSeconds int         `aws:"DurationSeconds"`
	SerialNumber    string      `aws:"SerialNumber"`
	TokenCode       string      `aws:"TokenCode"`
}

func (a *GetSessionToken) Execute(client *aws.Client) (*GetSessionTokenResponse, error) {
	params, err := aws.ParamsForAction(a)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", "https://sts.amazonaws.com/?"+params.Encode(), nil)
	tim := time.Now().UTC()
	client.SignAwsRequestV2(req, tim)
	for k, v := range req.Header {
		log.Printf("%s: %s", k, v)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.Status[0] != '2' {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("expected status 2xx, got %s. %s", rsp.Status, string(b))
	}
	var r *GetSessionTokenResponse
	return r, xml.NewDecoder(rsp.Body).Decode(&r)
}

type GetSessionTokenResponse struct {
	XMLName     xml.Name     `xml:"GetSessionTokenResponse"`
	Credentials *Credentials `xml:"GetSessionTokenResult>Credentials"`
}

type Credentials struct {
	AccessKeyID     string    `xml:"AccessKeyId,omitempty"`
	Expiration      time.Time `xml:",omitempty"`
	SecretAccessKey string    `xml:",omitempty"`
	SessionToken    string    `xml:",omitempty"`
}
