package config

import (
	"fmt"
)

func Example_parseMaster() {
	txtConfig := `Master "leader" {
		host = "localhost"
		port = 5051
		leader = true
	}
	Master "follower" {
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
	// <nil>
	// &{localhost 8086 mesos 30}
	// 30
	// 1
}

func Example_parseSlave() {
	txtConfig := `Slave "0" {
		host = "localhost"
		port = 5051
	}
	slave "1" {
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
	// <nil>
	// &{localhost 8086 mesos 30}
	// 100
	// 1
}

func Example_parseMarathon() {
	txtConfig := `Marathon {
		Server "0" {
			host = "marathon1"
			port = 8080
		}
		Server "1" {
			host = "marathon2"
			port = 8080
		}
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
	// &{[{marathon1 8080} {marathon2 8080}] false  0 0}
	// &{localhost 8086 mesos 30}
	// 30
	// 1
}

func Example_parseMarathonWithEvents() {
	txtConfig := `Marathon {
		events = true
		host = "localhost"
		port = 8088
		bufferSize = 10000
		Server "0" {
			host = "marathon1"
			port = 8080
		}
		Server "1" {
			host = "marathon2"
			port = 8080
		}
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
	// &{[{marathon1 8080} {marathon2 8080}] true localhost 8088 10000}
	// &{localhost 8086 mesos 30}
	// 30
	// 1
}

func Example_parseInfluxDB() {
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
	fmt.Println(c.HAProxy)
	// Output:
	// <nil>
	// []
	// []
	// <nil>
	// &{influx 18086 custom 30}
	// 30
	// 3600
	// <nil>
}

func Example_parseHAProxy() {
	txtConfig := `haproxy {
		user = "admin"
		password = "admin"
		endPoint = "haproxy_stats"
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
	fmt.Println(c.HAProxy)
	// Output:
	// <nil>
	// []
	// []
	// <nil>
	// &{localhost 8086 mesos 30}
	// 30
	// 3600
	// &{admin admin haproxy_stats}
}
