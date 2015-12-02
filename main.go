package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kpacha/mesos-influxdb-collector/collector"
	"github.com/kpacha/mesos-influxdb-collector/config"
	"github.com/kpacha/mesos-influxdb-collector/store"
)

const (
	ConfigPath   = "conf.hcl"
	InfluxdbHost = "localhost"
	InfluxdbPort = 8086
	InfluxdbDB   = "mesos"
	InfluxdbUser = "root"
	InfluxdbPass = "root"

	InfluxdbEnvName     = "INFLUXDB_HOST"
	InfluxdbDBEnvName   = "INFLUXDB_DB"
	InfluxdbPortEnvName = "INFLUXDB_PORT"
	ConfigPathEnvName   = "CONFIG_FILE"
	InfluxdbUserEnvName = "INFLUXDB_USER"
	InfluxdbPassEnvName = "INFLUXDB_PWD"
)

func main() {
	ihost := flag.String("Ih", getStringParam(InfluxdbEnvName, InfluxdbHost), "influxdb host")
	iport := flag.Int("Ip", getIntParam(InfluxdbPortEnvName, InfluxdbPort), "influxdb port")
	idb := flag.String("Id", getStringParam(InfluxdbDBEnvName, InfluxdbDB), "influxdb database")
	configPath := flag.String("c", ConfigPath, "path to the config file")
	flag.Parse()

	defaultConfig := *config.DefaultConfig
	defaultConfig.InfluxDB = &config.InfluxDB{
		Host:       *ihost,
		Port:       *iport,
		DB:         *idb,
		CheckLapse: config.DefaultConfig.InfluxDB.CheckLapse,
	}

	cp := config.ConfigParser{
		Path:     *configPath,
		AllowDNS: true,
		Default:  &defaultConfig,
	}

	conf, err := cp.Parse()
	if err != nil {
		log.Println("Error parsing config file:", err.Error())
		return
	}
	col := collector.NewCollectorFromConfig(conf)

	influxdb := store.NewInfluxdbFromConfig(
		conf,
		getStringParam(InfluxdbUserEnvName, InfluxdbUser),
		getStringParam(InfluxdbPassEnvName, InfluxdbPass))

	subscription := NewCollectorSubscription(&conf.Lapse, &col, &influxdb)

	go report(&subscription, conf.InfluxDB.CheckLapse)

	time.Sleep(time.Second * time.Duration(conf.DieAfter))
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

func report(subscription *Subscription, lapse int) {
	ticker := time.NewTicker(time.Second * time.Duration(lapse))
	var collects int
	for _ = range ticker.C {
		collects = <-subscription.Stats
		log.Println("Total collected stats:", collects)
	}
}
