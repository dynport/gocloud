default:
	go install github.com/dynport/gocloud/gocloud

test:
	cd jiffybox && go test -v
	cd aws/ec2 && go test -v
	cd aws/elb && go test -v
	cd aws/route53 && go test -v
