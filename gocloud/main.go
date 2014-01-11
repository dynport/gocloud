package main

import (
	"github.com/dynport/dgtk/cli"
	"github.com/dynport/gologger"
	"log"
	"os"
)

var (
	logger = gologger.NewFromEnv()
	router = cli.NewRouter()
)

func init() {
	log.SetFlags(0)
}

func main() {
	if e := router.RunWithArgs(); e != nil {
		logger.Error(e.Error())
		os.Exit(1)
	}
}
