package store

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/influxdb/influxdb/client"
	"github.com/kpacha/mesos-influxdb-collector/config"
)

type Store interface {
	Store(points []Point) error
}

type Point struct {
	Measurement string
	Tags        map[string]string
	Time        time.Time
	Fields      map[string]interface{}
	Precision   string
	Raw         string
}

func (p Point) normalize() client.Point {
	return client.Point{
		Measurement: p.Measurement,
		Tags:        p.Tags,
		Time:        p.Time,
		Fields:      p.Fields,
		Precision:   p.Precision,
		Raw:         p.Raw,
	}
}

type InfluxdbConfig struct {
	Host       string
	Port       int
	DB         string
	Username   string
	Password   string
	CheckLapse int
}

type Influxdb struct {
	Connection *client.Client
	Config     InfluxdbConfig
}

func NewInfluxdbFromConfig(conf *config.Config, user, password string) Store {
	return NewInfluxdb(InfluxdbConfig{
		Host:       conf.InfluxDB.Host,
		Port:       conf.InfluxDB.Port,
		DB:         conf.InfluxDB.DB,
		Username:   user,
		Password:   password,
		CheckLapse: conf.InfluxDB.CheckLapse,
	})
}

func NewInfluxdb(conf InfluxdbConfig) Store {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", conf.Host, conf.Port))
	if err != nil {
		log.Fatal("Error creating the influxdb url: ", err)
	}

	connectionConf := client.Config{
		URL:      *u,
		Username: conf.Username,
		Password: conf.Password,
	}

	con, err := client.NewClient(connectionConf)
	if err != nil {
		log.Fatal("Error connecting to the influxdb store: ", err)
	}

	i := Influxdb{con, conf}

	go i.report()

	return i
}

func (i *Influxdb) report() {
	ticker := time.NewTicker(time.Second * time.Duration(i.Config.CheckLapse))
	for _ = range ticker.C {
		dur, ver, err := i.Connection.Ping()
		if err != nil {
			log.Fatal("Error pinging the influxdb store: ", err)
		}
		log.Printf("InfluxDb [%s] Ping: %v", ver, dur)
	}
}

func (i Influxdb) Store(points []Point) error {
	ps := make([]client.Point, len(points))
	for i, p := range points {
		ps[i] = p.normalize()
	}
	bps := client.BatchPoints{
		Points:          ps,
		Database:        i.Config.DB,
		RetentionPolicy: "default",
	}
	_, err := i.Connection.Write(bps)
	return err
}
