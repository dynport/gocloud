default: build

build:
	cd aws/pricing && make assets
	go get ./...

test:
	cd jiffybox && go test -v
	cd aws/ec2 && go test -v
	cd aws/elb && go test -v
	cd aws/pricing && go test -v
	cd aws/route53 && go test -v
	cd digitalocean && go test -v
