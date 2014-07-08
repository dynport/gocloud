package main

import (
	"fmt"
	"strings"

	"github.com/dynport/gocli"
)

type imagesList struct {
}

func (r *imagesList) Run() error {
	cl, e := client()
	if e != nil {
		return e
	}
	rsp, e := cl.Images()
	if e != nil {
		return e
	}
	t := gocli.NewTable()
	for _, i := range rsp.Images {
		t.Add(i.Id, i.Slug, i.Name, strings.Join(i.Regions, ","))
	}
	fmt.Println(t)
	return nil
}
