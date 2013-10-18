package cloudwatch

import (
	"encoding/xml"
	"github.com/dynport/gocloud/aws"
	"net/url"
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
	ENDPOINT = "https://monitoring.us-east-1.amazonaws.com"
	VERSION  = "2010-08-01"
)

type Client struct {
	*aws.Client
}

func (client *Client) ListMetrics() (rsp *ListMetricsResponse, e error) {
	values := &url.Values{}
	values.Add("Version", VERSION)
	values.Add("Action", "ListMetrics")
	raw, e := client.DoSignedRequest("GET", ENDPOINT, values.Encode(), nil)
	if e != nil {
		return nil, e
	}
	e = xml.Unmarshal(raw.Content, &rsp)
	return rsp, e
}
