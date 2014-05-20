package main

import (
	"log"
	"net/http"
	"strings"
)

func fallback(r *http.Request, log *log.Logger) (int, string) {
	name := strings.TrimPrefix(r.URL.Path, "/")
	log.Printf("serving %q", name)
	for _, n := range assetNames() {
		if n == name {
			b, e := readAsset(n)
			if e != nil {
				return 500, e.Error()
			}
			log.Printf("serving goasset %q", n)
			return 200, string(b)
		}
	}
	b, e := readAsset("layout.html")
	if e != nil {
		return 500, e.Error()
	}
	return 200, string(b)
}
