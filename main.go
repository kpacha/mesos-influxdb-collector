package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	InfluxdbHost    = "localhost"
	InfluxdbPort    = 8086
	InfluxdbDB      = "mesos"
	InfluxdbUser    = "root"
	InfluxdbPass    = "root"
	MesosHost       = "localhost"
	MesosPort       = 5050
	DefaultLapse    = 1
	DefaultLifeTime = 300
	DefaultLogLapse = 30

	InfluxdbEnvName     = "INFLUXDB_HOST"
	InfluxdbDBEnvName   = "INFLUXDB_DB"
	InfluxdbPortEnvName = "INFLUXDB_PORT"
	InfluxdbUserEnvName = "INFLUXDB_USER"
	InfluxdbPassEnvName = "INFLUXDB_PWD"
	MesosHostEnvName    = "MESOS_HOST"
	MesosPortEnvName    = "MESOS_PORT"
	LapseEnvName        = "COLLECTOR_LAPSE"
	LifeTimeEnvName     = "COLLECTOR_LIFETIME"
)

func main() {
	ihost := flag.String("Ih", getStringParam(InfluxdbEnvName, InfluxdbHost), "influxdb host")
	iport := flag.Int("Ip", getIntParam(InfluxdbPortEnvName, InfluxdbPort), "influxdb port")
	idb := flag.String("Id", getStringParam(InfluxdbDBEnvName, InfluxdbDB), "influxdb database")
	mhost := flag.String("Mh", getStringParam(MesosHostEnvName, MesosHost), "mesos host")
	mport := flag.Int("Mp", getIntParam(MesosPortEnvName, MesosPort), "mesos port")
	lapse := flag.Int("l", getIntParam(LapseEnvName, DefaultLapse), "sleep time between collections in seconds")
	dieAfter := flag.Int("d", getIntParam(LifeTimeEnvName, DefaultLifeTime), "die after N seconds")
	flag.Parse()

	influxdb := NewInfluxdb(InfluxdbConfig{
		Host:       *ihost,
		Port:       *iport,
		DB:         *idb,
		Username:   getStringParam(InfluxdbUserEnvName, InfluxdbUser),
		Password:   getStringParam(InfluxdbPassEnvName, InfluxdbPass),
		CheckLapse: DefaultLogLapse,
	})

	mesos := NewMesosCollector(*mhost, *mport)

	subscription := NewCollectorSubscription(lapse, &mesos, &influxdb)

	go report(&subscription)

	time.Sleep(time.Second * time.Duration(*dieAfter))
	subscription.Cancel()

	log.Println("Mesos collector stopped")

}

func getStringParam(envName string, defaultValue string) string {
	env := os.Getenv(envName)
	if env == "" {
		return defaultValue
	}
	return env
}

func getIntParam(envName string, defaultValue int) int {
	env, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		return defaultValue
	}
	return env
}

func report(subscription *Subscription) {
	ticker := time.NewTicker(time.Second * time.Duration(DefaultLogLapse))
	var collects int
	for _ = range ticker.C {
		collects = <-subscription.Stats
		log.Println("Total collected stats:", collects)
	}
}
