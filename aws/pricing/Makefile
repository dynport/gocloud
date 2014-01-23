test:
	go test -v

build: assets
	go get .

assets: install_goassets
	rm -f ./compiled_assets.go
	goassets -pkg=pricing -path=./compiled_assets.go ./assets

install_goassets:
	which goassets || go get github.com/dynport/dgtk/goassets

update_instance_types:
	which aws-dump-instance-types || go get github.com/dynport/gocloud/aws/pricing/aws-dump-instance-types
	aws-dump-instance-types > assets/instance_types.json
