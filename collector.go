package main

import (
	"fmt"
	"github.com/kpacha/mesos-influxdb-collector/collector"
	"github.com/kpacha/mesos-influxdb-collector/parser/mesos"
	"log"
	"net/url"
)

func NewMesosMasterCollector(host string, port int) collector.Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/metrics/snapshot", host, port))
	if err != nil {
		log.Fatal("Error building the mesos master collector:", err)
	}
	return collector.UrlCollector{Url: u.String(), Parser: mesos.MasterParser{Node: host}}
}

func NewMesosSlaveCollector(host string, port int) collector.Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/metrics/snapshot", host, port))
	if err != nil {
		log.Fatal("Error building the mesos slave collector:", err)
	}
	return collector.UrlCollector{Url: u.String(), Parser: mesos.SlaveParser{Node: host}}
}
