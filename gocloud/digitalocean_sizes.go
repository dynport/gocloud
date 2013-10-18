package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"os"
	"strconv"
)

func init() {
	router.Register("do/size/list", &gocli.Action{Description: "List available droplet sizes", Handler: ListSizesAction})
}

func ListSizesAction(args *gocli.Args) error {
	logger.Debug("listing sizes")
	account, e := AccountFromEnv()
	if e != nil {
		return e
	}
	logger.Debugf("account is %+v", account)
	table := gocli.NewTable()
	table.Add("Id", "Name")
	sizes, e := account.Sizes()
	if e != nil {
		return e
	}
	for _, size := range sizes {
		table.Add(strconv.Itoa(size.Id), size.Name)
	}
	fmt.Fprintln(os.Stdout, table.String())
	return nil
}
