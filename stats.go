package main

import (
	"time"
)

type Stats struct {
	Allocator_eventQueueDispatches             int     `json:"allocator/event_queue_dispatches"`
	Master_cpusPercent                         int     `json:"master/cpus_percent"`
	Master_cpusRevocablePercent                int     `json:"master/cpus_revocable_percent"`
	Master_cpusRevocableTotal                  int     `json:"master/cpus_revocable_total"`
	Master_cpusRevocableUsed                   int     `json:"master/cpus_revocable_used"`
	Master_cpusTotal                           int     `json:"master/cpus_total"`
	Master_cpusUsed                            int     `json:"master/cpus_used"`
	Master_diskPercent                         int     `json:"master/disk_percent"`
	Master_diskRevocablePercent                int     `json:"master/disk_revocable_percent"`
	Master_diskRevocableTotal                  int     `json:"master/disk_revocable_total"`
	Master_diskRevocableUsed                   int     `json:"master/disk_revocable_used"`
	Master_diskTotal                           int     `json:"master/disk_total"`
	Master_diskUsed                            int     `json:"master/disk_used"`
	Master_droppedMessages                     int     `json:"master/dropped_messages"`
	Master_elected                             int     `json:"master/elected"`
	Master_eventQueueDispatches                int     `json:"master/event_queue_dispatches"`
	Master_eventQueueHTTPRequests              int     `json:"master/event_queue_http_requests"`
	Master_eventQueueMessages                  int     `json:"master/event_queue_messages"`
	Master_frameworksActive                    int     `json:"master/frameworks_active"`
	Master_frameworksConnected                 int     `json:"master/frameworks_connected"`
	Master_frameworksDisconnected              int     `json:"master/frameworks_disconnected"`
	Master_frameworksInactive                  int     `json:"master/frameworks_inactive"`
	Master_invalidFrameworkToExecutorMessages  int     `json:"master/invalid_framework_to_executor_messages"`
	Master_invalidStatusUpdateAcknowledgements int     `json:"master/invalid_status_update_acknowledgements"`
	Master_invalidStatusUpdates                int     `json:"master/invalid_status_updates"`
	Master_memPercent                          int     `json:"master/mem_percent"`
	Master_memRevocablePercent                 int     `json:"master/mem_revocable_percent"`
	Master_memRevocableTotal                   int     `json:"master/mem_revocable_total"`
	Master_memRevocableUsed                    int     `json:"master/mem_revocable_used"`
	Master_memTotal                            int     `json:"master/mem_total"`
	Master_memUsed                             int     `json:"master/mem_used"`
	Master_messagesAuthenticate                int     `json:"master/messages_authenticate"`
	Master_messagesDeactivateFramework         int     `json:"master/messages_deactivate_framework"`
	Master_messagesDeclineOffers               int     `json:"master/messages_decline_offers"`
	Master_messagesExitedExecutor              int     `json:"master/messages_exited_executor"`
	Master_messagesFrameworkToExecutor         int     `json:"master/messages_framework_to_executor"`
	Master_messagesKillTask                    int     `json:"master/messages_kill_task"`
	Master_messagesLaunchTasks                 int     `json:"master/messages_launch_tasks"`
	Master_messagesReconcileTasks              int     `json:"master/messages_reconcile_tasks"`
	Master_messagesRegisterFramework           int     `json:"master/messages_register_framework"`
	Master_messagesRegisterSlave               int     `json:"master/messages_register_slave"`
	Master_messagesReregisterFramework         int     `json:"master/messages_reregister_framework"`
	Master_messagesReregisterSlave             int     `json:"master/messages_reregister_slave"`
	Master_messagesResourceRequest             int     `json:"master/messages_resource_request"`
	Master_messagesReviveOffers                int     `json:"master/messages_revive_offers"`
	Master_messagesStatusUpdate                int     `json:"master/messages_status_update"`
	Master_messagesStatusUpdateAcknowledgement int     `json:"master/messages_status_update_acknowledgement"`
	Master_messagesUnregisterFramework         int     `json:"master/messages_unregister_framework"`
	Master_messagesUnregisterSlave             int     `json:"master/messages_unregister_slave"`
	Master_messagesUpdateSlave                 int     `json:"master/messages_update_slave"`
	Master_outstandingOffers                   int     `json:"master/outstanding_offers"`
	Master_recoverySlaveRemovals               int     `json:"master/recovery_slave_removals"`
	Master_slaveRegistrations                  int     `json:"master/slave_registrations"`
	Master_slaveRemovals                       int     `json:"master/slave_removals"`
	Master_slaveRemovals_reasonRegistered      int     `json:"master/slave_removals/reason_registered"`
	Master_slaveRemovals_reasonUnhealthy       int     `json:"master/slave_removals/reason_unhealthy"`
	Master_slaveRemovals_reasonUnregistered    int     `json:"master/slave_removals/reason_unregistered"`
	Master_slaveReregistrations                int     `json:"master/slave_reregistrations"`
	Master_slaveShutdownsCanceled              int     `json:"master/slave_shutdowns_canceled"`
	Master_slaveShutdownsCompleted             int     `json:"master/slave_shutdowns_completed"`
	Master_slaveShutdownsScheduled             int     `json:"master/slave_shutdowns_scheduled"`
	Master_slavesActive                        int     `json:"master/slaves_active"`
	Master_slavesConnected                     int     `json:"master/slaves_connected"`
	Master_slavesDisconnected                  int     `json:"master/slaves_disconnected"`
	Master_slavesInactive                      int     `json:"master/slaves_inactive"`
	Master_tasksError                          int     `json:"master/tasks_error"`
	Master_tasksFailed                         int     `json:"master/tasks_failed"`
	Master_tasksFinished                       int     `json:"master/tasks_finished"`
	Master_tasksKilled                         int     `json:"master/tasks_killed"`
	Master_tasksLost                           int     `json:"master/tasks_lost"`
	Master_tasksRunning                        int     `json:"master/tasks_running"`
	Master_tasksStaging                        int     `json:"master/tasks_staging"`
	Master_tasksStarting                       int     `json:"master/tasks_starting"`
	Master_uptimeSecs                          float64 `json:"master/uptime_secs"`
	Master_validFrameworkToExecutorMessages    int     `json:"master/valid_framework_to_executor_messages"`
	Master_validStatusUpdateAcknowledgements   int     `json:"master/valid_status_update_acknowledgements"`
	Master_validStatusUpdates                  int     `json:"master/valid_status_updates"`
	Registrar_queuedOperations                 int     `json:"registrar/queued_operations"`
	Registrar_registrySizeBytes                int     `json:"registrar/registry_size_bytes"`
	Registrar_stateFetchMs                     float64 `json:"registrar/state_fetch_ms"`
	Registrar_stateStoreMs                     float64 `json:"registrar/state_store_ms"`
	Registrar_stateStoreMs_count               int     `json:"registrar/state_store_ms/count"`
	Registrar_stateStoreMs_max                 float64 `json:"registrar/state_store_ms/max"`
	Registrar_stateStoreMs_min                 float64 `json:"registrar/state_store_ms/min"`
	Registrar_stateStoreMs_p50                 float64 `json:"registrar/state_store_ms/p50"`
	Registrar_stateStoreMs_p90                 float64 `json:"registrar/state_store_ms/p90"`
	Registrar_stateStoreMs_p95                 float64 `json:"registrar/state_store_ms/p95"`
	Registrar_stateStoreMs_p99                 float64 `json:"registrar/state_store_ms/p99"`
	Registrar_stateStoreMs_p999                float64 `json:"registrar/state_store_ms/p999"`
	Registrar_stateStoreMs_p9999               float64 `json:"registrar/state_store_ms/p9999"`
	System_cpusTotal                           int     `json:"system/cpus_total"`
	System_load15min                           float64 `json:"system/load_15min"`
	System_load1min                            float64 `json:"system/load_1min"`
	System_load5min                            float64 `json:"system/load_5min"`
	System_memFreeBytes                        int     `json:"system/mem_free_bytes"`
	System_memTotalBytes                       int     `json:"system/mem_total_bytes"`
	Time                                       time.Time
	Node                                       string
}
