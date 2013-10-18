package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/hetzner"
)

func init() {
	router.Register("hetzner/servers/describe", &gocli.Action{Handler: describeServer, Description: "describe server"})
	router.Register("hetzner/servers/list", &gocli.Action{Handler: listServers, Description: "list servers"})
	router.Register("hetzner/servers/rename", &gocli.Action{Handler: renameServer, Description: "rename server"})
}

func describeServer(args *gocli.Args) error {
	account, e := hetzner.AccountFromEnv()
	if e != nil {
		return e
	}

	if len(args.Args) != 1 {
		return fmt.Errorf("<ip>")
	}
	ip := args.Args[0]
	server, e := account.LoadServer(ip)
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("IP", server.ServerIp)
	table.Add("Number", server.ServerNumber)
	table.Add("Name", server.ServerName)
	table.Add("Product", server.Product)
	table.Add("DataCenter", server.Dc)
	table.Add("Status", server.Status)
	table.Add("Reset", server.Reset)
	table.Add("Rescue", server.Rescue)
	table.Add("VNC", server.Vnc)
	fmt.Println(table)
	return nil
}

func listServers(args *gocli.Args) error {
	account, e := hetzner.AccountFromEnv()
	if e != nil {
		return e
	}
	servers, e := account.Servers()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Number", "Name", "Product", "DC", "Ip", "Status")
	for _, server := range servers {
		table.Add(server.ServerNumber, server.ServerName, server.Product, server.Dc, server.ServerIp, server.Status)
	}
	fmt.Println(table)
	return nil
}

func renameServer(args *gocli.Args) error {
	account, e := hetzner.AccountFromEnv()
	if e != nil {
		return e
	}
	if len(args.Args) != 2 {
		return fmt.Errorf("<ip> <new_name>")
	}
	ip, name := args.Args[0], args.Args[1]
	logger.Infof("renaming servers %s to %s", ip, name)
	return account.RenameServer(ip, name)
}
