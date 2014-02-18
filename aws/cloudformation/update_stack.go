package cloudformation

import "net/url"

type UpdateStackParameters struct {
	BaseParameters
	StackPolicyDuringUpdateBody string
	StackPolicyDuringUpdateURL  string
}

func (p *UpdateStackParameters) values() url.Values {
	v := p.BaseParameters.values()
	mapping := map[string]string{
		"StackPolicyDuringUpdateBody": p.StackPolicyDuringUpdateBody,
		"StackPolicyDuringUpdateURL":  p.StackPolicyDuringUpdateURL,
	}

	for k, value := range mapping {
		if value != "" {
			v.Add(k, value)
		}
	}
	return v
}

type UpdateStackResponse struct {
	UpdateStackResult *UpdateStackResult `xml:"UpdateStackResult"`
}

type UpdateStackResult struct {
	StackId string `xml:"StackId"`
}

func (c *Client) UpdateStack(params UpdateStackParameters) (stackId string, e error) {
	r := &UpdateStackResponse{}
	e = c.loadCloudFormationResource("UpdateStack", params.values(), r)
	if e != nil {
		return "", e
	}
	return r.UpdateStackResult.StackId, nil
}
