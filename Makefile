default:
	go get github.com/dynport/gocloud/gocloud
	go get github.com/dynport/gocloud/aws/ec2
	go get github.com/dynport/gocloud/aws/s3

test:
	cd jiffybox && go test -v
	cd aws/ec2 && go test -v
	cd aws/elb && go test -v
	cd aws/route53 && go test -v
	cd digitalocean && go test -v
