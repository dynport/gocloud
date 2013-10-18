package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws"
	"github.com/dynport/gocloud/aws/cloudwatch"
)

func init() {
	router.Register("aws/cw", &gocli.Action{
		Handler: cloudwatchList,
	})
}

func cloudwatchList(args *gocli.Args) error {
	client := cloudwatch.Client{Client: aws.NewFromEnv()}
	rsp, e := client.ListMetrics()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, m := range rsp.Metrics {
		table.Add(m.Namespace, m.MetricName)
		for _, d := range m.Dimensions {
			table.Add("", d.Name, d.Value)
		}
	}
	fmt.Println(table)
	return nil
}
