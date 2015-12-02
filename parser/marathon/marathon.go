package marathon

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/kpacha/mesos-influxdb-collector/store"
)

type MarathonStatsParser struct {
	Node string
}

func (mp MarathonStatsParser) Parse(r io.ReadCloser) ([]store.Point, error) {
	defer r.Close()
	var stats MarathonStats
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("Error reading from", r)
		return []store.Point{}, err
	}
	if err = json.Unmarshal(body, &stats); err != nil {
		log.Println("Error parsing to MarathonStats")
		return []store.Point{}, err
	}
	stats.Node = mp.Node
	stats.Time = time.Now()
	return mp.getMarathonPoints(stats), nil
}

func (mp MarathonStatsParser) getMarathonPoints(stats MarathonStats) []store.Point {
	return []store.Point{
		mp.getCounterPoint(stats),
		mp.getGaugePoint(stats),
		mp.getRequestTimePoint(stats),
		mp.getDataSizePoint(stats),
		mp.getRequestPoint(stats),
		mp.getRequestErrorPoint(stats),
		mp.getResponsePoint(stats),
	}
}

func (mp MarathonStatsParser) getCounterPoint(stats MarathonStats) store.Point {
	return store.Point{
		Measurement: "marathon-servlet-context",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"active-dispatchers": stats.Counters.Org_eclipse_jetty_servlet_ServletContextHandler_active_dispatches,
			"active-requests":    stats.Counters.Org_eclipse_jetty_servlet_ServletContextHandler_active_requests,
			"suspended-request":  stats.Counters.Org_eclipse_jetty_servlet_ServletContextHandler_active_suspended_requests,
			"expires":            stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_expires,
			"resumes":            stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_resumes,
			"suspends":           stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_suspends,
		},
		Time: stats.Time,
	}
}

func (mp MarathonStatsParser) getGaugePoint(stats MarathonStats) store.Point {
	return store.Point{
		Measurement: "marathon-error",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"4xx_15m": stats.Gauges.Org_eclipse_jetty_servlet_ServletContextHandler_percent_4xx_15m,
			"4xx_1m":  stats.Gauges.Org_eclipse_jetty_servlet_ServletContextHandler_percent_4xx_1m,
			"4xx_5m":  stats.Gauges.Org_eclipse_jetty_servlet_ServletContextHandler_percent_4xx_5m,
			"5xx_15m": stats.Gauges.Org_eclipse_jetty_servlet_ServletContextHandler_percent_5xx_15m,
			"5xx_1m":  stats.Gauges.Org_eclipse_jetty_servlet_ServletContextHandler_percent_5xx_1m,
			"5xx_5m":  stats.Gauges.Org_eclipse_jetty_servlet_ServletContextHandler_percent_5xx_5m,
		},
		Time: stats.Time,
	}
}

func (mp MarathonStatsParser) getRequestTimePoint(stats MarathonStats) store.Point {
	return store.Point{
		Measurement: "marathon-request-time",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"AppRepository-read":         stats.Histograms.Mesosphere_marathon_state_AppRepository_read_request_time,
			"AppRepository-write":        stats.Histograms.Mesosphere_marathon_state_AppRepository_write_request_time,
			"DeploymentRepository-read":  stats.Histograms.Mesosphere_marathon_state_DeploymentRepository_read_request_time,
			"DeploymentRepository-write": stats.Histograms.Mesosphere_marathon_state_DeploymentRepository_write_request_time,
			"GroupRepository_read":       stats.Histograms.Mesosphere_marathon_state_GroupRepository_read_request_time,
			"GroupRepository_write":      stats.Histograms.Mesosphere_marathon_state_GroupRepository_write_request_time,
			"TaskTracker-read":           stats.Histograms.Mesosphere_marathon_tasks_TaskTracker_read_request_time,
			"TaskTracker-write":          stats.Histograms.Mesosphere_marathon_tasks_TaskTracker_write_request_time,
		},
		Time: stats.Time,
	}
}

