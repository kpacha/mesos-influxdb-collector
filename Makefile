all: deps build test

deps:
	go get -v -u github.com/hashicorp/hcl
	go get -v -u github.com/influxdata/influxdb/client
	go get -v -u github.com/stretchr/testify

gen:
	go fmt ./...

build:
	sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o mesos-influxdb-collector'

install: all do_install

do_install:
	go mod init
	go mod tidy
	go install

test:
	go test -cover ./...
	go vet ./...

docker:
	docker run --rm -it -e "GOPATH=/go" -v "${PWD}:/go/src/github.com/kpacha/mesos-influxdb-collector" -w /go/src/github.com/kpacha/mesos-influxdb-collector golang:1.5.3 make
	docker build -f Dockerfile-min -t kpacha/mesos-influxdb-collector:latest-min .
