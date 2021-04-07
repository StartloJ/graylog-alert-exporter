// Package prom use to expose metrics for prometheus
package prom

import (
	"graylog-alert-exporter/pkg/meta"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(buildInfoMetrics)
	buildInfoMetrics.Set(1)
	promHandler = promhttp.HandlerFor(Registry, promhttp.HandlerOpts{})
}

// Registry global var for prometheus registry
var Registry *prometheus.Registry
var promHandler http.Handler
var (
	buildInfoMetrics = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "graylog_alert_exporter_build_info",
			Help: "Exporter build information",
			ConstLabels: prometheus.Labels{
				"version": meta.Version,
				"commit":  meta.Commit,
				"date":    meta.Date,
			},
		},
	)
)

func PrometheusHandler() func(ctx *fiber.Ctx) error {
	return adaptor.HTTPHandler(promHandler)
}
