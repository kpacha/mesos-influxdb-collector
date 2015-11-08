FROM golang:1.5.1
MAINTAINER kpacha

ENV INFLUXDB_USER=root
ENV INFLUXDB_PWD=root

RUN mkdir -p /go/src/github.com/kpacha/mesos-influxdb-collector
COPY . /go/src/github.com/kpacha/mesos-influxdb-collector

WORKDIR /go/src/github.com/kpacha/mesos-influxdb-collector
RUN make install

ENTRYPOINT ["/go/bin/mesos-influxdb-collector"]

CMD ["-c", "conf.hcl"]