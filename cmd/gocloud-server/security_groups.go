package main

import (
	"net/http"

	"github.com/dynport/gocloud/aws/ec2"
	"github.com/dynport/martini"
)

func securityGroupsShow(w http.ResponseWriter, params martini.Params) (int, []byte) {
	groups, e := ec2Client.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsParameters{
		GroupIds: []string{params["id"]},
	})
	if e != nil {
		return returnError(e)
	}
	if len(groups) != 1 {
		return http.StatusNotFound, nil
	}
	raw := groups[0]
	return renderJson(w, NewSecurityGroup(raw))
}

func securityGroupsInstances(w http.ResponseWriter, params martini.Params) (int, []byte) {
	rsp, e := ec2Client.DescribeInstances()
	if e != nil {
		return returnError(e)
	}
	instances := []*Instance{}
	for _, raw := range rsp {
		i := NewInstance(raw)
		if i.HasSecurityGroup(params["id"]) {
			instances = append(instances, NewInstance(raw))
		}
	}
	return renderJson(w, instances)
}

func securityGroupsList(w http.ResponseWriter) (int, []byte) {
	raw, e := ec2Client.DescribeSecurityGroups(nil)
	if e != nil {
		return returnError(e)
	}
	groups := make([]*SecurityGroup, 0, len(raw))
	for _, g := range raw {
		if g.GroupId != "" {
			groups = append(groups, NewSecurityGroup(g))
		}
	}
	return renderJson(w, groups)
}
