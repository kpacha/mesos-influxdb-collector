package config

import (
	"fmt"
	"log"
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
	Port     int
}

type ConfigParser struct {
	Cfg      *viper.Viper
	AllowDNS bool
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

func NewConfigParser(format, path, configName string, allowDNS bool) *ConfigParser {
	return &ConfigParser{newViper(format, path, configName), allowDNS}
}

func (cp *ConfigParser) Parse() (*Config, error) {
	err := cp.Cfg.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	var c Config
	err = cp.Cfg.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	c = cp.updateNestedValues(&c)

	if cp.AllowDNS && c.MesosDNS != nil {
		dns := NewDNSResolver(&c)
		if err = dns.resolve(); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func (cp *ConfigParser) updateNestedValues(c *Config) Config {
	var tmp Config
	tmp = *c

	tmp.InfluxDB.CheckLapse = cp.getNestedInt("influxdb.checkLapse", DefaultInfluxdb.CheckLapse)
	tmp.InfluxDB.Port = cp.getNestedInt("influxdb.port", DefaultInfluxdb.Port)
	tmp.InfluxDB.Host = cp.getNestedString("influxdb.host", DefaultInfluxdb.Host)
	tmp.InfluxDB.DB = cp.getNestedString("influxdb.db", DefaultInfluxdb.DB)
	tmp.InfluxDB.User = cp.getNestedString("influxdb.user", DefaultInfluxdb.User)
	tmp.InfluxDB.Pass = cp.getNestedString("influxdb.pass", DefaultInfluxdb.Pass)

	if tmp.HAProxy != nil {
		tmp.HAProxy.EndPoint = cp.Cfg.GetString("haproxy.endPoint")
		tmp.HAProxy.Port = cp.Cfg.GetInt("haproxy.port")
		tmp.HAProxy.User = cp.Cfg.GetString("haproxy.user")
		tmp.HAProxy.Password = cp.Cfg.GetString("haproxy.password")
	}

	log.Printf("%+v\n", tmp)
	return tmp
}

func (cp *ConfigParser) getNestedString(key, defaultValue string) string {
	value := cp.Cfg.GetString(key)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func (cp *ConfigParser) getNestedInt(key string, defaultValue int) int {
	value := cp.Cfg.GetInt(key)
	if value == 0 {
		return defaultValue
	} else {
		return value
	}
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

	cfg.SetDefault("lapse", DefaultLapse)
	cfg.SetDefault("dieAfter", DefaultDieAfter)
	cfg.SetDefault("influxdb", DefaultInfluxdb)

	return cfg
}
