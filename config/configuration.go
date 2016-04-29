package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
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
	User       string
	Pass       string
}

type HAProxy struct {
	User     string
	Password string
	EndPoint string
}

type ConfigParser struct {
	cfg *viper.Viper
}

var (
	DefaultInfluxdb = InfluxDB{
		Host:       "localhost",
		DB:         "mesos",
		Port:       8086,
		CheckLapse: 30,
		User:       "root",
		Pass:       "root",
	}
	DefaultLapse    = 10
	DefaultDieAfter = 3600
	Debug           = false
)

func NewConfigParser(format, path, configName string) *ConfigParser {
	return &ConfigParser{newViper(format, path, configName)}
}

func (cp *ConfigParser) Parse() (*Config, error) {
	err := cp.cfg.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	var c Config
	err = cp.cfg.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return &c, nil
}

func newViper(format, path, configName string) *viper.Viper {
	cfg := viper.New()

	if Debug {
		cfg.Debug()
	}

	cfg.SetEnvPrefix("mic")
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	cfg.SetConfigType(format)
	cfg.AddConfigPath(path)
	cfg.SetConfigName(configName)

	cfg.SetDefault("influxdb", DefaultInfluxdb)
	cfg.SetDefault("lapse", DefaultLapse)
	cfg.SetDefault("dieAfter", DefaultDieAfter)

	return cfg
}

// func (cp ConfigParser) ParseConfigFile(file string) (*Config, error) {
// 	hclText, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cp.ParseConfig(string(hclText))
// }

// func (cp ConfigParser) ParseConfig(hclText string) (*Config, error) {
// 	var tmp Config
// 	if cp.Default != nil {
// 		tmp = *cp.Default
// 	} else {
// 		tmp = *DefaultConfig
// 	}
// 	err := cp.UpdateConfig(hclText, &tmp)
// 	return &tmp, err
// }

// func (cp ConfigParser) UpdateConfig(hclText string, result *Config) error {
// 	var tmp Config
// 	tmp = *result
// 	if err := hcl.Decode(&tmp, hclText); err != nil {
// 		return err
// 	}

// 	if tmp.InfluxDB.CheckLapse == 0 {
// 		tmp.InfluxDB.CheckLapse = DefaultConfig.InfluxDB.CheckLapse
// 	}

// 	if cp.AllowDNS && tmp.MesosDNS != nil {
// 		if _, err := NewDNSResolver(&tmp); err != nil {
// 			return err
// 		}
// 	}

// 	log.Printf("%+v\n", tmp)
// 	*result = tmp

// 	return nil
// }

// func (cp ConfigParser) ParseAndMerge(ihost *string, iport *int, idb *string) (*Config, error) {
// 	conf, err := cp.Parse()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if *ihost != EmptyString {
// 		conf.InfluxDB.Host = *ihost
// 	}
// 	if *iport != EmptyInt {
// 		conf.InfluxDB.Port = *iport
// 	}
// 	if *idb != EmptyString {
// 		conf.InfluxDB.DB = *idb
// 	}

// 	return conf, nil
// }
