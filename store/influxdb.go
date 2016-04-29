package store

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
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

func (p Point) normalize() *client.Point {
	pt, err := client.NewPoint(p.Measurement, p.Tags, p.Fields, p.Time)
	if err != nil {
		panic(err)
	}
	return pt
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
	Connection client.Client
	Config     InfluxdbConfig
}

func NewInfluxdbFromConfig(conf *config.Config) Store {
	return NewInfluxdb(InfluxdbConfig{
		Host:       conf.InfluxDB.Host,
		Port:       conf.InfluxDB.Port,
		DB:         conf.InfluxDB.DB,
		Username:   conf.InfluxDB.User,
		Password:   conf.InfluxDB.Pass,
		CheckLapse: conf.InfluxDB.CheckLapse,
	})
}

func NewInfluxdb(conf InfluxdbConfig) Store {
	u := fmt.Sprintf("http://%s:%d", conf.Host, conf.Port)
	fmt.Println("Connecting to:", u)
	con, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     u,
		Username: conf.Username,
		Password: conf.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}

	i := Influxdb{con, conf}
	go i.report()
	return i
}

func (i *Influxdb) report() {
	ticker := time.NewTicker(time.Second * time.Duration(i.Config.CheckLapse))
	for _ = range ticker.C {
		dur, ver, err := i.Connection.Ping(10 * time.Second)
		if err != nil {
			log.Fatal("Error pinging the influxdb store: ", err)
		}
		log.Printf("InfluxDb Ping[%s]: %v", ver, dur)
	}
}

func (i Influxdb) Store(points []Point) error {
	bps, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Config.DB,
		Precision: "s",
	})
	if err != nil {
		panic(err)
	}
	for _, p := range points {
		bps.AddPoint(p.normalize())
	}
	err = i.Connection.Write(bps)
	return err
}
