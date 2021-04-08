// Package handlers provide web service for router
package handlers

import (
	"graylog-alert-exporter/pkg/database"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

const (
	Firing   = 1
	Resolved = 0
)

// PrometheusHandler is handler to control prometheus metrics
func PrometheusHandler(ctx *fiber.Ctx) error {
	Registry := prometheus.NewRegistry()

	// Default metrics
	BuildInfoMetrics := prometheus.NewBuildInfoCollector()
	err := Registry.Register(BuildInfoMetrics)
	if err != nil {
		logrus.Error(err)
	}

	// Alert metrics
	alerts := database.GetAllAlerts()
	for _, alert := range alerts {
		m := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name:        "graylog_alert_status",
				Help:        "Graylog alert status [0 is Resolved, 1 is Firing]",
				ConstLabels: alert.Data,
			},
		)
		m.Set(Firing)
		if alert.Timeout <= 0 {
			m.Set(Resolved)
		}
		err := Registry.Register(m)
		if err != nil {
			logrus.Error(err)
		}
	}

	return adaptor.HTTPHandler(promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))(ctx)
}
