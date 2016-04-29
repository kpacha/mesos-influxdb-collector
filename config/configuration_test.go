package config

import (
	"fmt"
	"os"
)

func Example_defaultConfig() {
	os.Setenv("MIC_INFLUXDB_USER", "supu")
	os.Setenv("MIC_INFLUXDB_PASS", "secret")
	os.Setenv("MIC_HAPROXY_PASSWORD", "super_secret")
	os.Setenv("MIC_LAPSE", "1")
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
	// &{influxdb.marathon.mesos 8086 mesos 30 supu secret}
	// 1
	// 300
	// &{admin super_secret haproxy?stats 9090}
}
