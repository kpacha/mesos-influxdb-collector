package config

import (
	"fmt"
)

func ExampleParseMaster() {
	txtConfig := `Master {
		host = "localhost"
		port = 5051
		leader = true
	}
	Master {
		host = "localhost"
		port = 5052
	}
	dieafter = 1`
	cp := ConfigParser{}
	c, _ := cp.ParseConfig(txtConfig)
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	fmt.Println(c.InfluxDB)
	fmt.Println(c.Lapse)
	fmt.Println(c.DieAfter)
	// Output:
	// <nil>
	// [{localhost 5051 true} {localhost 5052 false}]
	// []
	// []
	// &{localhost 8086 mesos 30}
	// 30
	// 1
}

func ExampleParseSlave() {
	txtConfig := `Slave {
		host = "localhost"
		port = 5051
	}
	Slave {
		host = "localhost"
		port = 5052
	}
	lapse=100
	dieAfter = 1`
	cp := ConfigParser{}
	c, _ := cp.ParseConfig(txtConfig)
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	fmt.Println(c.InfluxDB)
	fmt.Println(c.Lapse)
	fmt.Println(c.DieAfter)
	// Output:
	// <nil>
	// []
	// [{localhost 5051} {localhost 5052}]
	// []
	// &{localhost 8086 mesos 30}
	// 100
	// 1
}

func ExampleParseMarathon() {
	txtConfig := `Marathon {
		host = "localhost"
		port = 5051
	}
	Marathon {
		host = "localhost"
		port = 5052
	}
	DieAfter = 1`
	cp := ConfigParser{}
	c, _ := cp.ParseConfig(txtConfig)
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	fmt.Println(c.InfluxDB)
	fmt.Println(c.Lapse)
	fmt.Println(c.DieAfter)
	// Output:
	// <nil>
	// []
	// []
	// [{localhost 5051} {localhost 5052}]
	// &{localhost 8086 mesos 30}
	// 30
	// 1
}

func ExampleParseInfluxDB() {
	txtConfig := `influxdb {
		host = "influx"
		port = 18086
		db = "custom"
	}`
	cp := ConfigParser{}
	c, _ := cp.ParseConfig(txtConfig)
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	fmt.Println(c.InfluxDB)
	fmt.Println(c.Lapse)
	fmt.Println(c.DieAfter)
	// Output:
	// <nil>
	// []
	// []
	// []
	// &{influx 18086 custom 30}
	// 30
	// 3600
}
