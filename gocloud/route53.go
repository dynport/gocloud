package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/route53"
)

const (
	USAGE_LIST_RRS = "HOSTEDZONE"
)

// route53
func init() {
	router.Register("aws/route53/hosted-zones/list", &gocli.Action{
		Handler: route53ListHostedZones, Description: "Describe load balancers",
	})

	router.Register("aws/route53/rrs/list", &gocli.Action{
		Handler: route53ListResourceRecordSet, Description: "Describe load balancers", Usage: USAGE_LIST_RRS,
	})
}

func route53ListHostedZones(args *gocli.Args) error {
	logger.Info("describing hosted zones")
	client := route53.NewFromEnv()
	zones, e := client.ListHostedZones()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Name", "record_set_count", "rrs count")
	for _, zone := range zones {
		table.Add(zone.Code(), zone.Name, zone.ResourceRecordSetCount)
	}
	fmt.Println(table)
	return nil
}

func route53ListResourceRecordSet(args *gocli.Args) error {
	if len(args.Args) != 1 {
		return fmt.Errorf(USAGE_LIST_RRS)
	}
	client := route53.NewFromEnv()
	rrsets, e := client.ListResourceRecordSets(args.Args[0])
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("name", "type", "ttl", "weight", "id", "hc id", "value")
	maxLen := 64
	for _, rrs := range rrsets {
		weight := ""
		if rrs.Weight > 0 {
			weight = fmt.Sprintf("%d", rrs.Weight)
		}
		col := []string{
			rrs.Name, rrs.Type, fmt.Sprintf("%d", rrs.TTL), rrs.SetIdentifier, weight, rrs.HealthCheckId,
		}
		for i, record := range rrs.ResourceRecords {
			v := record.Value
			if len(v) > maxLen {
				v = v[0:maxLen]
			}
			if i == 0 {
				col = append(col, v)
				table.AddStrings(col)
			} else {
				table.Add("", "", "", "", "", "", v)
			}
		}
	}
	fmt.Println(table)
	return nil
}
