package route53

import (
	"encoding/xml"
	"fmt"
	"github.com/dynport/gocloud/aws"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewFromEnv() *Client {
	return &Client{
		aws.NewFromEnv(),
	}
}

// http://docs.aws.amazon.com/Route53/latest/APIReference/Welcome.html
type Client struct {
	*aws.Client
}

type HostedZone struct {
	Id                     string `xml:"Id"`
	Name                   string `xml:"Name"`
	CallerReference        string `xml:"CallerReference"`
	ResourceRecordSetCount int    `xml:"ResourceRecordSetCount"`
}

func (zone *HostedZone) Code() string {
	chunks := strings.Split(zone.Id, "/")
	if len(chunks) == 3 {
		return chunks[2]
	}
	return ""
}

const API_VERSION = "2012-12-12"

type HttpResponse struct {
	StatusCode int
	Content    []byte
}

type ResourceRecord struct {
	Value string `xml:"Value"`
}

type ResourceRecordSet struct {
	Name            string            `xml:"Name"`
	Type            string            `xml:"Type"`
	TTL             int               `xml:"TTL"`
	HealthCheckId   string            `xml:"HealthCheckId"`
	SetIdentifier   string            `xml:"SetIdentifier"`
	Weight          int               `xml:"Weight"`
	ResourceRecords []*ResourceRecord `xml:"ResourceRecords>ResourceRecord"`
}

type ListResourceRecordSetsResponse struct {
	XMLName            xml.Name             `xml:"ListResourceRecordSetsResponse"`
	ResourceRecordSets []*ResourceRecordSet `xml:"ResourceRecordSets>ResourceRecordSet"`
}

func (client *Client) ListResourceRecordSets(id string) (rrs []*ResourceRecordSet, e error) {
	raw, e := client.doRequest("GET", "hostedzone/"+id+"/rrset")
	if e != nil {
		return nil, e
	}
	rsp := &ListResourceRecordSetsResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	return rsp.ResourceRecordSets, e
}

func (client *Client) GetHostedZone(id string) (zone *HostedZone, e error) {
	rsp, e := client.doRequest("GET", "hostedzone/"+id)
	if e != nil {
		return nil, e
	}
	fmt.Println(string(rsp.Content))
	return nil, nil
}

type GetHostedZoneResponse struct {
	XMLName     xml.Name   `xml:"GetHostedZoneResponse"`
	HostedZone  HostedZone `xml:"HostedZone"`
	NameServers []string   `xml:"DelegationSet>NameServers>NameServer"`
}

func (client *Client) ListHostedZones() (zones []*HostedZone, e error) {
	rsp, e := client.doRequest("GET", "hostedzone")
	if e != nil {
		return zones, e
	}
	zonesResponse := &ListHostedZonesResponse{}
	e = xml.Unmarshal(rsp.Content, zonesResponse)
	if e != nil {
		return zones, e
	}
	return zonesResponse.HostedZones, nil
}

type ListHostedZonesResponse struct {
	XMLName     xml.Name      `xml:"ListHostedZonesResponse"`
	IsTruncated bool          `xml:"IsTruncated"`
	MaxItems    int           `xml:"MaxItems"`
	HostedZones []*HostedZone `xml:"HostedZones>HostedZone"`
}

func (client *Client) doRequest(method, path string) (rsp *HttpResponse, e error) {
	request, e := http.NewRequest(method, "https://route53.amazonaws.com/"+API_VERSION+"/"+path, nil)
	if e != nil {
		return nil, e
	}
	client.SignAwsRequest(request)
	raw, e := http.DefaultClient.Do(request)

	if e != nil {
		return nil, e
	}
	defer raw.Body.Close()
	rsp = &HttpResponse{
		StatusCode: raw.StatusCode,
	}
	rsp.Content, e = ioutil.ReadAll(raw.Body)
	return rsp, e
}
