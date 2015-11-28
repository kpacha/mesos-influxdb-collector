package config

import (
	"io/ioutil"
	"log"

	"github.com/hashicorp/hcl"
)

type Config struct {
	MesosDNS *MesosDNS
	Master   []Master
	Slave    []Server
	Marathon *Marathon
	InfluxDB *InfluxDB
	Lapse    int
	DieAfter int
}

type MesosDNS struct {
	Domain   string
	Marathon bool
	Host     string
	Port     int
}

type Server struct {
	Host string
	Port int
}

type Marathon struct {
	Server     []Server
	Host       string
	Port       int
	BufferSize int
}

type Master struct {
	Host   string
	Port   int
	Leader bool
}

type InfluxDB struct {
	Host       string
	Port       int
	DB         string
	CheckLapse int
}

type ConfigParser struct {
	Path     string
	AllowDNS bool
}

func (cp ConfigParser) Parse() (*Config, error) {
	return cp.ParseConfigFile(cp.Path)
}

func (cp ConfigParser) ParseConfigFile(file string) (*Config, error) {
	hclText, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return cp.ParseConfig(string(hclText))
}

func (cp ConfigParser) ParseConfig(hclText string) (*Config, error) {
	result := &Config{}

	if err := hcl.Decode(&result, hclText); err != nil {
		return nil, err
	}

	if result.Lapse == 0 {
		result.Lapse = 30
	}
	if result.DieAfter == 0 {
		result.DieAfter = 3600
	}

	if result.InfluxDB == nil {
		result.InfluxDB = &InfluxDB{"localhost", 8086, "mesos", 30}
	}
	if result.InfluxDB.CheckLapse == 0 {
		result.InfluxDB.CheckLapse = 30
	}

	if cp.AllowDNS && result.MesosDNS != nil {
		if _, err := NewDNSResolver(result); err != nil {
			return nil, err
		}
	}

	log.Printf("%+v\n", result)

	return result, nil
}
