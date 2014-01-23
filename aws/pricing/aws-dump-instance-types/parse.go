package main

import (
	"fmt"
	"github.com/dynport/gocloud/aws/pricing"
	"github.com/moovweb/gokogiri"
	"strconv"
	"strings"
)

func parseInstanceTypes(b []byte) (instances []*pricing.InstanceTypeConfig, e error) {
	doc, e := gokogiri.ParseHtml(b)
	if e != nil {
		return nil, e
	}
	tables, e := doc.Search("//table")
	if e != nil {
		return nil, e
	}
	types := map[string]*pricing.InstanceTypeConfig{}
	for _, t := range tables {
		trs, e := t.Search(".//tr")
		if e != nil {
			return nil, e
		}
		headers := []string{}
		for _, tr := range trs {
			tds, e := tr.Search(".//td")
			if e != nil {
				return nil, e
			}
			if len(headers) == 0 {
				for _, td := range tds {
					headers = append(headers, td.Content())
				}
			} else {
				var instance *pricing.InstanceTypeConfig
				if len(tds) > 2 {
					family := tds[0].Content()
					instanceType := strings.TrimSpace(tds[1].Content())
					var ok bool
					instance, ok = types[instanceType]
					if !ok {
						instance = &pricing.InstanceTypeConfig{Family: family, Name: instanceType}
						types[instanceType] = instance
					}
				}
				for i, td := range tds[2:] {
					i += 2
					if i >= len(headers) {
						return nil, fmt.Errorf("tried to acces header %d but I only got %#v", i, headers)
					}
					header := headers[i]
					value := strings.TrimSpace(td.Content())
					switch header {
					case "Processor Arch":
						instance.Arch = strings.Replace(value, "\n", " ", -1)
					case "vCPU":
						instance.Cpus, _ = strconv.Atoi(value)
					case "ECU":
						instance.ECUs, e = strconv.ParseFloat(value, 64)
						if e != nil {
							instance.ECUText = value
						}
					case "Memory (GiB)":
						instance.Memory, _ = strconv.ParseFloat(value, 64)
					case "Intel® Turbo":
						instance.Turbo = value == "Yes"
					case "Intel® AVX†":
						instance.AVX = value == "Yes"
					case "Intel® AES-NI":
						instance.AES = value == "Yes"
					case "Physical Processor":
						instance.PhysicalProcessor = value
					case "Instance Storage (GB)":
						instance.Storage = value
					case "EBS-optimized Available":
						instance.EbsOptimizable = value == "Yes"
					case "Network Performance":
						instance.NetworkPerformance = value
					default:
						// should not happen
					}
				}
			}
		}
	}
	for _, instance := range types {
		instances = append(instances, instance)
	}
	return instances, nil
}
