package iam

import (
	"encoding/xml"
	"time"

	"github.com/dynport/gocloud/aws"
)

const (
	ENDPOINT    = "https://iam.amazonaws.com"
	API_VERSION = "2010-05-08"
)

type Client struct {
	*aws.Client
}

func NewFromEnv() *Client {
	return &Client{
		aws.NewFromEnv(),
	}
}

type GetUserResponse struct {
	User *User `xml:"GetUserResult>User"`
}

type Entry struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type SummaryMap struct {
	Entries []*Entry `xml:"entry"`
}

type GetAccountSummaryResponse struct {
	SummaryMap *SummaryMap `xml:"GetAccountSummaryResult>SummaryMap"`
}

type User struct {
	Path     string `xml:"Path"`
	UserName string `xml:"UserName"`
	UserId   string `xml:"UserId"`
	Arn      string `xml:"Arn"`
}

type Role struct {
	Path                     string `xml:"Path"`                     ///application_abc/component_xyz/</Path>
	Arn                      string `xml:"Arn"`                      //arn:aws:iam::123456789012:role/application_abc/component_xyz/S3Access</Arn>
	RoleName                 string `xml:"RoleName"`                 //S3Access</RoleName>
	AssumeRolePolicyDocument string `xml:"AssumeRolePolicyDocument"` //{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":["ec2.amazonaws.com"]},"Action":["sts:AssumeRole"]}]}</AssumeRolePolicyDocument>
	CreateDate               string `xml:"CreateDate"`               //2012-05-09T15:45:35Z</CreateDate>
	RoleId                   string `xml:"RoleId"`                   //AROACVSVTSZYEXAMPLEYK</RoleId>
}

func (client *Client) GetUser(userName string) (user *User, e error) {
	raw, e := client.DoSignedRequest("GET", ENDPOINT, aws.QueryPrefix(API_VERSION, "GetUser"), nil)
	if e != nil {
		return user, e
	}
	rsp := &GetUserResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return user, e
	}
	return rsp.User, nil
}

func (client *Client) GetAccountSummary() (m *SummaryMap, e error) {
	raw, e := client.DoSignedRequest("GET", ENDPOINT, aws.QueryPrefix(API_VERSION, "GetAccountSummary"), nil)
	if e != nil {
		return m, e
	}
	rsp := &GetAccountSummaryResponse{}
	if e := aws.ExtractError(raw.Content); e != nil {
		return nil, e
	}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return m, e
	}
	return rsp.SummaryMap, nil
}

type ListInstanceProfilesResponse struct {
	XMLName          xml.Name           `xml:"ListInstanceProfilesResponse"`
	InstanceProfiles []*InstanceProfile `xml:"ListInstanceProfilesResult>InstanceProfiles>member"`
}

type InstanceProfile struct {
	Id                  string    `xml:"Id"`                  //AIPACZLSXM2EYYEXAMPLE</Id>
	Roles               []*Role   `xml:"Roles>member"`        //
	InstanceProfileName string    `xml:"InstanceProfileName"` //Webserver</InstanceProfileName>
	Path                string    `xml:"Path"`                ///application_abc/component_xyz/</Path>
	Arn                 string    `xml:"Arn"`
	CreateDate          time.Time `xml:"CreateDate"`
}

type ListRolesResponse struct {
	XMLName xml.Name `xml:"ListRolesResponse"`
	Roles   []*Role  `xml:"ListRolesResult>Roles>member"`
}

type ListUsersResponse struct {
	Users []*User `xml:"ListUsersResult>Users>member"`
}

func (client *Client) ListUsers() (users *ListUsersResponse, e error) {
	raw, e := client.DoSignedRequest("GET", ENDPOINT, aws.QueryPrefix(API_VERSION, "ListUsers"), nil)
	if e != nil {
		return users, e
	}
	rsp := &ListUsersResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return users, e
	}
	return rsp, nil
}

type ListAccountAliasesResponse struct {
	AccountAliases []string `xml:"ListAccountAliasesResult>AccountAliases>member"`
	IsTruncated    bool     `ListAccountAliasesResult>IsTruncated`
}

func (client *Client) ListAccountAliases() (aliases *ListAccountAliasesResponse, e error) {
	raw, e := client.DoSignedRequest("GET", ENDPOINT, aws.QueryPrefix(API_VERSION, "ListAccountAliases"), nil)
	if e != nil {
		return aliases, e
	}
	rsp := &ListAccountAliasesResponse{}
	e = xml.Unmarshal(raw.Content, rsp)
	if e != nil {
		return rsp, e
	}
	return rsp, nil
}

func (client *Client) ListMFADevices(username string) (*ListMFADevicesResponse, error) {
	params := map[string]string{"UserName": username}
	raw, err := client.DoSignedRequest(
		"GET", ENDPOINT,
		aws.QueryPrefix(API_VERSION, "ListMFADevices"),
		params)
	if err != nil {
		return nil, err
	}
	rsp := &ListMFADevicesResponse{}
	err = xml.Unmarshal(raw.Content, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type ListMFADevicesResponse struct {
	MFADevices  []*MFADevice `xml:"ListMFADevicesResult>MFADevices>member"`
	IsTruncated bool         `ListAccountAliasesResult>IsTruncated`
}

type MFADevice struct {
	UserName     string `xml:"UserName"`
	SerialNumber string `xml:"SerialNumber"`
}
