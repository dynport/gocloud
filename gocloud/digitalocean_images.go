package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"os"
	"strconv"
)

func init() {
	router.RegisterFunc("do/image/list", ListImagesAction, "List available droplet images")
}

func ListImagesAction() error {
	logger.Debug("listing images")
	logger.Debugf("account is %+v", CurrentAccount())
	table := gocli.NewTable()
	table.Add("Id", "Name")
	images, e := account.Images()
	if e != nil {
		return e
	}
	for _, image := range images {
		table.Add(strconv.Itoa(image.Id), image.Name)
	}
	fmt.Fprintln(os.Stdout, table.String())
	return nil
}
