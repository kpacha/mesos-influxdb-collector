package mesos

import (
	"encoding/json"
	"github.com/kpacha/mesos-influxdb-collector/store"
	"io"
	"io/ioutil"
	"log"
	"time"
)

type MasterParser struct {
	Node   string
	Leader bool
}

func (mp MasterParser) Parse(r io.ReadCloser) ([]store.Point, error) {
	defer r.Close()
	var stats MasterStats
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("Error reading from", r)
		return []store.Point{}, err
	}
	if err = json.Unmarshal(body, &stats); err != nil {
		log.Println("Error parsing to MasterStats")
		return []store.Point{}, err
	}
	stats.Time = time.Now()
	return mp.getMesosPoints(stats), nil
}

func (mp MasterParser) getMesosPoints(stats MasterStats) []store.Point {
	ps := []store.Point{mp.getSystemPoint(stats)}
	if mp.Leader {
		ps = append(ps,
			mp.getCpuPoint(stats),
			mp.getDiskPoint(stats),
			mp.getMemPoint(stats),
			//mp.getRegistrarPoint(stats),
			mp.getTasksPoint(stats),
			mp.getFrameworksPoint(stats),
			mp.getSlavesPoint(stats),
			mp.getEventQueuePoint(stats),
			mp.getMessagesPoint(stats),
			mp.getGlobalPoint(stats))
	}
	return ps
}

func (mp MasterParser) getCpuPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "cpu",
		Fields: map[string]interface{}{
			"percent": stats.Master_cpusPercent,
			"total":   stats.Master_cpusTotal,
			"used":    stats.Master_cpusUsed,
		},
		Time: stats.Time,
	}
}

func (mp MasterParser) getDiskPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "disk",
		Fields: map[string]interface{}{
			"percent": stats.Master_diskPercent,
			"total":   stats.Master_diskTotal,
			"used":    stats.Master_diskUsed,
		},
		Time: stats.Time,
	}
}

func (mp MasterParser) getMemPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "mem",
		Fields: map[string]interface{}{
			"percent": stats.Master_memPercent,
			"total":   stats.Master_memTotal,
			"used":    stats.Master_memUsed,
		},
		Time: stats.Time,
	}
}

func (mp MasterParser) getTasksPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "tasks",
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

func (mp MasterParser) getSystemPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "system",
		Tags: map[string]string{
			"node": mp.Node,
		},
		Fields: map[string]interface{}{
			"cpus_total":      stats.System_cpusTotal,
			"load_15min":      stats.System_load15min,
			"load_1min":       stats.System_load1min,
			"load_5min":       stats.System_load5min,
			"mem_free_bytes":  int(stats.System_memFreeBytes),
			"mem_total_bytes": int(stats.System_memTotalBytes),
		},
		Time: stats.Time,
	}
}

func (mp MasterParser) getFrameworksPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "frameworks",
		Fields: map[string]interface{}{
			"active":       stats.Master_frameworksActive,
			"connected":    stats.Master_frameworksConnected,
			"disconnected": stats.Master_frameworksDisconnected,
			"inactive":     stats.Master_frameworksInactive,
		},
		Time: stats.Time,
	}
}

func (mp MasterParser) getSlavesPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "slaves",
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

func (mp MasterParser) getEventQueuePoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "event_queue",
		Fields: map[string]interface{}{
			"dispatches":    stats.Master_eventQueueDispatches,
			"http_requests": stats.Master_eventQueueHTTPRequests,
			"messages":      stats.Master_eventQueueMessages,
		},
		Time: stats.Time,
	}
}

func (mp MasterParser) getMessagesPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "messages",
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

