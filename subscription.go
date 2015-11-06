package main

import (
	"github.com/kpacha/mesos-influxdb-collector/collector"
	"github.com/kpacha/mesos-influxdb-collector/store"
	"log"
	"time"
)

type Subscription struct {
	close chan bool
	Stats chan int
}

func NewCollectorSubscription(lapse *int, collector *collector.Collector, st *store.Store) Subscription {
	s := Subscription{make(chan bool), make(chan int)}
	go s.manageCollects(lapse, collector, st)
	return s
}

func (s *Subscription) Cancel() {
	s.close <- true
}

func (s *Subscription) manageCollects(lapse *int, collector *collector.Collector, st *store.Store) {
	ticker := time.NewTicker(time.Second * time.Duration(*lapse))
	collects := 0
	for {
		select {
		case <-ticker.C:
			go s.collect(collector, st)
			collects++
		case <-s.close:
			ticker.Stop()
			return
		case s.Stats <- collects:
			continue
		}
	}
}

func (s *Subscription) collect(collector *collector.Collector, st *store.Store) {
	points, err := (*collector).Collect()
	if err != nil {
		log.Fatal("Error collecting stats: ", err)
	}

	if err = (*st).Store(points); err != nil {
		log.Fatal("Error storing stats: ", err)
	}

	log.Printf("Collection completed! Collected %d points from %s", len(points), *collector)
}
