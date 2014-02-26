package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

var (
	requiredRequestRegexp = regexp.MustCompile("Required: (Yes|No)")
	typeRegexp            = regexp.MustCompile("Type: (.*)")
	validValuesRegexp     = regexp.MustCompile("Valid values: (.*)")
)

func firstNode(doc xml.Node, i interface{}) (xml.Node, error) {
	nodes, e := doc.Search(i)
	if e != nil {
		return nil, e
	}
	if len(nodes) > 0 {
		return nodes[0], nil
	}
	return nil, nil
}

func parseDocu(b []byte) (*DocumentationPage, error) {
	doc, e := gokogiri.ParseHtml(b)
	if e != nil {
		return nil, e
	}
	return parseDocuNode(doc)
}

func urlRoot(theUrl string) string {
	parts := strings.Split(ec2ApiRoot, "/")
	if len(parts) > 1 {
		return strings.Join(parts[0:len(parts)-1], "/")
	}
	return ""
}

type Link struct {
	Name string
	Url  string
}

func extractLinks(doc xml.Node) (links []*Link, e error) {
	nodes, e := doc.Search("//div[@class='highlights']/.//a[starts-with(@href, 'ApiReference')]")
	if e != nil {
		return nil, e
	}
	root := urlRoot(ec2ApiRoot)
	for _, node := range nodes {
		links = append(links, &Link{Name: node.Content(), Url: root + "/" + node.Attr("href")})
	}
	return links, nil
}

func parseActions(doc xml.Node) (actions []*DocumentationPage, e error) {
	nodes, e := doc.Search("//div[@id='query-apis']/.//a[starts-with(@href, 'ApiReference')]")
	if e != nil {
		return nil, e
	}
	root := urlRoot(ec2ApiRoot)
	for _, node := range nodes {
		actions = append(actions, &DocumentationPage{Name: node.Content(), DocumentationUrl: root + "/" + node.Attr("href")})
	}
	return actions, nil
}

var normalizeTitleRegexp = regexp.MustCompile("[\\s]+")

func normalizeTitlePage(s string) string {
	return strings.TrimSpace(normalizeTitleRegexp.ReplaceAllString(s, " "))
}

func parseDocuNode(doc xml.Node) (*DocumentationPage, error) {
	action := &DocumentationPage{}
	if node, _ := firstNode(doc, "//h1[@class='topictitle']"); node != nil {
		action.Type = node.Content()
	}
	sections, e := doc.Search(".//div[@class='section']")
	if e != nil {
		return nil, e
	}

	action.AllProperties = map[string][]*Property{}

	for _, section := range sections {
		node, e := firstNode(section, ".//div[@class='titlepage']")
		if e != nil || node == nil {
			continue
		}
		list, e := firstNode(section, ".//div[@class='variablelist']")
		if e != nil || list == nil {
			continue
		}

		title := normalizeTitlePage(node.Content())
		if action.AllProperties[title] == nil {
			action.AllProperties[title] = []*Property{}
		}
		props, e := extractProperties(list)
		if e != nil {
			return nil, e
		}
		action.AllProperties[title] = props
	}
	return action, nil
}

func extractProperties(list xml.Node) (properties []*Property, e error) {
	dts, e := list.Search("./dl/dt")
	if e != nil {
		return nil, e
	}
	dds, e := list.Search("./dl/dd")
	if e != nil {
		return nil, e
	}
	for i, dt := range dts {
		property := &Property{}
		properties = append(properties, property)
		property.Name = dt.Content()
		if i < len(dds) {
			dd := dds[i]
			ps, e := dd.Search(".//p")
			if e != nil {
				return nil, e
			}
			for i, p := range ps {
				value := p.Content()
				if i == 0 {
					property.Description = value
					continue
				}
				if required := findFirst(value, requiredRequestRegexp); required != "" {
					property.Required = required == "Yes"
				} else if typ := findFirst(value, typeRegexp); typ != "" {
					if strings.HasPrefix(typ, "A list of ") {
						property.List = true
						property.Type = strings.TrimSuffix(strings.TrimPrefix(typ, "A list of "), ".")
					} else {
						property.Type = typ
					}
				} else if strings.HasPrefix(value, "Valid values: ") {
					for _, v := range strings.Split(strings.TrimPrefix(value, "Valid values: "), "|") {
						property.ValidValues = append(property.ValidValues, strings.TrimSpace(v))
					}
				}
			}
		}
	}
	return properties, nil
}

func findFirst(value string, rx *regexp.Regexp) string {
	m := rx.FindStringSubmatch(value)
	if len(m) > 1 {
		return m[1]
	}
	return ""

}

func parse(theUrl string) error {
	u, e := url.Parse("http://127.0.0.1:1234")
	if e != nil {
		return e
	}
	proxy := http.ProxyURL(u)
	http.DefaultClient.Transport = &http.Transport{Proxy: proxy}

	rsp, e := http.Get(theUrl)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	log.Printf("status %s %d", rsp.Status, len(b))
	return nil
}
