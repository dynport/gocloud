package main

import (
	"net/http"

	"github.com/dynport/gocloud/aws/ec2"
)

type ImageFilter struct {
	Owner string
	Name  string
}

func listImagesFactory(filter *ImageFilter) func(w http.ResponseWriter) (int, []byte) {
	return func(w http.ResponseWriter) (int, []byte) {
		rsp, e := ec2Client.DescribeImagesWithFilter(
			&ec2.ImageFilter{Owner: filter.Owner, Name: filter.Name},
		)
		if e != nil {
			return returnError(e)
		}
		return renderJson(w, rsp)
	}
}
