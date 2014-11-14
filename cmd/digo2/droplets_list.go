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
	t.Add("Id", "Status", "IP", "Private IP", "Name", "Region", "Size", "ImageId:ImageName (ImageSlug)", "CreatedAt")
	for _, d := range rsp.Droplets {
		imageName := fmt.Sprintf("%d:%s", d.Image.Id, d.Image.Name)
		if d.Image.Slug != "" {
			imageName += " (" + d.Image.Slug + ")"
		}
		var public, private string
		if d.Networks != nil {
			for _, i := range d.Networks.V4 {
				switch i.Type {
				case "public":
					public = i.IpAddress
				case "private":
					private = i.IpAddress

				}
			}
		}
		reg := func() string {
			if d.Region != nil {
				return d.Region.Slug
			}
			return ""
		}()
		created := func() string {
			if !d.CreatedAt.IsZero() {
				return d.CreatedAt.Format("2006-01-02 15:04:05")
			}
			return ""
		}()
		size := func() string {
			if d.Size != nil {
				return d.Size.Slug
			}
			return d.SizeSlug
		}()
		t.Add(d.Id, d.Status, public, private, d.Name, reg, size, imageName, created)
	}
	fmt.Println(t)
	return nil
}
