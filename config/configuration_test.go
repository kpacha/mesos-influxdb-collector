package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Example_defaultConfig() {
	os.Setenv("MIC_INFLUXDB_USER", "supu")
	os.Setenv("MIC_INFLUXDB_PASS", "secret")
	os.Setenv("MIC_LAPSE", "1")
	Debug = true
	cp := NewConfigParser("json", "../.", "conf")
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
	fmt.Println(viper.Get("influxdb.user"))
	fmt.Println(viper.AllSettings())
	// Output:
	// &{mesos true slave.mesos 8123}
	// []
	// []
	// &{[] false  8080 0}
	// &{influxdb.marathon.mesos 8086 mesos 30 supu secret}
	// 1
	// 300

}
