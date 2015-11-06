package marathon

import (
	"time"
)

type histogram struct {
	Count  int     `json:"count"`
	Max    float64 `json:"max"`
	Mean   float64 `json:"mean"`
	Min    float64 `json:"min"`
	P50    float64 `json:"p50"`
	P75    float64 `json:"p75"`
	P95    float64 `json:"p95"`
	P98    float64 `json:"p98"`
	P99    float64 `json:"p99"`
	P999   float64 `json:"p999"`
	Stddev float64 `json:"stddev"`
}

type meter struct {
	Count    int     `json:"count"`
	M15Rate  float64 `json:"m15_rate"`
	M1Rate   float64 `json:"m1_rate"`
	M5Rate   float64 `json:"m5_rate"`
	MeanRate float64 `json:"mean_rate"`
}

type timer struct {
	Count    int     `json:"count"`
	M15Rate  float64 `json:"m15_rate"`
	M1Rate   float64 `json:"m1_rate"`
	M5Rate   float64 `json:"m5_rate"`
	Max      float64 `json:"max"`
	Mean     float64 `json:"mean"`
	MeanRate float64 `json:"mean_rate"`
	Min      float64 `json:"min"`
	P50      float64 `json:"p50"`
	P75      float64 `json:"p75"`
	P95      float64 `json:"p95"`
	P98      float64 `json:"p98"`
	P99      float64 `json:"p99"`
	P999     float64 `json:"p999"`
	Stddev   float64 `json:"stddev"`
}

