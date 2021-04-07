// Package handlers provide web service for router
package handlers

import (
	"graylog-alert-exporter/pkg/database"
	"strconv"

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
	BuildInfoMetrics := prometheus.NewBuildInfoCollector()
	Registry.Register(BuildInfoMetrics)

	alerts := database.GetAllAlerts()
	for _, alert := range alerts {
		m := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "main_metrics",
				Help: "Main metric",
				ConstLabels: map[string]string{
					"event_title":       alert.Title,
					"event_description": alert.Description,
					"event_source":      alert.Source,
					"event_priority":    strconv.Itoa(alert.Priority),
				},
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
