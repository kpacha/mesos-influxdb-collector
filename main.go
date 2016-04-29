package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/kpacha/mesos-influxdb-collector/collector"
	"github.com/kpacha/mesos-influxdb-collector/config"
	"github.com/kpacha/mesos-influxdb-collector/store"
)

const (
	DefaultConfigPath      = "."
	DefaultConfigFile      = "conf"
	DefaultFormat          = "json"
	DefaultMesosDNSEnabled = true

	FormatDBEnvName     = "FORMAT"
	ConfigPathDBEnvName = "CONFIG_PATH"
	ConfigFileDBEnvName = "CONFIG_FILE"
)

func main() {
	allowDNS := flag.Bool("dns", DefaultMesosDNSEnabled, "enable mesos-dns")
	format := flag.String("f", getStringParam(FormatDBEnvName, DefaultFormat), "config format")
	configPath := flag.String("d", getStringParam(ConfigPathDBEnvName, DefaultConfigPath), "path to the config folder")
	configFile := flag.String("c", getStringParam(ConfigFileDBEnvName, DefaultConfigFile), "name of the config file")
	flag.Parse()

	cp := config.NewConfigParser(*format, *configPath, *configFile, *allowDNS)
	conf, err := cp.Parse()
	if err != nil {
		log.Println("Error parsing config file:", err.Error())
		return
	}
	col := collector.NewCollectorFromConfig(conf)

	influxdb := store.NewInfluxdbFromConfig(conf)

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

func report(subscription *Subscription, lapse int) {
	ticker := time.NewTicker(time.Second * time.Duration(lapse))
	var collects int
	for _ = range ticker.C {
		collects = <-subscription.Stats
		log.Println("Total collected stats:", collects)
	}
}
