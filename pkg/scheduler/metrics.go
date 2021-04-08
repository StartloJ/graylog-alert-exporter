// Package scheduler provide worker to run background task
package scheduler

import (
	"graylog-alert-exporter/pkg/database"
	"time"
)

// StartTimeoutScheduler enable scheduler reduce timeout in struct. 'interval' are set in second.
func StartTimeoutScheduler(interval int) {
	go func() {
		for {
			alerts := database.GetAllAlerts()
			for _, alert := range alerts {
				if alert.Timeout > 0 {
					alert.Timeout -= interval
					database.InsertAlert(alert)
				}
			}
			time.Sleep(time.Second * time.Duration(interval))
		}
	}()
}
