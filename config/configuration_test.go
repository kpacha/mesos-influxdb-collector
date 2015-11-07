package config

import (
	"fmt"
	"testing"
)

func TestParseMesosDNS(t *testing.T) {
	txtConfig := `mesosDNS {
		domain = "mesos"
		marathon = true
		host = "localhost"
		port = 53
	}`
	cp := ConfigParser{}
	c, err := cp.ParseConfig(txtConfig)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	// Output: &{mesos true localhost 53}
	// []
	// []
	// []
}

func TestParseMaster(t *testing.T) {
	txtConfig := `Master {
		host = "localhost"
		port = 5051
		leader = true
	}
	Master {
		host = "localhost"
		port = 5052
	}`
	cp := ConfigParser{}
	c, err := cp.ParseConfig(txtConfig)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	// Output: nil
	// [{localhost 5051 true} {localhost 5052 false}]
	// []
	// []
}

func TestParseSlave(t *testing.T) {
	txtConfig := `Slave {
		host = "localhost"
		port = 5051
	}
	Slave {
		host = "localhost"
		port = 5052
	}`
	cp := ConfigParser{}
	c, err := cp.ParseConfig(txtConfig)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	// Output: nil
	// []
	// [{localhost 5051} {localhost 5052}]
	// []
}

func TestParseMarathon(t *testing.T) {
	txtConfig := `Marathon {
		host = "localhost"
		port = 5051
	}
	Marathon {
		host = "localhost"
		port = 5052
	}`
	cp := ConfigParser{}
	c, err := cp.ParseConfig(txtConfig)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c.MesosDNS)
	fmt.Println(c.Master)
	fmt.Println(c.Slave)
	fmt.Println(c.Marathon)
	// Output: nil
	// []
	// []
	// [{localhost 5051} {localhost 5052}]
}
