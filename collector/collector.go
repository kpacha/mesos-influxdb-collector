package collector

import (
	"log"
	"net/http"

	"github.com/kpacha/mesos-influxdb-collector/parser"
	"github.com/kpacha/mesos-influxdb-collector/store"
)

type Collector interface {
	Collect() ([]store.Point, error)
}

type MultiCollector struct {
	Collectors []Collector
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
	Url      string
	Parser   parser.Parser
	User     *string
	Password *string
}

func (uc UrlCollector) Collect() ([]store.Point, error) {
	req, err := http.NewRequest("GET", uc.Url, nil)
	if err != nil {
		log.Println("Error building request to", uc.Url)
		return []store.Point{}, err
	}
	if uc.User != nil && uc.Password != nil {
		req.SetBasicAuth(*uc.User, *uc.Password)
	}
	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		log.Println("Error connecting to", uc.Url)
		return []store.Point{}, err
	}

	return uc.Parser.Parse(r.Body)
}
