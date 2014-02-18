package cloudformation

import "net/url"

func (c *Client) DeleteStack(name string) error {
	return c.loadCloudFormationResource("DeleteStack", url.Values{"StackName": {name}}, nil)
}
