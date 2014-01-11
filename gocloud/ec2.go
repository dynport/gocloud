package main

import (
	"github.com/dynport/gocloud/cli/aws/ec2"
)

func init() {
	ec2.Register(router)
}
