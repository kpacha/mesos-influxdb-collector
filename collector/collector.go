package collector

import (
	"github.com/kpacha/mesos-influxdb-collector/parser"
	"github.com/kpacha/mesos-influxdb-collector/store"
	"log"
	"net/http"
)

type Collector interface {
	Collect() ([]store.Point, error)
}

type MultiCollector struct {
	Collectors []Collector
}

func NewMultiCollector(collectors []Collector) Collector {
	return MultiCollector{collectors}
}

func (mc MultiCollector) Collect() ([]store.Point, error) {
	var data []store.Point
	for _, c := range mc.Collectors {
		ps, err := c.Collect()
		if err != nil {
			log.Println("Error collecting from", c)
			return data, err
		}
		data = append(data, ps...)
	}
	return data, nil
}

type UrlCollector struct {
	Url    string
	Parser parser.Parser
}

func (uc UrlCollector) Collect() ([]store.Point, error) {
	r, err := http.Get(uc.Url)
	if err != nil {
		log.Println("Error connecting to", uc.Url)
		return []store.Point{}, err
	}

	return uc.Parser.Parse(r.Body)
}
