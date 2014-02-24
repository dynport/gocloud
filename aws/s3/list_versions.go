package s3

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type ListVersionsResult struct {
	Name            string            `xml:"Name"`            // example-bucket</Name>
	Prefix          string            `xml:"Prefix"`          // photos/2006/</Prefix>
	KeyMarker       string            `xml:"KeyMarker"`       // </KeyMarker>
	VersionIdMarker string            `xml:"VersionIdMarker"` // </VersionIdMarker>
	MaxKeys         int               `xml:"MaxKeys"`         // 1000</MaxKeys>
	Delimiter       string            `xml:"Delimiter"`       // /</Delimiter>
	IsTruncated     bool              `xml:"IsTruncated"`     // false</IsTruncated>
	CommonPrefixes  []*CommonPrefixes `xml:"CommonPrefixes"`
	Versions        []*Version        `xml:"Version"`
}

type Owner struct {
	ID          string `xml:"ID"`
	DisplayName string `xml:"DisplayName"`
}

type CommonPrefixes struct {
	Prefix string `xml:"Prefix"`
}

type Version struct {
	Key          string `xml:"Key"`          // photos/2006/</Key>
	VersionId    string `xml:"VersionId"`    // 3U275dAA4gz8ZOqOPHtJCUOi60krpCdy</VersionId>
	IsLatest     string `xml:"IsLatest"`     // true</IsLatest>
	LastModified string `xml:"LastModified"` // 2011-02-02T18:47:27.000Z</LastModified>
	ETag         string `xml:"ETag"`         // &quot;d41d8cd98f00b204e9800998ecf8427e&quot;</ETag>
	Size         string `xml:"Size"`         // 0</Size>
	Owner        *Owner `xml:"Owner"`
	StorageClass string `xml:"StorageClass"` // STANDARD</StorageClass>
}

func (client *Client) ListVersions(bucket string) (*ListVersionsResult, error) {
	req, e := http.NewRequest("GET", "https://"+client.EndpointHost()+"/"+bucket+"?versions", nil)
	if e != nil {
		return nil, e
	}
	req.Header.Add("Host", bucket+"."+client.EndpointHost())
	client.SignS3Request(req, bucket)
	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return nil, e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return nil, e
	}
	r := &ListVersionsResult{}
	e = xml.Unmarshal(b, r)
	return r, e
}
