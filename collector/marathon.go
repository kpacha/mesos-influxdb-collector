package collector

import (
	"time"

	"github.com/kpacha/mesos-influxdb-collector/collector/marathon"
	"github.com/kpacha/mesos-influxdb-collector/config"
	"github.com/kpacha/mesos-influxdb-collector/parser"
	"github.com/kpacha/mesos-influxdb-collector/store"
)

type MarathonRedisteredCallbacks struct {
	URL []string `json:"callbackUrls"`
}

type MarathonEventsCollector struct {
	buffer chan []store.Point
}

func NewMarathonEventsCollector(configuration *config.Marathon, p parser.ParserFrom) Collector {
	return MarathonEventsCollector{marathon.NewMarathonEventsSubscriber(configuration, p)}
}

func (mec MarathonEventsCollector) Collect() ([]store.Point, error) {
	points := []store.Point{}
	for {
		select {
		case ps := <-mec.buffer:
			points = append(points, ps...)
		case <-time.After(time.Second):
			return points, nil
		}
	}
}
