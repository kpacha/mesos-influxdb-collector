package collector

import (
	"fmt"
	"log"
	"net/url"

	"github.com/kpacha/mesos-influxdb-collector/config"
	"github.com/kpacha/mesos-influxdb-collector/parser/haproxy"
	"github.com/kpacha/mesos-influxdb-collector/parser/marathon"
	"github.com/kpacha/mesos-influxdb-collector/parser/mesos"
)

func NewCollectorFromConfig(configuration *config.Config) Collector {
	var collectors []Collector

	for _, master := range configuration.Master {
		collectors = append(collectors, NewMesosMasterCollector(master.Host, master.Port, master.Leader))
	}
	for _, slave := range configuration.Slave {
		collectors = append(collectors, NewMesosSlaveCollector(slave.Host, slave.Port))
	}
	if configuration.Marathon != nil {
		collectors = append(collectors, NewMarathonCollectors(configuration.Marathon)...)
	}
	if configuration.HAProxy != nil {
		for _, slave := range configuration.Slave {
			collectors = append(collectors, NewHAProxyCollector(slave.Host, configuration.HAProxy))
		}
	}

	log.Println("Total collectors created:", len(collectors))

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
	return NewMultiCollector([]Collector{
		NewMesosSlaveMainCollector(host, port),
		NewMesosSlaveStatsCollector(host, port),
	})
}

func NewMesosSlaveMainCollector(host string, port int) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/metrics/snapshot", host, port))
	if err != nil {
		log.Fatal("Error building the mesos slave collector:", err)
	}
	return UrlCollector{Url: u.String(), Parser: mesos.SlaveParser{Node: host}}
}

func NewMesosSlaveStatsCollector(host string, port int) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/monitor/statistics", host, port))
	if err != nil {
		log.Fatal("Error building the mesos slave stats collector:", err)
	}
	return UrlCollector{Url: u.String(), Parser: mesos.SlaveStatsParser{Node: host}}
}

func NewMarathonCollectors(configuration *config.Marathon) []Collector {
	if configuration.Server == nil {
		return []Collector{}
	}
	collectors := []Collector{}
	if configuration.Events {
		collectors = append(collectors, NewMarathonEventsCollector(configuration, marathon.MarathonEventsParser{}))
	}
	for _, marathonInstance := range configuration.Server {
		collectors = append(collectors, NewMarathonStatsCollector(marathonInstance.Host, marathonInstance.Port))
	}
	return collectors
}

func NewMarathonStatsCollector(host string, port int) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/metrics", host, port))
	if err != nil {
		log.Fatal("Error building the marathon collector:", err)
	}
	return UrlCollector{Url: u.String(), Parser: marathon.MarathonStatsParser{Node: host}}
}

func NewHAProxyCollector(host string, configuration *config.HAProxy) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s/%s;csv", host, configuration.EndPoint))
	if err != nil {
		log.Fatal("Error building the HAProxy collector:", err)
	}
	return UrlCollector{
		Url:      u.String(),
		Parser:   haproxy.NewHAProxy(host),
		User:     &configuration.User,
		Password: &configuration.Password,
	}
}
