package main

import "os"

type TestIntegration struct {
}

func (i *TestIntegration) Run() error {
	logger.Print("testing integration")
	client := &AwsClient{
		Endpoint:        "https://eu-west-1.ec2.amazonaws.com",
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
	}
	a := DescribeInstances{}
	client.loadResource(a)
	return nil
}

func init() {
	router.Register("integration", &TestIntegration{}, "Test Integration")
}
