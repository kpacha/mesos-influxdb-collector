package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Collector interface {
	Collect(stats *Stats) error
}

type MesosCollector struct {
	url  string
	node string
}

func NewMesosCollector(host string, port int) Collector {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/stats.json", host, port))
	if err != nil {
		log.Fatal("Error collecting the stats from mesos: ", err)
	}
	return MesosCollector{url: u.String(), node: host}
}

func (m MesosCollector) Collect(stats *Stats) error {
	r, err := http.Get(m.url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(stats); err != nil {
		return err
	}
	stats.Node = m.node
	stats.Time = time.Now()
	return nil
}
