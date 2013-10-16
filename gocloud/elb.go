package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/aws/elb"
	"strings"
)

const (
	USAGE_LB           = "LB"
	USAGE_ELB_REGISTER = "LB " + USAGE_TERMINATE_INSTANCES
)

// elb
func init() {
	router.Register("aws/elb/lbs/list", &gocli.Action{
		Handler: elbListLoadBalancers, Description: "Describe load balancers",
	})

	router.Register("aws/elb/lbs/deregister", &gocli.Action{
		Handler: elbDeregisterInstances, Description: "Deregister instances with load balancer",
		Usage: USAGE_ELB_REGISTER,
	})

	router.Register("aws/elb/lbs/register", &gocli.Action{
		Handler: elbRegisterInstances, Description: "Register instances with load balancer",
		Usage: USAGE_ELB_REGISTER,
	})

	router.Register("aws/elb/lbs/describe", &gocli.Action{
		Handler: elbDescribeLoadBalancer, Description: "Describe load balancers", Usage: USAGE_LB,
	})
}

func elbDescribeLoadBalancer(args *gocli.Args) error {
	if len(args.Args) != 1 {
		return fmt.Errorf(USAGE_LB)
	}
	elbClient := elb.NewFromEnv()
	states, e := elbClient.DescribeInstanceHealth(args.Args[0])
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	for _, state := range states {
		stateString := ""
		if state.State == "InService" {
			stateString = gocli.Green(state.State)
		} else {
			stateString = gocli.Red(state.State)
		}
		table.Add(state.InstanceId, stateString)
	}
	fmt.Println(table)
	return nil
}

func elbRegisterInstances(args *gocli.Args) error {
	if len(args.Args) < 2 {
		return fmt.Errorf(USAGE_ELB_REGISTER)
	}
	elbClient := elb.NewFromEnv()
	return elbClient.RegisterInstancesWithLoadBalancer(args.Args[0], args.Args[1:len(args.Args)])
}

func elbDeregisterInstances(args *gocli.Args) error {
	if len(args.Args) < 2 {
		return fmt.Errorf(USAGE_ELB_REGISTER)
	}
	elbClient := elb.NewFromEnv()
	return elbClient.DeregisterInstancesWithLoadBalancer(args.Args[0], args.Args[1:len(args.Args)])
}

func elbListLoadBalancers(args *gocli.Args) error {
	elbClient := elb.NewFromEnv()
	logger.Info("describing load balancers")
	lbs, e := elbClient.DescribeLoadBalancers()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Name", "DnsName", "Instances")
	for _, lb := range lbs {
		table.Add(
			lb.LoadBalancerName,
			lb.DNSName,
			strings.Join(lb.Instances, ", "),
		)
	}
	fmt.Print(table)
	return nil
}
