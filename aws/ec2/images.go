package ec2

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

func (client *Client) DescribeImages() (images []*Image, e error) {
	return client.DescribeImagesWithFilter(&ImageFilter{})
}

func (client *Client) DescribeImagesWithFilter(filter *ImageFilter) (images ImageList, e error) {
	query := "Version=" + API_VERSIONS_EC2 + "&Action=DescribeImages"
	if filter.Owner != "" {
		query += "&Owner.1=" + filter.Owner
	}
	if filter.Name != "" {
		query += "&Filter.1.Name=name&Filter.1.Value.0=" + filter.Name
	}
	raw, e := client.DoSignedRequest("GET", ENDPOINT, query, nil)
	if e != nil {
		return
	}
	rsp := &DescribeImagesResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return images, e
	}
	return rsp.Images, nil
}

type CreateImageOptions struct {
	InstancId   string // required
	Name        string // required
	Description string
	NoReboot    bool
}

type CreateImageResponse struct {
	XMLName   xml.Name `xml:"CreateImageResponse"`
	RequestId string   `xml:"requestId"`
	ImageId   string   `xml:"imageId"`
}

func (client *Client) CreateImage(opts *CreateImageOptions) (rsp *CreateImageResponse, e error) {
	values := &url.Values{}
	values.Add("Version", API_VERSIONS_EC2)
	values.Add("Action", "CreateImage")
	values.Add("Name", opts.Name)
	if opts.Description != "" {
		values.Add("Description", opts.Description)
	}
	values.Add("NoReboot", fmt.Sprintf("%t", opts.NoReboot))

	raw, e := client.DoSignedRequest("GET", ENDPOINT, values.Encode(), nil)
	if e != nil {
		return nil, e
	}
	rsp = &CreateImageResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	return rsp, e
}
