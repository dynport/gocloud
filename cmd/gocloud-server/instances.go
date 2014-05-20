package main

import (
	"net/http"

	"github.com/dynport/gocloud/aws/ec2"
	"github.com/dynport/martini"
)

func instancesShow(w http.ResponseWriter, params martini.Params) (int, []byte) {
	orig, e := ec2Client.DescribeInstancesWithOptions(&ec2.DescribeInstancesOptions{
		InstanceIds: []string{params["id"]},
	})
	if e != nil {
		return returnError(e)
	}
	if len(orig) != 1 {
		return http.StatusNotFound, nil
	}
	return renderJson(w, NewInstance(orig[0]))
}

func instancesList(w http.ResponseWriter) (int, []byte) {
	orig, e := ec2Client.DescribeInstances()
	if e != nil {
		return returnError(e)
	}
	instances := []*Instance{}
	for _, i := range orig {
		if i.InstanceStateName == "running" {
			instances = append(instances, NewInstance(i))
		}
	}
	return renderJson(w, instances)
}
