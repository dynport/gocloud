package main

import (
	"log"
	"os"

	"github.com/dynport/gocloud/aws/cloudformation"
	"github.com/dynport/gocloud/aws/ec2"
	"github.com/dynport/martini"
)

var (
	logger    = log.New(os.Stderr, "", 0)
	ec2Client = ec2.NewFromEnv()
	cfClient  = cloudformation.NewFromEnv()
	app       = martini.Classic()
)

func main() {
	app.Run()
}