func (mp MarathonStatsParser) getDataSizePoint(stats MarathonStats) store.Point {
	return store.Point{
		Measurement: "marathon-data-size",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"AppDefinition-read":   stats.Histograms.Mesosphere_marathon_state_MarathonStore_AppDefinition_read_data_size,
			"AppDefinition-write":  stats.Histograms.Mesosphere_marathon_state_MarathonStore_AppDefinition_write_data_size,
			"DeploymentPlan-read":  stats.Histograms.Mesosphere_marathon_state_MarathonStore_DeploymentPlan_read_data_size,
			"DeploymentPlan-write": stats.Histograms.Mesosphere_marathon_state_MarathonStore_DeploymentPlan_write_data_size,
			"Group-read":           stats.Histograms.Mesosphere_marathon_state_MarathonStore_Group_read_data_size,
			"Group-write":          stats.Histograms.Mesosphere_marathon_state_MarathonStore_Group_write_data_size,
			"TeskFailure-read":     stats.Histograms.Mesosphere_marathon_state_MarathonStore_TaskFailure_read_data_size,
			"TeskFailure-write":    stats.Histograms.Mesosphere_marathon_state_MarathonStore_TaskFailure_write_data_size,
		},
		Time: stats.Time,
	}
}

func (mp MarathonStatsParser) getRequestPoint(stats MarathonStats) store.Point {
	return store.Point{
		Measurement: "marathon-request",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"AppRepository-read":         stats.Meters.Mesosphere_marathon_state_AppRepository_read_requests,
			"AppRepository-write":        stats.Meters.Mesosphere_marathon_state_AppRepository_write_requests,
			"DeploymentRepository-read":  stats.Meters.Mesosphere_marathon_state_DeploymentRepository_read_requests,
			"DeploymentRepository-write": stats.Meters.Mesosphere_marathon_state_DeploymentRepository_write_requests,
			"GroupRepository_read":       stats.Meters.Mesosphere_marathon_state_GroupRepository_read_requests,
			"GroupRepository_write":      stats.Meters.Mesosphere_marathon_state_GroupRepository_write_requests,
			"TaskTracker-read":           stats.Meters.Mesosphere_marathon_tasks_TaskTracker_read_requests,
			"TaskTracker-write":          stats.Meters.Mesosphere_marathon_tasks_TaskTracker_write_requests,
			"total":                      stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_requests,
		},
		Time: stats.Time,
	}
}

func (mp MarathonStatsParser) getRequestErrorPoint(stats MarathonStats) store.Point {
	return store.Point{
		Measurement: "marathon-request-error",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"AppRepository-read":         stats.Meters.Mesosphere_marathon_state_AppRepository_read_request_errors,
			"AppRepository-write":        stats.Meters.Mesosphere_marathon_state_AppRepository_write_request_errors,
			"DeploymentRepository-read":  stats.Meters.Mesosphere_marathon_state_DeploymentRepository_read_request_errors,
			"DeploymentRepository-write": stats.Meters.Mesosphere_marathon_state_DeploymentRepository_write_request_errors,
			"GroupRepository_read":       stats.Meters.Mesosphere_marathon_state_GroupRepository_read_request_errors,
			"GroupRepository_write":      stats.Meters.Mesosphere_marathon_state_GroupRepository_write_request_errors,
			"TaskTracker-read":           stats.Meters.Mesosphere_marathon_tasks_TaskTracker_read_request_errors,
			"TaskTracker-write":          stats.Meters.Mesosphere_marathon_tasks_TaskTracker_write_request_errors,
		},
		Time: stats.Time,
	}
}

func (mp MarathonStatsParser) getResponsePoint(stats MarathonStats) store.Point {
	return store.Point{
		Measurement: "marathon-request-error",
		Tags: map[string]string{
			"node": stats.Node,
		},
		Fields: map[string]interface{}{
			"1xx": stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_1xx_responses,
			"2xx": stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_2xx_responses,
			"3xx": stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_3xx_responses,
			"4xx": stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_4xx_responses,
			"5xx": stats.Meters.Org_eclipse_jetty_servlet_ServletContextHandler_5xx_responses,
		},
		Time: stats.Time,
	}
}
