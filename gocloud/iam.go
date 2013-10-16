package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/iam"
	"strings"
)

func init() {
	router.Register("aws/iam/users/get", &gocli.Action{
		Handler: iamGetUser, Description: "Get user information",
	})

	router.Register("aws/iam/users/list", &gocli.Action{
		Handler: iamListUsers, Description: "List users",
	})

	router.Register("aws/iam/account-summary", &gocli.Action{
		Handler: iamGetAccountSummary, Description: "Get account summary",
	})

	router.Register("aws/iam/account-aliases/list", &gocli.Action{
		Handler: iamListAccountAliases, Description: "List account aliases",
	})
}

func iamGetUser(args *gocli.Args) error {
	client := iam.NewFromEnv()
	user, e := client.GetUser("")
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Id", user.UserId)
	table.Add("Name", user.UserName)
	table.Add("Arn", strings.TrimSpace(user.Arn))
	table.Add("Path", user.Path)
	fmt.Println(table)
	return nil
}

func iamGetAccountSummary(args *gocli.Args) error {
	client := iam.NewFromEnv()
	summary, e := client.GetAccountSummary()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, entry := range summary.Entries {
		table.Add(entry.Key, entry.Value)
	}
	fmt.Println(table)
	return nil
}

func iamListUsers(args *gocli.Args) error {
	client := iam.NewFromEnv()
	rsp, e := client.ListUsers()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, user := range rsp.Users {
		table.Add(user.UserId, user.UserName, strings.TrimSpace(user.Arn))
	}
	fmt.Println(table)
	return nil
}

func iamListAccountAliases(args *gocli.Args) error {
	client := iam.NewFromEnv()
	rsp, e := client.ListAccountAliases()
	if e != nil {
		return e
	}
	for _, alias := range rsp.AccountAliases {
		fmt.Println(alias)
	}
	return nil
}
