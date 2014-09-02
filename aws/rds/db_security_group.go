package rds

import "encoding/xml"

type CreateDBSecurityGroup struct {
	DBSecurityGroupName        string `xml:"DBSecurityGroupName"`
	DBSecurityGroupDescription string `xml:"DBSecurityGroupDescription"`
}

type CreateDBSecurityGroupResponse struct {
	XMLName xml.Name                     `xml:"CreateDBSecurityGroupResponse"`
	Result  *CreateDBSecurityGroupResult `xml:"CreateDBSecurityGroupResult"`
}

type CreateDBSecurityGroupResult struct {
	XMLName         xml.Name         `xml:"CreateDBSecurityGroupResult"`
	DBSecurityGroup *DBSecurityGroup `xml:"DBSecurityGroup"`
}

func (action *CreateDBSecurityGroup) Execute(client *Client) (res *CreateDBSecurityGroupResponse, e error) {
	v := newAction("CreateDBSecurityGroup")
	if e = loadValues(v, action); e != nil {
		return nil, e
	}

	res = &CreateDBSecurityGroupResponse{}
	return res, client.loadResource("GET", client.Endpoint()+"?"+v.Encode(), nil, res)
}

type AuthorizeDBSecurityGroupIngress struct {
	DBSecurityGroupName string `xml:"DBSecurityGroupName"`
	CIDRIP              string `xml:"CIDRIP"`
}

type AuthorizeDBSecurityGroupIngressResponse struct {
	XMLName xml.Name                        `xml:"AuthorizeDBSecurityGroupIngressResponse"`
	Result  *AuthorizeDBSecurityGroupResult `xml:"AuthorizeDBSecurityGroupIngressResult"`
}
type AuthorizeDBSecurityGroupResult struct {
	XMLName         xml.Name         `xml:"AuthorizeDBSecurityGroupIngressResult"`
	DBSecurityGroup *DBSecurityGroup `xml:"DBSecurityGroup"`
}

func (action *AuthorizeDBSecurityGroupIngress) Execute(client *Client) (res *AuthorizeDBSecurityGroupIngressResponse, e error) {
	v := newAction("AuthorizeDBSecurityGroupIngress")
	if e = loadValues(v, action); e != nil {
		return nil, e
	}

	res = &AuthorizeDBSecurityGroupIngressResponse{}
	return res, client.loadResource("GET", client.Endpoint()+"?"+v.Encode(), nil, res)
}

type DeleteDBSecurityGroup struct {
	DBSecurityGroupName string `xml:"DBSecurityGroupName"`
}

func (action *DeleteDBSecurityGroup) Execute(client *Client) (e error) {
	v := newAction("DeleteDBSecurityGroup")
	if e = loadValues(v, action); e != nil {
		return e
	}
	res := &struct {
		XMLName xml.Name `xml:"DeleteDBSecurityGroupResponse"`
	}{}
	return client.loadResource("GET", client.Endpoint()+"?"+v.Encode(), nil, res)
}
