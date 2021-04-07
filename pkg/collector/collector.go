package collector

import (
	"time"
)

type AlertMetrics struct {
	EventID          string
	EventTitle       string
	EventDescription string
	EventTimeStamp   *time.Time
	EventSource      string
	EventPriority    int
}

// var (
// 	GraylogAlertInfo = prometheus.NewGauge(
// 		prometheus.GaugeOpts{
// 			Name: "graylog_alert_events",
// 			Help: "Keep information about alert event-base if exporter's trickered",
// 			ConstLabels: prometheus.Labels{
// 				"event_id":        AlertMetrics.EventID,
// 				"event_title":     AlertMetrics.EventTitle,
// 				"event_desc":      AlertMetrics.EventDescription,
// 				"event_timestamp": AlertMetrics.EventTimeStamp.String(),
// 				"event_source":    AlertMetrics.EventSource,
// 				"event_priority":  AlertMetrics.EventPriority,
// 			},
// 		},
// 	)
// )
