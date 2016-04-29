package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kpacha/mesos-influxdb-collector/reader"
)

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
	leaders, err := r.getARecords(r.getMesosLeaderUrl())
	if err != nil {
		return err
	}
	masters, err := r.getARecords(r.getMesosMasterUrl())
	if err != nil {
		return err
	}

	log.Println("leaders", leaders)
	log.Println("masters", masters)

	for _, master := range masters {
		r.Config.Master = append(r.Config.Master, Master{master.IP, 5050, master.IP == leaders[0].IP})
	}

	return nil
}

func (r DNSResolver) resolveMesosSlaves() error {
	slaves, err := r.getARecords(r.getMesosSlaveUrl())
	if err != nil {
		return err
	}

	for _, slave := range slaves {
		r.Config.Slave = append(r.Config.Slave, Server{slave.IP, 5051})
	}

	return nil
}

func (r DNSResolver) resolveMarathon() error {
	if !r.Config.MesosDNS.Marathon {
		return nil
	}
	instances, err := r.getARecords(r.getMarathonUrl())
	if err != nil {
		return err
	}

	for _, instance := range instances {
		r.Config.Marathon.Server = append(r.Config.Marathon.Server, Server{instance.IP, 8080})
	}

	return nil
}

func (r DNSResolver) getARecords(url string) ([]ARecord, error) {
	var instances []ARecord
	body, err := reader.ReadUrl(url)
	if err != nil {
		return instances, err
	}

	err = json.Unmarshal(body, &instances)
	return instances, err
}

func (r DNSResolver) getMesosMasterUrl() string {
	return r.getUrl("v1/hosts/master")
}

func (r DNSResolver) getMesosLeaderUrl() string {
	return r.getUrl("v1/hosts/leader")
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
