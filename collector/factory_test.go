package collector

import (
	"fmt"

	"github.com/kpacha/mesos-influxdb-collector/config"
)

func Example_collectorFromConfig() {
	c := &config.Config{
		Master: []config.Master{
			config.Master{Host: "localhost", Port: 5050, Leader: true},
			config.Master{Host: "localhost", Port: 15050, Leader: false},
		},
		Slave: []config.Server{
			config.Server{Host: "localhost", Port: 5051},
			config.Server{Host: "localhost", Port: 5052},
		},
	}
	collector := NewCollectorFromConfig(c)
	fmt.Println(collector)

	// Output:
	// {[{http://localhost:5050/metrics/snapshot {localhost true} <nil> <nil>} {http://localhost:15050/metrics/snapshot {localhost false} <nil> <nil>} {[{http://localhost:5051/metrics/snapshot {localhost} <nil> <nil>} {http://localhost:5051/monitor/statistics {localhost} <nil> <nil>}]} {[{http://localhost:5052/metrics/snapshot {localhost} <nil> <nil>} {http://localhost:5052/monitor/statistics {localhost} <nil> <nil>}]}]}
}
