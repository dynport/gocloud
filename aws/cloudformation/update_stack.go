package cloudformation

type UpdateStackParameters struct {
	BaseParameters
	StackPolicyDuringUpdateBody string
	StackPolicyDuringUpdateURL  string
}

func (p *UpdateStackParameters) values() Values {
	v := p.BaseParameters.values()
	v["StackPolicyDuringUpdateBody"] = p.StackPolicyDuringUpdateBody
	v["StackPolicyDuringUpdateURL"] = p.StackPolicyDuringUpdateURL
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
