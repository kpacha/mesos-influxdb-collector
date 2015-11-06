package mesos

import (
	"encoding/json"
	"github.com/kpacha/mesos-influxdb-collector/store"
	"io"
	"io/ioutil"
	"log"
	"time"
)

type SlaveParser struct {
	Node string
}

func (mp SlaveParser) Parse(r io.ReadCloser) ([]store.Point, error) {
	defer r.Close()
	var stats SlaveStats
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("Error reading from", r)
		return []store.Point{}, err
	}
	if err = json.Unmarshal(body, &stats); err != nil {
		log.Println("Error parsing to SlaveStats")
		return []store.Point{}, err
	}
	stats.Node = mp.Node
	stats.Time = time.Now()
	return mp.getMesosPoints(stats), nil
}

func (mp SlaveParser) getMesosPoints(stats SlaveStats) []store.Point {
	return []store.Point{
		mp.getCpuPoint(stats),
		mp.getDiskPoint(stats),
		mp.getMemPoint(stats),
		mp.getSystemPoint(stats),
		mp.getTasksPoint(stats),
		mp.getExecutorPoint(stats),
		mp.getGlobalPoint(stats),
	}
}

func (mp SlaveParser) getCpuPoint(stats SlaveStats) store.Point {
	return store.Point{
		Measurement: "slave-cpu",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"percent": stats.Slave_cpusPercent,
			"total":   stats.Slave_cpusTotal,
			"used":    stats.Slave_cpusUsed,
		},
		Time: stats.Time,
	}
}

func (mp SlaveParser) getDiskPoint(stats SlaveStats) store.Point {
	return store.Point{
		Measurement: "slave-disk",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"percent": stats.Slave_diskPercent,
			"total":   stats.Slave_diskTotal,
			"used":    stats.Slave_diskUsed,
		},
		Time: stats.Time,
	}
}

func (mp SlaveParser) getMemPoint(stats SlaveStats) store.Point {
	return store.Point{
		Measurement: "slave-mem",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"percent": stats.Slave_memPercent,
			"total":   stats.Slave_memTotal,
			"used":    stats.Slave_memUsed,
		},
		Time: stats.Time,
	}
}

func (mp SlaveParser) getTasksPoint(stats SlaveStats) store.Point {
	return store.Point{
		Measurement: "slave-tasks",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"failed":   stats.Slave_tasksFailed,
			"finished": stats.Slave_tasksFinished,
			"killed":   stats.Slave_tasksKilled,
			"lost":     stats.Slave_tasksLost,
			"running":  stats.Slave_tasksRunning,
			"staging":  stats.Slave_tasksStaging,
			"starting": stats.Slave_tasksStarting,
		},
		Time: stats.Time,
	}
}

func (mp SlaveParser) getSystemPoint(stats SlaveStats) store.Point {
	return store.Point{
		Measurement: "system",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"cpus_total":      stats.System_cpusTotal,
			"load_15min":      stats.System_load15min,
			"load_1min":       stats.System_load1min,
			"load_5min":       stats.System_load5min,
			"mem_free_bytes":  stats.System_memFreeBytes,
			"mem_total_bytes": stats.System_memTotalBytes,
		},
		Time: stats.Time,
	}
}

func (mp SlaveParser) getExecutorPoint(stats SlaveStats) store.Point {
	return store.Point{
		Measurement: "slave-executor",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"directory_max_allowed_age_secs": stats.Slave_executorDirectoryMaxAllowedAgeSecs,
			"registering":                    stats.Slave_executorsRegistering,
			"running":                        stats.Slave_executorsRunning,
			"terminated":                     stats.Slave_executorsTerminated,
			"terminating":                    stats.Slave_executorsTerminating,
		},
		Time: stats.Time,
	}
}

func (mp SlaveParser) getGlobalPoint(stats SlaveStats) store.Point {
	return store.Point{
		Measurement: "slave-global",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"registered":                 stats.Slave_registered,
			"invalid_framework_messages": stats.Slave_invalidFrameworkMessages,
			"invalid_status_updates":     stats.Slave_invalidStatusUpdates,
			"uptime_secs":                stats.Slave_uptimeSecs,
			"valid_framework_messages":   stats.Slave_validFrameworkMessages,
			"valid_status_updates":       stats.Slave_validStatusUpdates,
			"framewors":                  stats.Slave_frameworksActive,
			"conatiner_launch_errors":    stats.Slave_containerLaunchErrors,
		},
		Time: stats.Time,
	}
}

