package cloudformation

import (
	"net/url"
	"strconv"
)

type BaseParameters struct {
	Capabilities    []string
	Parameters      []*StackParameter
	StackName       string
	StackPolicyBody string
	StackPolicyURL  string
	TemplateBody    string
	TemplateURL     string
}

func (c *BaseParameters) values() url.Values {
	v := url.Values{}
	for i, c := range c.Capabilities {
		v.Add("Capabilities.member."+strconv.Itoa(i+1), c)
	}
	for i, p := range c.Parameters {
		v.Add("Parameters.member."+strconv.Itoa(i+1)+".ParameterKey", p.ParameterKey)
		v.Add("Parameters.member."+strconv.Itoa(i+1)+".ParameterValue", p.ParameterValue)
	}

	mapping := map[string]string{
		"StackPolicyBody": c.StackPolicyBody,
		"StackPolicyURL":  c.StackPolicyURL,
		"TemplateBody":    c.TemplateBody,
		"TemplateURL":     c.TemplateURL,
		"StackName":       c.StackName,
	}

	for k, value := range mapping {
		if value != "" {
			v.Add(k, value)
		}
	}

	return v
}

type CreateStackParameters struct {
	BaseParameters
	DisableRollback  bool
	NotificationARNs []string
	OnFailure        string
	Tags             []*Tag
	TimeoutInMinutes int
}

func (c *CreateStackParameters) values() url.Values {
	v := c.BaseParameters.values()
	if c.DisableRollback {
		v.Add("DisableRollback", "true")
	}

	if c.OnFailure != "" {
		v.Add("OnFailure", c.OnFailure)
	}

	if c.TimeoutInMinutes > 0 {
		v.Add("TimeoutInMinutes", strconv.Itoa(c.TimeoutInMinutes))
	}

	for i, arn := range c.NotificationARNs {
		v.Add("NoNotificationARNs.member."+strconv.Itoa(i+1), arn)
	}
	return v
}

type StackParameter struct {
	ParameterKey   string
	ParameterValue string
}

type CreateStackResponse struct {
	CreateStackResult *CreateStackResult `xml:"CreateStackResult"`
}

type CreateStackResult struct {
	StackId string `xml:"StackId"`
}

func (client *Client) CreateStack(params CreateStackParameters) (stackId string, e error) {
	r := &CreateStackResponse{}
	v := params.values()
	e = client.loadCloudFormationResource("CreateStack", v, r)
	if e != nil {
		return "", e
	}
	return r.CreateStackResult.StackId, nil
}
