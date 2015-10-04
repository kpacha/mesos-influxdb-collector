package main

import (
	"log"
	"time"
)

type Subscription struct {
	close chan bool
	Stats chan int
}

func NewCollectorSubscription(lapse *int, collector *Collector, store *Store) Subscription {
	s := Subscription{make(chan bool), make(chan int)}
	go s.manageCollects(lapse, collector, store)
	return s
}

func (s *Subscription) Cancel() {
	s.close <- true
}

func (s *Subscription) manageCollects(lapse *int, collector *Collector, store *Store) {
	ticker := time.NewTicker(time.Second * time.Duration(*lapse))
	collects := 0
	for {
		select {
		case <-ticker.C:
			go s.collect(collector, store)
			collects++
		case <-s.close:
			ticker.Stop()
			return
		case s.Stats <- collects:
			continue
		}
	}
}

func (s *Subscription) collect(collector *Collector, store *Store) {
	var stats Stats
	if err := (*collector).Collect(&stats); err != nil {
		log.Fatal("Error collecting stats: ", err)
	}

	if err := (*store).Store(&stats); err != nil {
		log.Fatal("Error storing stats: ", err)
	}
}
