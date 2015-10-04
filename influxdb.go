package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/influxdb/influxdb/client"
)

type InfluxdbConfig struct {
	Host       string
	Port       int
	DB         string
	Username   string
	Password   string
	CheckLapse int
}

type Store interface {
	Store(stats *Stats) error
}

type Influxdb struct {
	Connection *client.Client
	Config     InfluxdbConfig
}

func NewInfluxdb(conf InfluxdbConfig) Store {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", conf.Host, conf.Port))
	if err != nil {
		log.Fatal("Error creating the influxdb url: ", err)
	}

	connectionConf := client.Config{
		URL:      *u,
		Username: conf.Username,
		Password: conf.Password,
	}

	con, err := client.NewClient(connectionConf)
	if err != nil {
		log.Fatal("Error connecting to the influxdb store: ", err)
	}

	i := Influxdb{con, conf}

	go i.report()

	return i
}

func (i *Influxdb) report() {
	ticker := time.NewTicker(time.Second * time.Duration(i.Config.CheckLapse))
	for _ = range ticker.C {
		dur, ver, err := i.Connection.Ping()
		if err != nil {
			log.Fatal("Error pinging the influxdb store: ", err)
		}
		log.Printf("InfluxDb [%s] Ping: %v", ver, dur)
	}
}

func (i Influxdb) Store(stats *Stats) error {
	pts := append(i.getMastersPoints(stats), i.getMesosPoints(stats)...)

	bps := client.BatchPoints{
		Points:          pts,
		Database:        i.Config.DB,
		RetentionPolicy: "default",
	}
	_, err := i.Connection.Write(bps)
	return err
}

func (i *Influxdb) getMesosPoints(stats *Stats) []client.Point {
	return []client.Point{
		i.getCpuPoint(stats),
		i.getDiskPoint(stats),
		i.getMemPoint(stats),
		i.getSystemPoint(stats),
		//i.getRegistrarPoint(stats),
		i.getTasksPoint(stats),
	}
}

func (i *Influxdb) getMastersPoints(stats *Stats) []client.Point {
	return []client.Point{
		i.getMastersCpuPoint(stats),
		i.getMastersDiskPoint(stats),
		i.getMastersMemPoint(stats),
		i.getMastersFrameworksPoint(stats),
		i.getMastersTasksPoint(stats),
		i.getMastersSlavesPoint(stats),
		i.getMastersEventQueuePoint(stats),
		i.getMastersMessagesPoint(stats),
		i.getMastersGlobalPoint(stats),
	}
}

