package main

import (
	"net/http"
	"strings"

	"github.com/dynport/gocloud/aws/cloudformation"
	"github.com/dynport/martini"
)

func stacksResources(w http.ResponseWriter, params martini.Params) (int, []byte) {
	rsp, e := cfClient.DescribeStackResources(cloudformation.DescribeStackResourcesParameters{
		StackName: params["name"],
	})
	if e != nil {
		return returnError(e)
	}
	resources := []*StackResource{}
	for _, r := range rsp.DescribeStackResourcesResult.StackResources {
		resources = append(resources, NewStackResource(r))
	}
	return renderJson(w, resources)
}

func stacksShow(w http.ResponseWriter, params martini.Params) (int, []byte) {
	rsp, e := cfClient.DescribeStacks(&cloudformation.DescribeStacksParameters{
		StackName: params["name"],
	})
	if e != nil {
		returnError(e)
	}
	if len(rsp.DescribeStacksResult.Stacks) != 1 {
		return http.StatusNotFound, nil
	}
	raw := rsp.DescribeStacksResult.Stacks[0]
	return renderJson(w, &Stack{Name: raw.StackName, Original: raw, Status: raw.StackStatus})
}

func stacksList(w http.ResponseWriter) (int, []byte) {
	raw, e := cfClient.ListStacks(nil)
	if e != nil {
		return returnError(e)
	}
	stacks := []*Stack{}
	for _, s := range raw.ListStacksResult.Stacks {
		if strings.HasPrefix(s.StackStatus, "DELETE_COMPLETE") {
			continue
		}
		stacks = append(stacks, &Stack{Name: s.StackName, Status: s.StackStatus})
	}
	return renderJson(w, stacks)
}
