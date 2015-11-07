package config

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/kpacha/mesos-influxdb-collector/reader"
)

type SRVRecord struct {
	Host    string `json:"host"`
	IP      string `json:"ip"`
	Port    string `json:"port"`
	Service string `json:"service"`
}

type ARecord struct {
	Host string `json:"host"`
	IP   string `json:"ip"`
}

type DNSResolver struct {
	Config *Config
}

func NewDNSResolver(config *Config) (*DNSResolver, error) {
	resolver := DNSResolver{config}
	if err := resolver.resolveMesosMasters(); err != nil {
		log.Println("Error resolving MesosMasters")
		return nil, err
	}
	if err := resolver.resolveMesosSlaves(); err != nil {
		log.Println("Error resolving MesosSlaves")
		return nil, err
	}
	if err := resolver.resolveMarathon(); err != nil {
		log.Println("Error resolving Marathon")
		return nil, err
	}
	return &resolver, nil
}

func (r DNSResolver) resolveMesosMasters() error {
	body, err := reader.ReadUrl(r.getMesosMasterUrl())
	if err != nil {
		return err
	}

	var masters []SRVRecord
	if err = json.Unmarshal(body, &masters); err != nil {
		return err
	}

	for _, master := range masters {
		port, err := strconv.Atoi(master.Port)
		if err != nil {
			return err
		}
		r.Config.Master = append(r.Config.Master, Master{master.IP, port, true})
	}

	return nil
}

func (r DNSResolver) resolveMesosSlaves() error {
	body, err := reader.ReadUrl(r.getMesosSlaveUrl())
	if err != nil {
		return err
	}

	var slaves []ARecord
	if err = json.Unmarshal(body, &slaves); err != nil {
		return err
	}

	for _, slave := range slaves {
		r.Config.Slave = append(r.Config.Slave, Server{slave.IP, 5051})
	}

	return nil
}

func (r DNSResolver) resolveMarathon() error {
	body, err := reader.ReadUrl(r.getMarathonUrl())
	if err != nil {
		return err
	}

	var instances []ARecord
	if err = json.Unmarshal(body, &instances); err != nil {
		return err
	}

	for _, instance := range instances {
		r.Config.Marathon = append(r.Config.Marathon, Server{instance.IP, 8080})
	}

	return nil
}

func (r DNSResolver) getMesosMasterUrl() string {
	//return r.getUrl("/v1/hosts/master")
	return r.getUrl("v1/services/_leader._tcp")
}

func (r DNSResolver) getMesosSlaveUrl() string {
	return r.getUrl("v1/hosts/slave")
}

func (r DNSResolver) getMarathonUrl() string {
	return r.getUrl("v1/hosts/marathon")
}

func (r DNSResolver) getUrl(partial string) string {
	return fmt.Sprintf("http://%s:%d/%s.%s.", r.Config.MesosDNS.Host, r.Config.MesosDNS.Port, partial, r.Config.MesosDNS.Domain)
}
