package cloudformation

import (
	"encoding/xml"
	"strconv"
)

type EstimateTemplateCostResponse struct {
	XMLName                    xml.Name                    `xml:"EstimateTemplateCostResponse"`
	EstimateTemplateCostResult *EstimateTemplateCostResult `xml:"EstimateTemplateCostResult"`
}

type EstimateTemplateCostResult struct {
	Url string `xml:"Url"`
}

type EstimateTemplateCostParameters struct {
	TemplateBody string
	TemplateURL  string
	Parameters   []*StackParameter
}

func (c *Client) EstimateTemplateCost(params EstimateTemplateCostParameters) (*EstimateTemplateCostResponse, error) {
	r := &EstimateTemplateCostResponse{}
	v := Values{}
	if params.TemplateBody != "" {
		v["TemplateBody"] = params.TemplateBody
	}
	if params.TemplateURL != "" {
		v["TemplateURL"] = params.TemplateURL
	}
	for i, p := range params.Parameters {
		v["Parameters.member."+strconv.Itoa(i+1)+".ParameterKey"] = p.ParameterKey
		v["Parameters.member."+strconv.Itoa(i+1)+".ParameterValue"] = p.ParameterValue
	}
	e := c.loadCloudFormationResource("EstimateTemplateCost", v, r)
	return r, e
}
