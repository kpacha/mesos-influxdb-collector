package main

import (
	"flag"
	"github.com/kpacha/mesos-influxdb-collector/collector"
	"github.com/kpacha/mesos-influxdb-collector/store"
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
	MesosMasterHost = "localhost"
	MesosMasterPort = 5050
	MesosSlaveHost  = "localhost"
	MesosSlavePort  = 5051
	DefaultLapse    = 1
	DefaultLifeTime = 300
	DefaultLogLapse = 30

	InfluxdbEnvName        = "INFLUXDB_HOST"
	InfluxdbDBEnvName      = "INFLUXDB_DB"
	InfluxdbPortEnvName    = "INFLUXDB_PORT"
	InfluxdbUserEnvName    = "INFLUXDB_USER"
	InfluxdbPassEnvName    = "INFLUXDB_PWD"
	MesosMasterHostEnvName = "MESOS_MASTER_HOST"
	MesosMasterPortEnvName = "MESOS_MASTER_PORT"
	MesosSlaveHostEnvName  = "MESOS_SLAVE_HOST"
	MesosSlavePortEnvName  = "MESOS_SLAVE_PORT"
	LapseEnvName           = "COLLECTOR_LAPSE"
	LifeTimeEnvName        = "COLLECTOR_LIFETIME"
)

func main() {
	ihost := flag.String("Ih", getStringParam(InfluxdbEnvName, InfluxdbHost), "influxdb host")
	iport := flag.Int("Ip", getIntParam(InfluxdbPortEnvName, InfluxdbPort), "influxdb port")
	idb := flag.String("Id", getStringParam(InfluxdbDBEnvName, InfluxdbDB), "influxdb database")
	mmhost := flag.String("Mmh", getStringParam(MesosMasterHostEnvName, MesosMasterHost), "mesos master host")
	mmport := flag.Int("Mmp", getIntParam(MesosMasterPortEnvName, MesosMasterPort), "mesos master port")
	mshost := flag.String("Msh", getStringParam(MesosSlaveHostEnvName, MesosSlaveHost), "mesos slave host")
	msport := flag.Int("Msp", getIntParam(MesosSlavePortEnvName, MesosSlavePort), "mesos slave port")
	lapse := flag.Int("l", getIntParam(LapseEnvName, DefaultLapse), "sleep time between collections in seconds")
	dieAfter := flag.Int("d", getIntParam(LifeTimeEnvName, DefaultLifeTime), "die after N seconds")
	flag.Parse()

	influxdb := store.NewInfluxdb(store.InfluxdbConfig{
		Host:       *ihost,
		Port:       *iport,
		DB:         *idb,
		Username:   getStringParam(InfluxdbUserEnvName, InfluxdbUser),
		Password:   getStringParam(InfluxdbPassEnvName, InfluxdbPass),
		CheckLapse: DefaultLogLapse,
	})

	col := collector.NewMultiCollector(
		[]collector.Collector{NewMesosMasterCollector(*mmhost, *mmport),
			NewMesosSlaveCollector(*mshost, *msport)})

	subscription := NewCollectorSubscription(lapse, &col, &influxdb)

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
