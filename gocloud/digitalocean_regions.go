package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"os"
	"strconv"
)

func init() {
	router.Register("do/region/list", &gocli.Action{Description: "List available droplet regions", Handler: ListRegionsAction})
}

func ListRegionsAction(args *gocli.Args) error {
	logger.Debug("listing regions")
	account, e := AccountFromEnv()
	if e != nil {
		return e
	}
	logger.Debugf("account is %+v", account)
	table := gocli.NewTable()
	table.Add("Id", "Name")
	regions, e := account.Regions()
	if e != nil {
		return e
	}
	for _, region := range regions {
		table.Add(strconv.Itoa(region.Id), region.Name)
	}
	fmt.Fprintln(os.Stdout, table.String())
	return nil
}
