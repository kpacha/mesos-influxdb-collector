package marathon

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/kpacha/mesos-influxdb-collector/store"
)

type MarathonEvent struct {
	Type   string `json:"eventType"`
	Status string `json:"taskStatus,omitempty"`
	ID     string `json:"appId,omitempty"`
	Time   time.Time
	Node   string
}

type MarathonEventsParser struct{}

func (mp MarathonEventsParser) Parse(r io.ReadCloser, from string) ([]store.Point, error) {
	defer r.Close()
	event, err := mp.parse(r, from)
	if err != nil {
		return []store.Point{}, err
	}
	return []store.Point{mp.getMarathoneventPoint(event)}, nil
}

func (mp MarathonEventsParser) parse(r io.ReadCloser, from string) (*MarathonEvent, error) {
	var event *MarathonEvent
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("Error reading from", r)
		return nil, err
	}
	if err = json.Unmarshal(body, &event); err != nil {
		log.Println("Error parsing to MarathonEvents")
		return nil, err
	}
	event.Node = from
	event.Time = time.Now()
	return event, nil
}

func (mp MarathonEventsParser) getMarathoneventPoint(event *MarathonEvent) store.Point {
	return store.Point{
		Measurement: "marathon-event",
		Tags: map[string]string{
			"node":   event.Node,
			"type":   event.Type,
			"status": event.Status,
			"appId":  event.ID,
		},
		Fields: map[string]interface{}{
			"value": 1,
		},
		Time: event.Time,
	}
}