type SlaveStats struct {
	Containerizer_mesos_containerDestroyErrors int     `json:"containerizer/mesos/container_destroy_errors"`
	Slave_containerLaunchErrors                int     `json:"slave/container_launch_errors"`
	Slave_cpusPercent                          float64 `json:"slave/cpus_percent"`
	Slave_cpusRevocablePercent                 float64 `json:"slave/cpus_revocable_percent"`
	Slave_cpusRevocableTotal                   int     `json:"slave/cpus_revocable_total"`
	Slave_cpusRevocableUsed                    int     `json:"slave/cpus_revocable_used"`
	Slave_cpusTotal                            int     `json:"slave/cpus_total"`
	Slave_cpusUsed                             float64 `json:"slave/cpus_used"`
	Slave_diskPercent                          float64 `json:"slave/disk_percent"`
	Slave_diskRevocablePercent                 float64 `json:"slave/disk_revocable_percent"`
	Slave_diskRevocableTotal                   int     `json:"slave/disk_revocable_total"`
	Slave_diskRevocableUsed                    int     `json:"slave/disk_revocable_used"`
	Slave_diskTotal                            int     `json:"slave/disk_total"`
	Slave_diskUsed                             int     `json:"slave/disk_used"`
	Slave_executorDirectoryMaxAllowedAgeSecs   float64 `json:"slave/executor_directory_max_allowed_age_secs"`
	Slave_executorsRegistering                 int     `json:"slave/executors_registering"`
	Slave_executorsRunning                     int     `json:"slave/executors_running"`
	Slave_executorsTerminated                  int     `json:"slave/executors_terminated"`
	Slave_executorsTerminating                 int     `json:"slave/executors_terminating"`
	Slave_frameworksActive                     int     `json:"slave/frameworks_active"`
	Slave_invalidFrameworkMessages             int     `json:"slave/invalid_framework_messages"`
	Slave_invalidStatusUpdates                 int     `json:"slave/invalid_status_updates"`
	Slave_memPercent                           float64 `json:"slave/mem_percent"`
	Slave_memRevocablePercent                  float64 `json:"slave/mem_revocable_percent"`
	Slave_memRevocableTotal                    int     `json:"slave/mem_revocable_total"`
	Slave_memRevocableUsed                     int     `json:"slave/mem_revocable_used"`
	Slave_memTotal                             int     `json:"slave/mem_total"`
	Slave_memUsed                              int     `json:"slave/mem_used"`
	Slave_recoveryErrors                       int     `json:"slave/recovery_errors"`
	Slave_registered                           int     `json:"slave/registered"`
	Slave_tasksFailed                          int     `json:"slave/tasks_failed"`
	Slave_tasksFinished                        int     `json:"slave/tasks_finished"`
	Slave_tasksKilled                          int     `json:"slave/tasks_killed"`
	Slave_tasksLost                            int     `json:"slave/tasks_lost"`
	Slave_tasksRunning                         int     `json:"slave/tasks_running"`
	Slave_tasksStaging                         int     `json:"slave/tasks_staging"`
	Slave_tasksStarting                        int     `json:"slave/tasks_starting"`
	Slave_uptimeSecs                           float64 `json:"slave/uptime_secs"`
	Slave_validFrameworkMessages               int     `json:"slave/valid_framework_messages"`
	Slave_validStatusUpdates                   int     `json:"slave/valid_status_updates"`
	System_cpusTotal                           int     `json:"system/cpus_total"`
	System_load15min                           float64 `json:"system/load_15min"`
	System_load1min                            float64 `json:"system/load_1min"`
	System_load5min                            float64 `json:"system/load_5min"`
	System_memFreeBytes                        int     `json:"system/mem_free_bytes"`
	System_memTotalBytes                       int     `json:"system/mem_total_bytes"`
	Time                                       time.Time
	Node                                       string
}
