package cloudformation

import (
	"fmt"
	"time"

	"github.com/dynport/dgtk/cli"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/cloudformation"
)

type StacksList struct {
	IncludeDeleted bool `cli:"opt --deleted"`
}

var client = cloudformation.NewFromEnv()

func (list *StacksList) Run() error {
	rsp, e := client.ListStacks(nil)
	if e != nil {
		return e
	}
	t := gocli.NewTable()
	for _, s := range rsp.ListStacksResult.Stacks {
		if !list.IncludeDeleted && s.StackStatus == "DELETE_COMPLETE" {
			continue
		}
		t.Add(s.StackName, s.StackStatus, s.CreationTime.Format(time.RFC3339))
	}
	fmt.Println(t)
	return nil
}

type StackResources struct {
	Name string `cli:"arg required"`
}

func (r *StackResources) Run() error {
	rsp, e := client.DescribeStackResources(cloudformation.DescribeStackResourcesParameters{
		StackName: r.Name,
	})
	if e != nil {
		return e
	}
	t := gocli.NewTable()
	for _, r := range rsp.DescribeStackResourcesResult.StackResources {
		t.Add(r.LogicalResourceId, r.PhysicalResourceId)
	}
	fmt.Println(t)
	return nil

}

func Register(router *cli.Router) {
	router.Register("aws/cloudformation/stacks/list", &StacksList{}, "List Cloudformation stacks")
	router.Register("aws/cloudformation/stacks/resources", &StackResources{}, "Describe Stack Resources")
}
