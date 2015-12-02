package marathon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kpacha/mesos-influxdb-collector/config"
	"github.com/kpacha/mesos-influxdb-collector/parser"
	"github.com/kpacha/mesos-influxdb-collector/store"
)

type MarathonRedisteredCallbacks struct {
	URL []string `json:"callbackUrls"`
}

type MarathonEventsSubscriber struct {
	configuration *config.Marathon
	parser        parser.ParserFrom
	buffer        chan []store.Point
	callback      string
}

func NewMarathonEventsSubscriber(configuration *config.Marathon, p parser.ParserFrom) chan []store.Point {
	callback := fmt.Sprintf("http://%s:%d/marathon", strings.TrimSuffix(configuration.Host, "/"), configuration.Port)
	mes := MarathonEventsSubscriber{
		configuration: configuration,
		parser:        p,
		buffer:        make(chan []store.Point, configuration.BufferSize),
		callback:      callback,
	}
	go mes.run()
	return mes.buffer
}

func (mes MarathonEventsSubscriber) run() {
	mes.register()
	defer mes.unregister()

	http.HandleFunc("/marathon", mes.marathonListener)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", mes.configuration.Port), nil))
	log.Println("DONE!")
}

func (mes MarathonEventsSubscriber) register() error {
	var subsEndpoint string
	for _, server := range mes.configuration.Server {
		subsEndpoint = fmt.Sprintf("http://%s:%d/v2/eventSubscriptions", server.Host, server.Port)
		if !mes.isAlreadyRegistered(subsEndpoint) {
			if err := mes.registerEndpoint(subsEndpoint); err != nil {
				return err
			}
		}
	}
	return nil
}

func (mes MarathonEventsSubscriber) unregister() error {
	var subsEndpoint string
	for _, server := range mes.configuration.Server {
		subsEndpoint = fmt.Sprintf("http://%s:%d/v2/eventSubscriptions", server.Host, server.Port)
		if err := mes.unregisterEndpoint(subsEndpoint); err != nil {
			return err
		}
	}
	return nil
}

func (mes MarathonEventsSubscriber) marathonListener(w http.ResponseWriter, r *http.Request) {
	event, err := mes.parser.Parse(r.Body, r.RemoteAddr)
	if err == nil {
		mes.buffer <- event
	}
}

func (mes MarathonEventsSubscriber) registerEndpoint(subsEndpoint string) error {
	err := mes.sendHttpRequest("POST", subsEndpoint)
	if err != nil {
		return fmt.Errorf("Error registering: %s", err)
	}
	return nil

}

func (mes MarathonEventsSubscriber) unregisterEndpoint(subsEndpoint string) error {
	err := mes.sendHttpRequest("DELETE", subsEndpoint)
	if err != nil {
		return fmt.Errorf("Error unregistering: %s", err)
	}
	return nil
}

func (mes MarathonEventsSubscriber) sendHttpRequest(method, subsEndpoint string) error {
	endpoint := fmt.Sprintf("%s?callbackUrl=%s", subsEndpoint, mes.callback)

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error sending %s to %s. Status code [%d]", method, subsEndpoint, resp.StatusCode)
	}

	return nil
}

func (mes MarathonEventsSubscriber) isAlreadyRegistered(subsEndpoint string) bool {
	resp, err := http.Get(subsEndpoint)
	if err != nil {
		log.Fatalf("Error retrieving event subscribers list: %s", err)
	}
	defer resp.Body.Close()

	callbacks := MarathonRedisteredCallbacks{}
	if err := json.NewDecoder(resp.Body).Decode(&callbacks); nil != err {
		log.Fatalf("Error decoding event subscribers list: %s", err)
	}

	for _, url := range callbacks.URL {
		if url == mes.callback {
			return true
		}
	}

	return false
}