type counters struct {
	Org_eclipse_jetty_servlet_ServletContextHandler_active_dispatches struct {
		Count int `json:"count"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.active-dispatches"`
	Org_eclipse_jetty_servlet_ServletContextHandler_active_requests struct {
		Count int `json:"count"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.active-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_active_suspended_requests struct {
		Count int `json:"count"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.active-suspended-requests"`
}

type gauges struct {
	Org_eclipse_jetty_servlet_ServletContextHandler_percent_4xx_15m struct {
		Value float64 `json:"value"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.percent-4xx-15m"`
	Org_eclipse_jetty_servlet_ServletContextHandler_percent_4xx_1m struct {
		Value float64 `json:"value"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.percent-4xx-1m"`
	Org_eclipse_jetty_servlet_ServletContextHandler_percent_4xx_5m struct {
		Value float64 `json:"value"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.percent-4xx-5m"`
	Org_eclipse_jetty_servlet_ServletContextHandler_percent_5xx_15m struct {
		Value float64 `json:"value"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.percent-5xx-15m"`
	Org_eclipse_jetty_servlet_ServletContextHandler_percent_5xx_1m struct {
		Value float64 `json:"value"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.percent-5xx-1m"`
	Org_eclipse_jetty_servlet_ServletContextHandler_percent_5xx_5m struct {
		Value float64 `json:"value"`
	} `json:"org.eclipse.jetty.servlet.ServletContextHandler.percent-5xx-5m"`
}

type histograms struct {
	Mesosphere_marathon_state_AppRepository_read_request_time              histogram `json:"mesosphere.marathon.state.AppRepository.read-request-time"`
	Mesosphere_marathon_state_AppRepository_write_request_time             histogram `json:"mesosphere.marathon.state.AppRepository.write-request-time"`
	Mesosphere_marathon_state_DeploymentRepository_read_request_time       histogram `json:"mesosphere.marathon.state.DeploymentRepository.read-request-time"`
	Mesosphere_marathon_state_DeploymentRepository_write_request_time      histogram `json:"mesosphere.marathon.state.DeploymentRepository.write-request-time"`
	Mesosphere_marathon_state_GroupRepository_read_request_time            histogram `json:"mesosphere.marathon.state.GroupRepository.read-request-time"`
	Mesosphere_marathon_state_GroupRepository_write_request_time           histogram `json:"mesosphere.marathon.state.GroupRepository.write-request-time"`
	Mesosphere_marathon_state_MarathonStore_AppDefinition_read_data_size   histogram `json:"mesosphere.marathon.state.MarathonStore.AppDefinition.read-data-size"`
	Mesosphere_marathon_state_MarathonStore_AppDefinition_write_data_size  histogram `json:"mesosphere.marathon.state.MarathonStore.AppDefinition.write-data-size"`
	Mesosphere_marathon_state_MarathonStore_DeploymentPlan_read_data_size  histogram `json:"mesosphere.marathon.state.MarathonStore.DeploymentPlan.read-data-size"`
	Mesosphere_marathon_state_MarathonStore_DeploymentPlan_write_data_size histogram `json:"mesosphere.marathon.state.MarathonStore.DeploymentPlan.write-data-size"`
	Mesosphere_marathon_state_MarathonStore_Group_read_data_size           histogram `json:"mesosphere.marathon.state.MarathonStore.Group.read-data-size"`
	Mesosphere_marathon_state_MarathonStore_Group_write_data_size          histogram `json:"mesosphere.marathon.state.MarathonStore.Group.write-data-size"`
	Mesosphere_marathon_state_MarathonStore_TaskFailure_read_data_size     histogram `json:"mesosphere.marathon.state.MarathonStore.TaskFailure.read-data-size"`
	Mesosphere_marathon_state_MarathonStore_TaskFailure_write_data_size    histogram `json:"mesosphere.marathon.state.MarathonStore.TaskFailure.write-data-size"`
	Mesosphere_marathon_tasks_TaskTracker_read_request_time                histogram `json:"mesosphere.marathon.tasks.TaskTracker.read-request-time"`
	Mesosphere_marathon_tasks_TaskTracker_write_request_time               histogram `json:"mesosphere.marathon.tasks.TaskTracker.write-request-time"`
}

type meters struct {
	Mesosphere_marathon_state_AppRepository_read_request_errors         meter `json:"mesosphere.marathon.state.AppRepository.read-request-errors"`
	Mesosphere_marathon_state_AppRepository_read_requests               meter `json:"mesosphere.marathon.state.AppRepository.read-requests"`
	Mesosphere_marathon_state_AppRepository_write_request_errors        meter `json:"mesosphere.marathon.state.AppRepository.write-request-errors"`
	Mesosphere_marathon_state_AppRepository_write_requests              meter `json:"mesosphere.marathon.state.AppRepository.write-requests"`
	Mesosphere_marathon_state_DeploymentRepository_read_request_errors  meter `json:"mesosphere.marathon.state.DeploymentRepository.read-request-errors"`
	Mesosphere_marathon_state_DeploymentRepository_read_requests        meter `json:"mesosphere.marathon.state.DeploymentRepository.read-requests"`
	Mesosphere_marathon_state_DeploymentRepository_write_request_errors meter `json:"mesosphere.marathon.state.DeploymentRepository.write-request-errors"`
	Mesosphere_marathon_state_DeploymentRepository_write_requests       meter `json:"mesosphere.marathon.state.DeploymentRepository.write-requests"`
	Mesosphere_marathon_state_GroupRepository_read_request_errors       meter `json:"mesosphere.marathon.state.GroupRepository.read-request-errors"`
	Mesosphere_marathon_state_GroupRepository_read_requests             meter `json:"mesosphere.marathon.state.GroupRepository.read-requests"`
	Mesosphere_marathon_state_GroupRepository_write_request_errors      meter `json:"mesosphere.marathon.state.GroupRepository.write-request-errors"`
	Mesosphere_marathon_state_GroupRepository_write_requests            meter `json:"mesosphere.marathon.state.GroupRepository.write-requests"`
	Mesosphere_marathon_tasks_TaskTracker_read_request_errors           meter `json:"mesosphere.marathon.tasks.TaskTracker.read-request-errors"`
	Mesosphere_marathon_tasks_TaskTracker_read_requests                 meter `json:"mesosphere.marathon.tasks.TaskTracker.read-requests"`
	Mesosphere_marathon_tasks_TaskTracker_write_request_errors          meter `json:"mesosphere.marathon.tasks.TaskTracker.write-request-errors"`
	Mesosphere_marathon_tasks_TaskTracker_write_requests                meter `json:"mesosphere.marathon.tasks.TaskTracker.write-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_1xx_responses       meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.1xx-responses"`
	Org_eclipse_jetty_servlet_ServletContextHandler_2xx_responses       meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.2xx-responses"`
	Org_eclipse_jetty_servlet_ServletContextHandler_3xx_responses       meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.3xx-responses"`
	Org_eclipse_jetty_servlet_ServletContextHandler_4xx_responses       meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.4xx-responses"`
	Org_eclipse_jetty_servlet_ServletContextHandler_5xx_responses       meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.5xx-responses"`
	Org_eclipse_jetty_servlet_ServletContextHandler_expires             meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.expires"`
	Org_eclipse_jetty_servlet_ServletContextHandler_requests            meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_resumes             meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.resumes"`
	Org_eclipse_jetty_servlet_ServletContextHandler_suspends            meter `json:"org.eclipse.jetty.servlet.ServletContextHandler.suspends"`
}

type timers struct {
	Mesosphere_marathon_api_v2_AppVersionsResource_index                  timer `json:"mesosphere.marathon.api.v2.AppVersionsResource.index"`
	Mesosphere_marathon_api_v2_AppVersionsResource_show                   timer `json:"mesosphere.marathon.api.v2.AppVersionsResource.show"`
	Mesosphere_marathon_api_v2_AppsResource_create                        timer `json:"mesosphere.marathon.api.v2.AppsResource.create"`
	Mesosphere_marathon_api_v2_AppsResource_delete                        timer `json:"mesosphere.marathon.api.v2.AppsResource.delete"`
	Mesosphere_marathon_api_v2_AppsResource_index                         timer `json:"mesosphere.marathon.api.v2.AppsResource.index"`
	Mesosphere_marathon_api_v2_AppsResource_replace                       timer `json:"mesosphere.marathon.api.v2.AppsResource.replace"`
	Mesosphere_marathon_api_v2_AppsResource_replaceMultiple               timer `json:"mesosphere.marathon.api.v2.AppsResource.replaceMultiple"`
	Mesosphere_marathon_api_v2_AppsResource_show                          timer `json:"mesosphere.marathon.api.v2.AppsResource.show"`
	Mesosphere_marathon_api_v2_EventSubscriptionsResource_listSubscribers timer `json:"mesosphere.marathon.api.v2.EventSubscriptionsResource.listSubscribers"`
	Mesosphere_marathon_api_v2_EventSubscriptionsResource_subscribe       timer `json:"mesosphere.marathon.api.v2.EventSubscriptionsResource.subscribe"`
	Mesosphere_marathon_api_v2_EventSubscriptionsResource_unsubscribe     timer `json:"mesosphere.marathon.api.v2.EventSubscriptionsResource.unsubscribe"`
	Mesosphere_marathon_api_v2_GroupsResource_create                      timer `json:"mesosphere.marathon.api.v2.GroupsResource.create"`
	Mesosphere_marathon_api_v2_GroupsResource_createWithPath              timer `json:"mesosphere.marathon.api.v2.GroupsResource.createWithPath"`
	Mesosphere_marathon_api_v2_GroupsResource_delete                      timer `json:"mesosphere.marathon.api.v2.GroupsResource.delete"`
	Mesosphere_marathon_api_v2_GroupsResource_group                       timer `json:"mesosphere.marathon.api.v2.GroupsResource.group"`
	Mesosphere_marathon_api_v2_GroupsResource_root                        timer `json:"mesosphere.marathon.api.v2.GroupsResource.root"`
	Mesosphere_marathon_api_v2_GroupsResource_update                      timer `json:"mesosphere.marathon.api.v2.GroupsResource.update"`
	Mesosphere_marathon_api_v2_GroupsResource_updateRoot                  timer `json:"mesosphere.marathon.api.v2.GroupsResource.updateRoot"`
	Mesosphere_marathon_api_v2_QueueResource_index                        timer `json:"mesosphere.marathon.api.v2.QueueResource.index"`
	Mesosphere_marathon_api_v2_TasksResource_indexJSON                    timer `json:"mesosphere.marathon.api.v2.TasksResource.indexJson"`
	Mesosphere_marathon_api_v2_TasksResource_indexTxt                     timer `json:"mesosphere.marathon.api.v2.TasksResource.indexTxt"`
	Mesosphere_marathon_api_v2_TasksResource_killTasks                    timer `json:"mesosphere.marathon.api.v2.TasksResource.killTasks"`
	Org_eclipse_jetty_servlet_ServletContextHandler_connect_requests      timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.connect-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_delete_requests       timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.delete-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_dispatches            timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.dispatches"`
	Org_eclipse_jetty_servlet_ServletContextHandler_get_requests          timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.get-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_head_requests         timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.head-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_options_requests      timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.options-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_other_requests        timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.other-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_patch_requests        timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.patch-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_post_requests         timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.post-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_put_requests          timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.put-requests"`
	Org_eclipse_jetty_servlet_ServletContextHandler_trace_requests        timer `json:"org.eclipse.jetty.servlet.ServletContextHandler.trace-requests"`
}

type MarathonStats struct {
	Counters   counters   `json:"counters"`
	Gauges     gauges     `json:"gauges"`
	Histograms histograms `json:"histograms"`
	Meters     meters     `json:"meters"`
	Timers     timers     `json:"timers"`
	Version    string     `json:"version"`
	Time       time.Time
	Node       string
}
