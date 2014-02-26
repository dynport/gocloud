package main

import (
	"fmt"
	"log"
	"os"
)

type Test struct {
}

var logger = log.New(os.Stderr, "", 0)

func (m *Test) Run() error {
	doc, e := loadDoc(ec2TypesUrl)
	if e != nil {
		return e
	}
	links, e := extractLinks(doc)
	if e != nil {
		return e
	}
	if len(links) == 0 {
		return fmt.Errorf("no links found")
	}
	for _, link := range links {
		log.Printf("%s: %s", link.Name, link.Url)
		linkDoc, e := loadDoc(link.Url)
		if e != nil {
			return e
		}
		a, e := parseDocuNode(linkDoc)
		if e != nil {
			return e
		}
		for k, v := range a.AllProperties {
			log.Printf("%s", k)
			log.Print(v)
		}
		break
	}
	return nil
}

func init() {
	router.Register("test", &Test{}, "Test")
}
