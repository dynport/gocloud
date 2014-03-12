package cloudformation

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dynport/gocloud/aws"
)

type Client struct {
	*aws.Client
}

func NewFromEnv() *Client {
	return &Client{Client: aws.NewFromEnv()}
}

func (client *Client) Endpoint() string {
	prefix := "https://cloudformation"
	if client.Client.Region != "" {
		prefix += "." + client.Client.Region
	}
	return prefix + ".amazonaws.com"
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
	case 404:
		return ErrorNotFound
	case 200:
		if i != nil {
			return xml.Unmarshal(b, i)
		}
		return nil
	default:
		ersp := &ErrorResponse{}
		e = xml.Unmarshal(b, ersp)
		if e != nil {
			return fmt.Errorf("expected status 2xx but got %s (%s)", rsp.Status, string(b))

		}
		if strings.Contains(ersp.Error.Message, "does not exist") {
			return ErrorNotFound
		}
		return fmt.Errorf(ersp.Error.Message)
	}
}

func defaultParams() url.Values {
	return url.Values{
		"Version": {"2010-05-15"},
	}
}

var (
	ErrorNotFound = fmt.Errorf("Error not found")
)

func (client *Client) signedCloudFormationRequest(action string, params url.Values) (*http.Request, error) {
	values := defaultParams()
	for k, vs := range params {
		for _, v := range vs {
			values.Add(k, v)
		}
	}
	values.Add("Action", action)
	theUrl := client.Endpoint() + "?" + values.Encode()
	req, e := http.NewRequest("GET", theUrl, nil)
	if e != nil {
		return nil, e
	}
	client.SignAwsRequestV2(req, time.Now())
	return req, nil
}
