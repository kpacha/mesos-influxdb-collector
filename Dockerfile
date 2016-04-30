FROM golang:1.5.3
MAINTAINER kpacha

ENV FORMAT=json
ENV CONFIG_PATH=/go/src/github.com/kpacha/mesos-influxdb-collector
ENV CONFIG_FILE=conf

RUN mkdir -p /go/src/github.com/kpacha/mesos-influxdb-collector
COPY . /go/src/github.com/kpacha/mesos-influxdb-collector

WORKDIR /go/src/github.com/kpacha/mesos-influxdb-collector
RUN make install

ENTRYPOINT ["/go/bin/mesos-influxdb-collector"]
