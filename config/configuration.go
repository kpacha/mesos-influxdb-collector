package config

import (
	"io/ioutil"
	"log"

	"github.com/hashicorp/hcl"
)

type Config struct {
	MesosDNS *MesosDNS
	Master   []Master `hcl:"master,expand"`
	Slave    []Server `hcl:"slave,expand"`
	Marathon *Marathon
	InfluxDB *InfluxDB
	Lapse    int
	DieAfter int
	HAProxy  *HAProxy
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
	Server     []Server `hcl:"server,expand"`
	Events     bool
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

type HAProxy struct {
	User     string
	Password string
	EndPoint string
	Port     int
}

type ConfigParser struct {
	Path     string
	AllowDNS bool
	Default  *Config
}

var (
	EmptyString   = ""
	EmptyInt      = -1
	DefaultConfig = &Config{
		InfluxDB: &InfluxDB{"localhost", 8086, "mesos", 30},
		Lapse:    30,
		DieAfter: 3600,
	}
)

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
	var tmp Config
	if cp.Default != nil {
		tmp = *cp.Default
	} else {
		tmp = *DefaultConfig
	}
	err := cp.UpdateConfig(hclText, &tmp)
	return &tmp, err
}

func (cp ConfigParser) UpdateConfig(hclText string, result *Config) error {
	var tmp Config
	tmp = *result
	if err := hcl.Decode(&tmp, hclText); err != nil {
		return err
	}

	if tmp.InfluxDB.CheckLapse == 0 {
		tmp.InfluxDB.CheckLapse = DefaultConfig.InfluxDB.CheckLapse
	}

	if cp.AllowDNS && tmp.MesosDNS != nil {
		if _, err := NewDNSResolver(&tmp); err != nil {
			return err
		}
	}

	log.Printf("%+v\n", tmp)
	*result = tmp

	return nil
}

func (cp ConfigParser) ParseAndMerge(ihost *string, iport *int, idb *string) (*Config, error) {
	conf, err := cp.Parse()
	if err != nil {
		return nil, err
	}

	if *ihost != EmptyString {
		conf.InfluxDB.Host = *ihost
	}
	if *iport != EmptyInt {
		conf.InfluxDB.Port = *iport
	}
	if *idb != EmptyString {
		conf.InfluxDB.DB = *idb
	}

	return conf, nil
}
