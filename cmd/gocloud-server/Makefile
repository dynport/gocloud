.PHONY: assets

default: build

build: assets 
	go get .
	
assets:
	cd assets && coffee -c .
	rm -f assets.go
	goassets assets
