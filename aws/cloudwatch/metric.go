package cloudwatch

import (
	"encoding/xml"
	"net/url"

	"github.com/dynport/gocloud/aws"
)

type Dimension struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type Metric struct {
	Dimensions []*Dimension `xml:"Dimensions>member"`
	MetricName string       `xml:"MetricName"`
	Namespace  string       `xml:"Namespace"`
}

type ListMetricsResponse struct {
	XMLName   xml.Name  `xml:"ListMetricsResponse"`
	Metrics   []*Metric `xml:"ListMetricsResult>Metrics>member"`
	NextToken string    `xml:"ListMetricsResult>NextToken"`
}

const (
	VERSION = "2010-08-01"
)

func endpoint(client *aws.Client) string {
	return "https://monitoring." + client.Region + ".amazonaws.com"
}

type Client struct {
	*aws.Client
}

func (client *Client) Endpoint() string {
	prefix := "https://monitoring"
	if client.Client.Region != "" {
		prefix += "." + client.Client.Region
	}
	return prefix + ".amazonaws.com"
}

type ListMetricsOptions struct {
	NextToken string
}

type ListMetricsOption func(*ListMetricsOptions)

func OptNextToken(i string) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.NextToken = i
	}
}

func (client *Client) ListMetrics(funcs ...ListMetricsOption) (rsp *ListMetricsResponse, e error) {
	values := &url.Values{}
	values.Add("Version", VERSION)
	values.Add("Action", "ListMetrics")
	o := &ListMetricsOptions{}
	for _, f := range funcs {
		f(o)
	}
	if o.NextToken != "" {
		values.Add("NextToken", o.NextToken)
	}
	logger.Printf("%#v", values.Encode())
	raw, e := client.DoSignedRequest("GET", client.Endpoint(), values.Encode(), nil)
	if e != nil {
		return nil, e
	}
	e = xml.Unmarshal(raw.Content, &rsp)
	return rsp, e
}