func (i *Influxdb) getCpuPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "cpu",
		Fields: map[string]interface{}{
			"percent": stats.CpusPercent,
			"total":   stats.CpusTotal,
			"used":    stats.CpusUsed,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getDiskPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "disk",
		Fields: map[string]interface{}{
			"percent": stats.DiskPercent,
			"total":   stats.DiskTotal,
			"used":    stats.DiskUsed,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMemPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "mem",
		Fields: map[string]interface{}{
			"percent": stats.MemPercent,
			"total":   stats.MemTotal,
			"used":    stats.MemUsed,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getTasksPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "tasks",
		Fields: map[string]interface{}{
			"failed":   stats.FailedTasks,
			"finished": stats.FinishedTasks,
			"killed":   stats.KilledTasks,
			"lost":     stats.LostTasks,
			"staging":  stats.StagedTasks,
			"starting": stats.StartedTasks,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getGlobalPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "global",
		Fields: map[string]interface{}{
			"activated_slaves":       stats.ActivatedSlaves,
			"active_schedulers":      stats.ActiveSchedulers,
			"active_tasks_gauge":     stats.ActiveTasksGauge,
			"deactivated_slaves":     stats.DeactivatedSlaves,
			"elected":                stats.Elected,
			"invalid_status_updates": stats.InvalidStatusUpdates,
			"outstanding_offers":     stats.OutstandingOffers,
			"total_schedulers":       stats.TotalSchedulers,
			"uptime":                 stats.Uptime,
			"valid_status_updates":   stats.ValidStatusUpdates,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getSystemPoint(stats *Stats) client.Point {
	return client.Point{
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

func (i *Influxdb) getMastersCpuPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.cpu",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"percent": stats.Master_cpusPercent,
			"total":   stats.Master_cpusTotal,
			"used":    stats.Master_cpusUsed,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersDiskPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.disk",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"percent": stats.Master_diskPercent,
			"total":   stats.Master_diskTotal,
			"used":    stats.Master_diskUsed,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersMemPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.mem",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"percent": stats.Master_memPercent,
			"total":   stats.Master_memTotal,
			"used":    stats.Master_memUsed,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersFrameworksPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.frameworks",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"active":       stats.Master_frameworksActive,
			"connected":    stats.Master_frameworksConnected,
			"disconnected": stats.Master_frameworksDisconnected,
			"inactive":     stats.Master_frameworksInactive,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersTasksPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.tasks",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"error":    stats.Master_tasksError,
			"failed":   stats.Master_tasksFailed,
			"finished": stats.Master_tasksFinished,
			"killed":   stats.Master_tasksKilled,
			"lost":     stats.Master_tasksLost,
			"running":  stats.Master_tasksRunning,
			"staging":  stats.Master_tasksStaging,
			"starting": stats.Master_tasksStarting,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersSlavesPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.slaves",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"recovery_slave_removals":   stats.Master_recoverySlaveRemovals,
			"slave_registrations":       stats.Master_slaveRegistrations,
			"slave_removals":            stats.Master_slaveRemovals,
			"slave_reregistrations":     stats.Master_slaveReregistrations,
			"slave_shutdowns_canceled":  stats.Master_slaveShutdownsCanceled,
			"slave_shutdowns_scheduled": stats.Master_slaveShutdownsScheduled,
			"slaves_active":             stats.Master_slavesActive,
			"slaves_connected":          stats.Master_slavesConnected,
			"slaves_disconnected":       stats.Master_slavesDisconnected,
			"slaves_inactive":           stats.Master_slavesInactive,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersEventQueuePoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.event_queue",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"dispatches":    stats.Master_eventQueueDispatches,
			"http_requests": stats.Master_eventQueueHTTPRequests,
			"messages":      stats.Master_eventQueueMessages,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersMessagesPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.messages",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"dropped_messages":                       stats.Master_droppedMessages,
			"messages_authenticate":                  stats.Master_messagesAuthenticate,
			"messages_deactivate_framework":          stats.Master_messagesDeactivateFramework,
			"messages_decline_offers":                stats.Master_messagesDeclineOffers,
			"messages_exited_executor":               stats.Master_messagesExitedExecutor,
			"messages_framework_to_executor":         stats.Master_messagesFrameworkToExecutor,
			"messages_kill_task":                     stats.Master_messagesKillTask,
			"messages_launch_tasks":                  stats.Master_messagesLaunchTasks,
			"messages_reconcile_tasks":               stats.Master_messagesReconcileTasks,
			"messages_register_framework":            stats.Master_messagesRegisterFramework,
			"messages_register_slave":                stats.Master_messagesRegisterSlave,
			"messages_reregister_framework":          stats.Master_messagesReregisterFramework,
			"messages_reregister_slave":              stats.Master_messagesReregisterSlave,
			"messages_resource_request":              stats.Master_messagesResourceRequest,
			"messages_revive_offers":                 stats.Master_messagesReviveOffers,
			"messages_status_update":                 stats.Master_messagesStatusUpdate,
			"messages_status_update_acknowledgement": stats.Master_messagesStatusUpdateAcknowledgement,
			"messages_unregister_framework":          stats.Master_messagesUnregisterFramework,
			"messages_unregister_slave":              stats.Master_messagesUnregisterSlave,
		},
		Time: stats.Time,
	}
}

func (i *Influxdb) getMastersGlobalPoint(stats *Stats) client.Point {
	return client.Point{
		Measurement: "master.global",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"elected": stats.Master_elected,
			"invalid_framework_to_executor_messages": stats.Master_invalidFrameworkToExecutorMessages,
			"invalid_status_update_acknowledgements": stats.Master_invalidStatusUpdateAcknowledgements,
			"invalid_status_updates":                 stats.Master_invalidStatusUpdates,
			"outstanding_offers":                     stats.Master_outstandingOffers,
			"uptime_secs":                            stats.Master_uptimeSecs,
			"valid_framework_to_executor_messages":   stats.Master_validFrameworkToExecutorMessages,
			"valid_status_update_acknowledgements":   stats.Master_validStatusUpdateAcknowledgements,
			"valid_status_updates":                   stats.Master_validStatusUpdates,
		},
		Time: stats.Time,
	}
}
