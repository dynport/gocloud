package cloudformation

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dynport/gocloud/aws"
)

type Client struct {
	*aws.Client
}

func NewFromEnv() *Client {
	return &Client{Client: aws.NewFromEnv()}
}

func (client *Client) loadCloudFormationResource(action string, params url.Values, i interface{}) error {
	req, e := client.signedCloudFormationRequest(action, params)

	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	switch rsp.StatusCode {
	case 400, 404:
		return ErrorNotFound
	case 200:
		return xml.Unmarshal(b, i)
	default:
		return fmt.Errorf("expected status 2xx but got %s (%s)", rsp.Status, string(b))
	}
}

func defaultParams() url.Values {
	return url.Values{
		"Version": {"2010-05-15"},
	}
}

var (
	ErrorNotFound = fmt.Errorf("Error not found")
	urlRoot       = "https://cloudformation.eu-west-1.amazonaws.com/"
)

func (client *Client) signedCloudFormationRequest(action string, params url.Values) (*http.Request, error) {
	values := defaultParams()
	for k, vs := range params {
		for _, v := range vs {
			values.Add(k, v)
		}
	}
	values.Add("Action", action)
	theUrl := urlRoot + "?" + values.Encode()
	req, e := http.NewRequest("GET", theUrl, nil)
	if e != nil {
		return nil, e
	}
	client.SignAwsRequestV2(req, time.Now())
	return req, nil
}
