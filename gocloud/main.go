package main

import (
	"github.com/dynport/gocli"
	"github.com/dynport/gologger"
	"os"
)

var logger = gologger.NewFromEnv()
var router = gocli.NewRouter(nil)

func main() {
	if e := router.Handle(os.Args); e != nil {
		logger.Error(e.Error())
		os.Exit(1)
	}
}
