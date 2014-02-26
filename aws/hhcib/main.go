package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/dynport/dgtk/cli"
	"github.com/dynport/gocli"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

var (
	cloudFormationRoot = "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html"
	ec2ApiRoot         = "http://docs.aws.amazon.com/AWSEC2/latest/APIReference/query-apis.html"
	ec2TypesUrl        = "http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API-ItemTypes.html"
	elbRoot            = "http://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_Operations.html"
	rdsRoot            = "http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Operations.html"
)

var router = cli.NewRouter()

func init() {
	//	router.Register("ec2", &Main{}, "Describe EC2 API")
	//	router.Register("scrape", &Scrape{}, "Scrape aws documentation")
}

func main() {
	e := router.RunWithArgs()
	if e != nil {
		switch e {
		case cli.ErrorHelpRequested, cli.ErrorNoRoute:
			// ignore
		default:
			log.Fatal(e.Error())
		}
	}
}

type Main struct {
}

func (main *Main) Run() error {
	log.Print("running")
	doc, e := loadDoc(ec2ApiRoot)
	if e != nil {
		return e
	}
	actions, e := extractLinks(doc)
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, a := range actions {
		log.Printf(a.Url)
		return nil
		doc, e := loadDoc(a.Url)
		if e != nil {
			return e
		}
		action, e := parseDocuNode(doc)
		if e != nil {
			return e
		}
		table.Add(a.Name)
		for _, param := range action.RequestParameters() {
			table.Add("", param.Name, param.Type, param.Required)
		}
	}
	fmt.Println(table)
	return nil
}

var customFields = map[string]*Type{}

//	"Filter": {
//		Name: "Filter",
//		Properties: []*Field{
//			{Name: "Name", Type: "string"},
//			{Name: "Values", Type: "[]string", AttributeName: "Value.m"},
//		},
//	},
//	"ReservedInstanceLimitPriceType":              {name: "ReservedInstanceLimitPriceType"},
//	"ReservedInstancesConfigurationSetItemType":   {name: "ReservedInstancesConfigurationSetItemType"},
//	"AssignPrivateIpAddressesSetItemRequestType":  {name: "AssignPrivateIpAddressesSetItemRequestType"},
//	"SecurityGroupIdSetItemType":                  {name: "SecurityGroupIdSetItemType"},
//	"PriceScheduleRequestSetItemType":             {name: "PriceScheduleRequestSetItemType"},
//	"DescribeReservedInstancesListingSetItemType": {name: "DescribeReservedInstancesListingSetItemType"},
//	"DescribeReservedInstancesSetItemType":        {name: "DescribeReservedInstancesSetItemType"},
//	"InstanceBlockDeviceMappingItemType":          {name: "InstanceBlockDeviceMappingItemType"},
//}

func loadDoc(theUrl string) (xml.Node, error) {
	rsp, e := http.Get(theUrl)
	if e != nil {
		return nil, e
	}
	defer rsp.Body.Close()
	if rsp.Status[0] != '2' {
		return nil, fmt.Errorf("expected response code 2xx, got %s", rsp.Status)
	}
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return nil, e
	}
	doc, e := gokogiri.ParseHtml(b)
	return doc, e
}

func init() {
	u, e := url.Parse("http://127.0.0.1:1234")
	if e != nil {
		log.Fatal(e.Error())
	}
	proxy := http.ProxyURL(u)
	http.DefaultClient.Transport = &http.Transport{Proxy: proxy}
}
