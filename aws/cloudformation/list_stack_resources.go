package cloudformation

import "encoding/xml"

type ListStackResourcesResponse struct {
	XMLName                          xml.Name                          `xml:"ListStackResourcesResponse"`
	InternalListStackResourcesResult *InternalListStackResourcesResult `xml:"ListStackResourcesResult"`
}

type InternalListStackResourcesResult struct {
	NextToken      *string          `xml:"NextToken"`
	StackResources []*StackResource `xml:"StackResourceSummaries>member"`
}

type ListStackResourcesResult struct {
	StackResources []*StackResource
}

type ListStackResourcesParameters struct {
	LogicalResourceId  string
	PhysicalResourceId string
	StackName          string
}

func (client *Client) ListStackResources(params ListStackResourcesParameters) (*ListStackResourcesResult, error) {
	var NextToken string
	result := &ListStackResourcesResult{
		StackResources: []*StackResource{},
	}
	for {
		r := &ListStackResourcesResponse{}
		values := Values{
			"StackName":          params.StackName,
			"PhysicalResourceId": params.PhysicalResourceId,
			"LogicalResourceId":  params.LogicalResourceId,
			"NextToken":          NextToken,
		}
		e := client.loadCloudFormationResource("ListStackResources", values, r)
		if e != nil {
			return nil, e
		}

		result.StackResources = append(result.StackResources, r.InternalListStackResourcesResult.StackResources...)
		if r.InternalListStackResourcesResult.NextToken == nil {
			break
		} else {
			NextToken = *r.InternalListStackResourcesResult.NextToken
		}
	}
	return result, nil
}
