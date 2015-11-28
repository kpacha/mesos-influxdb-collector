package collector

import (
	"fmt"

	"github.com/kpacha/mesos-influxdb-collector/config"
)

func ExampleCollectorFromConfig() {
	txtConfig := `master "leader" {
		host = "localhost"
		port = 5050
		leader = true
	}
	master "follower" {
		host = "localhost"
		port = 15050
	}
	slave "0" {
		host = "localhost"
		port = 5051
	}
	slave "1" {
		host = "localhost"
		port = 5052
	}`
	cp := config.ConfigParser{}
	c, err := cp.ParseConfig(txtConfig)
	if err != nil {
		fmt.Printf("Error parsing the config:", err.Error())
	}
	collector := NewCollectorFromConfig(c)
	fmt.Println(collector)

	// Output:
	// {[{http://localhost:5050/metrics/snapshot {localhost true}} {http://localhost:15050/metrics/snapshot {localhost false}} {http://localhost:5051/metrics/snapshot {localhost}} {http://localhost:5052/metrics/snapshot {localhost}}]}
}
