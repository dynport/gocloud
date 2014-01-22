package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var instanceTypeUrl = "http://aws.amazon.com/ec2/instance-types/"

func main() {
	if e := run(); e != nil {
		log.Fatal("ERROR: " + e.Error())
	}
}

func run() error {
	rsp, e := http.Get(instanceTypeUrl)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)

	types, e := parseInstanceTypes(b)
	if e != nil {
		return e
	}
	return json.NewEncoder(os.Stdout).Encode(types)
}
