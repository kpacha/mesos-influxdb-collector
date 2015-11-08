all: deps build test

deps:
	go get github.com/hashicorp/hcl
	go get github.com/influxdb/influxdb/client

gen:
	go fmt ./...

build:
	go build

install: all do_install

do_install:
	go install

test:
	go test -cover ./...
