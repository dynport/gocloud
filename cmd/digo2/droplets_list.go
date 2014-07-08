package main

import (
	"fmt"

	"github.com/dynport/gocli"
)

type dropletsList struct {
}

func (r *dropletsList) Run() error {
	cl, e := client()
	if e != nil {
		return e
	}
	rsp, e := cl.Droplets()
	if e != nil {
		return e
	}
	t := gocli.NewTable()
	for _, d := range rsp.Droplets {
		imageName := fmt.Sprintf("%d: %s", d.Image.Id, d.Image.Name)
		if d.Image.Slug != "" {
			imageName += " (" + d.Image.Slug + ")"
		}
		t.Add(d.Id, d.Status, d.CreatedAt.Format("2006-01-02T15:04:05"), d.Name, d.Region.Name, d.Size.Slug, imageName)
	}
	fmt.Println(t)
	return nil
}
