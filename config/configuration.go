package config

import (
	"log"

	"github.com/hashicorp/hcl"
	"io/ioutil"
)

type Config struct {
	MesosDNS *MesosDNS
	Master   []Master
	Slave    []Server
	Marathon []Server
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

type Master struct {
	Host   string
	Port   int
	Leader bool
}

type InfluxDB struct {
	Host string
	Port int
	DB   string
}

type ConfigParser struct {
	Path string
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

	if result.InfluxDB == nil {
		result.InfluxDB = &InfluxDB{"localhost", 8086, "mesos"}
	}

	log.Printf("%+v\n", result)

	return result, nil
}
