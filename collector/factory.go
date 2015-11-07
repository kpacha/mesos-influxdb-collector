package collector

import (
	"fmt"
	"github.com/kpacha/mesos-influxdb-collector/config"
	"github.com/kpacha/mesos-influxdb-collector/parser/marathon"
	"github.com/kpacha/mesos-influxdb-collector/parser/mesos"
	"log"
	"net/url"
)

func NewCollectorFromConfig(configuration *config.Config) Collector {
	var collectors []Collector

	for _, master := range configuration.Master {
		collectors = append(collectors, NewMesosMasterCollector(master.Host, master.Port, master.Leader))
	}
	for _, slave := range configuration.Slave {
		collectors = append(collectors, NewMesosSlaveCollector(slave.Host, slave.Port))
	}
	for _, marathonInstance := range configuration.Marathon {
		collectors = append(collectors, NewMarathonCollector(marathonInstance.Host, marathonInstance.Port))
	}

	return NewMultiCollector(collectors)
}

func NewMultiCollector(collectors []Collector) Collector {
	return MultiCollector{collectors}
}

func NewMesosLeaderCollector(host string, port int) Collector {
	return NewMesosMasterCollector(host, port, true)
}

func NewMesosMasterCollector(host string, port int, leader bool) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/metrics/snapshot", host, port))
	if err != nil {
		log.Fatal("Error building the mesos master collector:", err)
	}
	return UrlCollector{Url: u.String(), Parser: mesos.MasterParser{Node: host, Leader: leader}}
}

func NewMesosSlaveCollector(host string, port int) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/metrics/snapshot", host, port))
	if err != nil {
		log.Fatal("Error building the mesos slave collector:", err)
	}
	return UrlCollector{Url: u.String(), Parser: mesos.SlaveParser{Node: host}}
}

func NewMarathonCollector(host string, port int) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/metrics", host, port))
	if err != nil {
		log.Fatal("Error building the marathon collector:", err)
	}
	return UrlCollector{Url: u.String(), Parser: marathon.MarathonParser{Node: host}}
}
