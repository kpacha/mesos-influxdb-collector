package config

import (
	"fmt"
	"os"
)

func Example_defaultConfig() {
	cp := NewConfigParser("json", "../.", "conf", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	fmt.Println(c.InfluxDB)
	fmt.Println(c.Lapse)
	fmt.Println(c.DieAfter)
	fmt.Println(c.HAProxy)
	// Output:
	// &{mesos true slave.mesos 8123}
	// []
	// []
	// &{[] false  8080 0}
	// &{influxdb.marathon.mesos 8086 mesos 30 root root}
	// 5
	// 300
	// &{admin admin haproxy?stats 9090}
}

func Example_envarHaveHigherPrio() {
	os.Setenv("MIC_INFLUXDB_USER", "supu")
	os.Setenv("MIC_INFLUXDB_PASS", "secret")
	os.Setenv("MIC_HAPROXY_PASSWORD", "super_secret")
	os.Setenv("MIC_LAPSE", "1")
	cp := NewConfigParser("json", "../fixtures/.", "conf1", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	fmt.Println(c.InfluxDB)
	fmt.Println(c.Lapse)
	fmt.Println(c.DieAfter)
	fmt.Println(c.HAProxy)
	os.Clearenv()
	// Output:
	// &{mesos true slave.mesos 8123}
	// []
	// []
	// &{[] false  8080 0}
	// &{influxdb.marathon.mesos 8086 mesos 30 supu secret}
	// 1
	// 300
	// &{admin super_secret haproxy?stats 9090}
}

func Example_parseMasters() {
	cp := NewConfigParser("json", "../fixtures/.", "conf2", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
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
	// [{localhost 5050 true} {localhost 5051 false}]
	// []
	// <nil>
	// &{localhost 8086 mesos 30 root root}
	// 10
	// 2
	// <nil>
}

func Example_parseSlaves() {
	cp := NewConfigParser("json", "../fixtures/.", "conf3", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
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
	// [{localhost 5051} {localhost 5052}]
	// <nil>
	// &{localhost 8086 mesos 30 root root}
	// 100
	// 3600
	// <nil>
}

func Example_parseMarathon() {
	cp := NewConfigParser("json", "../fixtures/.", "conf4", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
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
	// &{[{marathon1 8080} {marathon2 8080}] false  0 0}
	// &{localhost 8086 mesos 30 root root}
	// 30
	// 1
	// <nil>
}

func Example_parseMarathonWithEvents() {
	cp := NewConfigParser("json", "../fixtures/.", "conf5", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
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
	// &{[{marathon1 8080} {marathon2 8080}] true localhost 8088 10000}
	// &{localhost 8086 mesos 30 root root}
	// 10
	// 1
	// <nil>
}

func Example_parseInfluxDB() {
	cp := NewConfigParser("json", "../fixtures/.", "conf6", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
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
	// &{influx 18086 custom 30 root root}
	// 10
	// 3600
	// <nil>
}

func Example_parseHAProxy() {
	cp := NewConfigParser("json", "../fixtures/.", "conf7", false)
	c, err := cp.Parse()
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
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
	// &{localhost 8086 mesos 30 root root}
	// 10
	// 3600
	// &{admin super_secret haproxy_stats 19090}
}
