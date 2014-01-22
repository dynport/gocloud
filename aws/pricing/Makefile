test:
	go test -v

build: assets
	go get .

assets: install_goassets
	rm -f ./compiled_assets.go
	goassets -pkg=pricing -path=./compiled_assets.go ./assets

install_goassets:
	which goassets || go get github.com/dynport/dgtk/goassets
