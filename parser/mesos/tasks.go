package mesos

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/kpacha/mesos-influxdb-collector/store"
)

type SlaveStatsParser struct {
	Node string
}

func (mp SlaveStatsParser) Parse(r io.ReadCloser) ([]store.Point, error) {
	defer r.Close()
	var stats []SlaveTaskStats
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("Error reading from", r)
		return []store.Point{}, err
	}
	if err = json.Unmarshal(body, &stats); err != nil {
		log.Println("Error parsing to SlaveTaskStats")
		return []store.Point{}, err
	}
	return mp.getTaskPoints(stats), nil
}

func (mp SlaveStatsParser) getTaskPoints(stats []SlaveTaskStats) []store.Point {
	ts := time.Now()
	points := []store.Point{}
	for _, taskStats := range stats {
		points = append(points, mp.getTaskPoint(taskStats, ts))
	}
	return points
}

func (mp SlaveStatsParser) getTaskPoint(stats SlaveTaskStats, ts time.Time) store.Point {
	return store.Point{
		Measurement: "task-consumption",
		Tags: map[string]string{
			"node":     mp.Node,
			"executor": stats.ExecutorID,
		},
		Fields: map[string]interface{}{
			"cpu_limit": stats.Statistics.CpusLimit,
			"cpu_sys":   stats.Statistics.CpusSystemTimeSecs,
			"cpu_usr":   stats.Statistics.CpusUserTimeSecs,
			"mem_lim":   stats.Statistics.MemLimitBytes,
			"mem_rss":   stats.Statistics.MemRssBytes,
		},
		Time: ts,
	}
}

type SlaveTaskStats struct {
	ExecutorID string `json:"executor_id"`
	Statistics struct {
		CpusLimit          float64 `json:"cpus_limit"`
		CpusSystemTimeSecs float64 `json:"cpus_system_time_secs"`
		CpusUserTimeSecs   float64 `json:"cpus_user_time_secs"`
		MemLimitBytes      int     `json:"mem_limit_bytes"`
		MemRssBytes        int     `json:"mem_rss_bytes"`
	} `json:"statistics"`
	Time time.Time
	Node string
}
