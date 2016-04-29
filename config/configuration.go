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

	c = cp.updateNestedDefaults(&c)

	if cp.AllowDNS && c.MesosDNS != nil {
		if _, err := NewDNSResolver(&c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func (cp *ConfigParser) updateNestedDefaults(c *Config) Config {
	var tmp Config
	tmp = *c

	tmp.InfluxDB.CheckLapse = cp.Cfg.GetInt("influxdb.checkLapse")
	tmp.InfluxDB.Port = cp.Cfg.GetInt("influxdb.port")
	tmp.InfluxDB.Host = cp.Cfg.GetString("influxdb.host")
	tmp.InfluxDB.DB = cp.Cfg.GetString("influxdb.db")
	tmp.InfluxDB.User = cp.Cfg.GetString("influxdb.user")
	tmp.InfluxDB.Pass = cp.Cfg.GetString("influxdb.pass")

	if tmp.HAProxy != nil {
		tmp.HAProxy.EndPoint = cp.Cfg.GetString("haproxy.endPoint")
		tmp.HAProxy.Port = cp.Cfg.GetInt("haproxy.port")
		tmp.HAProxy.User = cp.Cfg.GetString("haproxy.user")
		tmp.HAProxy.Password = cp.Cfg.GetString("haproxy.password")
	}

	log.Printf("%+v\n", tmp)
	return tmp
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

	return cfg
}
