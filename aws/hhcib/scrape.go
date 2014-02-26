package main

import (
	"log"
	"net/url"
	"strings"
	"time"
)

var Null = struct{}{}

type Scrape struct {
	found map[string]struct{}
}

var prefix = "http://aws.amazon.com/documentation"
var fetched = map[string]time.Time{}

func (scrape *Scrape) Run() error {
	theUrl := "http://aws.amazon.com/documentation/"
	theUrl = "http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Operations.html"
	return fetchUrl(theUrl)
}

func fetchUrl(theUrl string) error {
	if _, ok := fetched[theUrl]; ok {
		return nil
	}
	log.Printf("fetching %s", theUrl)
	u, e := url.Parse(theUrl)
	if e != nil {
		return e
	}
	doc, e := loadDoc(theUrl)
	if e != nil {
		return e
	}
	fetched[theUrl] = time.Now()
	links, e := doc.Search("//a")
	if e != nil {
		return e
	}
	for _, link := range links {
		theUrl := normalizeUrl(u, link.Attr("href"))
		if doFetch(theUrl) {
			e := fetchUrl(theUrl)
			if e != nil {
				log.Println("ERROR: " + e.Error())
			}
		}
	}
	if strings.Contains(doc.InnerHtml(), `myDefaultPage = "Welcome.html";`) && strings.HasSuffix(theUrl, "/") {
		e := fetchUrl(theUrl + "Welcome.html")
		if e != nil {
			log.Println("ERROR: " + e.Error())
		}
	}
	fetched[theUrl] = time.Now()
	return nil
}

func normalizeUrl(u *url.URL, href string) string {
	if strings.HasPrefix(href, "//") {
		href = u.Scheme + ":" + href
	} else if strings.HasPrefix(href, "/") {
		href = u.Scheme + "://" + u.Host + href
	}
	return strings.Split(href, "#")[0]
}

func doFetch(theUrl string) bool {
	if strings.HasSuffix(theUrl, ".pdf") {
		return false
	}
	return strings.HasPrefix(theUrl, prefix) || strings.HasPrefix(theUrl, "http://docs.aws.amazon.com/")
}
