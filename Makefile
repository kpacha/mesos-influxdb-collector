all: deps build test

deps:
	go get -v github.com/hashicorp/hcl
	go get -v github.com/influxdb/influxdb/client

gen:
	go fmt ./...

build:
	sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o mesos-influxdb-collector'

install: all do_install

do_install:
	go install

test:
	go test -cover ./...
	go vet ./...

docker:
	docker run --rm -it -e "GOPATH=/go" -v "${PWD}:/go/src/github.com/kpacha/mesos-influxdb-collector" -w /go/src/github.com/kpacha/mesos-influxdb-collector golang:1.5.3 make
	docker build -f Dockerfile-min -t kpacha/mesos-influxdb-collector:latest-min .