func (mp MasterParser) getGlobalPoint(stats MasterStats) store.Point {
	return store.Point{
		Measurement: "global",
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

type MasterStats struct {
	Allocator_eventQueueDispatches             float64 `json:"allocator/event_queue_dispatches"`
	Master_cpusPercent                         float64 `json:"master/cpus_percent"`
	Master_cpusRevocablePercent                float64 `json:"master/cpus_revocable_percent"`
	Master_cpusRevocableTotal                  float64 `json:"master/cpus_revocable_total"`
	Master_cpusRevocableUsed                   float64 `json:"master/cpus_revocable_used"`
	Master_cpusTotal                           float64 `json:"master/cpus_total"`
	Master_cpusUsed                            float64 `json:"master/cpus_used"`
	Master_diskPercent                         float64 `json:"master/disk_percent"`
	Master_diskRevocablePercent                float64 `json:"master/disk_revocable_percent"`
	Master_diskRevocableTotal                  float64 `json:"master/disk_revocable_total"`
	Master_diskRevocableUsed                   float64 `json:"master/disk_revocable_used"`
	Master_diskTotal                           float64 `json:"master/disk_total"`
	Master_diskUsed                            float64 `json:"master/disk_used"`
	Master_droppedMessages                     float64 `json:"master/dropped_messages"`
	Master_elected                             float64 `json:"master/elected"`
	Master_eventQueueDispatches                float64 `json:"master/event_queue_dispatches"`
	Master_eventQueueHTTPRequests              float64 `json:"master/event_queue_http_requests"`
	Master_eventQueueMessages                  float64 `json:"master/event_queue_messages"`
	Master_frameworksActive                    float64 `json:"master/frameworks_active"`
	Master_frameworksConnected                 float64 `json:"master/frameworks_connected"`
	Master_frameworksDisconnected              float64 `json:"master/frameworks_disconnected"`
	Master_frameworksInactive                  float64 `json:"master/frameworks_inactive"`
	Master_invalidFrameworkToExecutorMessages  float64 `json:"master/invalid_framework_to_executor_messages"`
	Master_invalidStatusUpdateAcknowledgements float64 `json:"master/invalid_status_update_acknowledgements"`
	Master_invalidStatusUpdates                float64 `json:"master/invalid_status_updates"`
	Master_memPercent                          float64 `json:"master/mem_percent"`
	Master_memRevocablePercent                 float64 `json:"master/mem_revocable_percent"`
	Master_memRevocableTotal                   float64 `json:"master/mem_revocable_total"`
	Master_memRevocableUsed                    float64 `json:"master/mem_revocable_used"`
	Master_memTotal                            float64 `json:"master/mem_total"`
	Master_memUsed                             float64 `json:"master/mem_used"`
	Master_messagesAuthenticate                float64 `json:"master/messages_authenticate"`
	Master_messagesDeactivateFramework         float64 `json:"master/messages_deactivate_framework"`
	Master_messagesDeclineOffers               float64 `json:"master/messages_decline_offers"`
	Master_messagesExitedExecutor              float64 `json:"master/messages_exited_executor"`
	Master_messagesFrameworkToExecutor         float64 `json:"master/messages_framework_to_executor"`
	Master_messagesKillTask                    float64 `json:"master/messages_kill_task"`
	Master_messagesLaunchTasks                 float64 `json:"master/messages_launch_tasks"`
	Master_messagesReconcileTasks              float64 `json:"master/messages_reconcile_tasks"`
	Master_messagesRegisterFramework           float64 `json:"master/messages_register_framework"`
	Master_messagesRegisterSlave               float64 `json:"master/messages_register_slave"`
	Master_messagesReregisterFramework         float64 `json:"master/messages_reregister_framework"`
	Master_messagesReregisterSlave             float64 `json:"master/messages_reregister_slave"`
	Master_messagesResourceRequest             float64 `json:"master/messages_resource_request"`
	Master_messagesReviveOffers                float64 `json:"master/messages_revive_offers"`
	Master_messagesStatusUpdate                float64 `json:"master/messages_status_update"`
	Master_messagesStatusUpdateAcknowledgement float64 `json:"master/messages_status_update_acknowledgement"`
	Master_messagesUnregisterFramework         float64 `json:"master/messages_unregister_framework"`
	Master_messagesUnregisterSlave             float64 `json:"master/messages_unregister_slave"`
	Master_messagesUpdateSlave                 float64 `json:"master/messages_update_slave"`
	Master_outstandingOffers                   float64 `json:"master/outstanding_offers"`
	Master_recoverySlaveRemovals               float64 `json:"master/recovery_slave_removals"`
	Master_slaveRegistrations                  float64 `json:"master/slave_registrations"`
	Master_slaveRemovals                       float64 `json:"master/slave_removals"`
	Master_slaveRemovals_reasonRegistered      float64 `json:"master/slave_removals/reason_registered"`
	Master_slaveRemovals_reasonUnhealthy       float64 `json:"master/slave_removals/reason_unhealthy"`
	Master_slaveRemovals_reasonUnregistered    float64 `json:"master/slave_removals/reason_unregistered"`
	Master_slaveReregistrations                float64 `json:"master/slave_reregistrations"`
	Master_slaveShutdownsCanceled              float64 `json:"master/slave_shutdowns_canceled"`
	Master_slaveShutdownsCompleted             float64 `json:"master/slave_shutdowns_completed"`
	Master_slaveShutdownsScheduled             float64 `json:"master/slave_shutdowns_scheduled"`
	Master_slavesActive                        float64 `json:"master/slaves_active"`
	Master_slavesConnected                     float64 `json:"master/slaves_connected"`
	Master_slavesDisconnected                  float64 `json:"master/slaves_disconnected"`
	Master_slavesInactive                      float64 `json:"master/slaves_inactive"`
	Master_tasksError                          float64 `json:"master/tasks_error"`
	Master_tasksFailed                         float64 `json:"master/tasks_failed"`
	Master_tasksFinished                       float64 `json:"master/tasks_finished"`
	Master_tasksKilled                         float64 `json:"master/tasks_killed"`
	Master_tasksLost                           float64 `json:"master/tasks_lost"`
	Master_tasksRunning                        float64 `json:"master/tasks_running"`
	Master_tasksStaging                        float64 `json:"master/tasks_staging"`
	Master_tasksStarting                       float64 `json:"master/tasks_starting"`
	Master_uptimeSecs                          float64 `json:"master/uptime_secs"`
	Master_validFrameworkToExecutorMessages    float64 `json:"master/valid_framework_to_executor_messages"`
	Master_validStatusUpdateAcknowledgements   float64 `json:"master/valid_status_update_acknowledgements"`
	Master_validStatusUpdates                  float64 `json:"master/valid_status_updates"`
	Registrar_queuedOperations                 float64 `json:"registrar/queued_operations"`
	Registrar_registrySizeBytes                float64 `json:"registrar/registry_size_bytes"`
	Registrar_stateFetchMs                     float64 `json:"registrar/state_fetch_ms"`
	Registrar_stateStoreMs                     float64 `json:"registrar/state_store_ms"`
	Registrar_stateStoreMs_count               float64 `json:"registrar/state_store_ms/count"`
	Registrar_stateStoreMs_max                 float64 `json:"registrar/state_store_ms/max"`
	Registrar_stateStoreMs_min                 float64 `json:"registrar/state_store_ms/min"`
	Registrar_stateStoreMs_p50                 float64 `json:"registrar/state_store_ms/p50"`
	Registrar_stateStoreMs_p90                 float64 `json:"registrar/state_store_ms/p90"`
	Registrar_stateStoreMs_p95                 float64 `json:"registrar/state_store_ms/p95"`
	Registrar_stateStoreMs_p99                 float64 `json:"registrar/state_store_ms/p99"`
	Registrar_stateStoreMs_p999                float64 `json:"registrar/state_store_ms/p999"`
	Registrar_stateStoreMs_p9999               float64 `json:"registrar/state_store_ms/p9999"`
	System_cpusTotal                           float64 `json:"system/cpus_total"`
	System_load15min                           float64 `json:"system/load_15min"`
	System_load1min                            float64 `json:"system/load_1min"`
	System_load5min                            float64 `json:"system/load_5min"`
	System_memFreeBytes                        float64 `json:"system/mem_free_bytes"`
	System_memTotalBytes                       float64 `json:"system/mem_total_bytes"`
	Time                                       time.Time
}
