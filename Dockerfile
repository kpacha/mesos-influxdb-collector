FROM golang:latest
MAINTAINER kpacha

ENV INFLUXDB_HOST=influxdb
ENV INFLUXDB_PORT=8086
ENV INFLUXDB_DB=mesos
ENV INFLUXDB_USER=root
ENV INFLUXDB_PWD=root

RUN mkdir -p /go/src/github.com/kpacha/mesos-influxdb-collector
COPY . /go/src/github.com/kpacha/mesos-influxdb-collector

WORKDIR /go/src/github.com/kpacha/mesos-influxdb-collector
RUN export GO111MODULE=on && make install

ENTRYPOINT ["/go/bin/mesos-influxdb-collector"]

CMD ["-c", "conf.hcl"]
