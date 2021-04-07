package main

import (
	"flag"
	"fmt"
	"graylog-alert-exporter/pkg/client"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func parseUnixSocketAddr(addr string) (string, string, error) {
	addrPart := strings.Split(addr, ":")
	addrPartLen := len(addrPart)

	if addrPartLen > 3 || addrPartLen < 1 {
		return "", "", fmt.Errorf("address for unix domain socket has wrong format")
	}

	unixSocketPath := addrPart[1]
	requestPath := ""
	if addrPartLen == 3 {
		requestPath = addrPart[2]
	}
	return unixSocketPath, requestPath, nil
}

func getListener(listenAddress string) (net.Listener, error) {
	var listener net.Listener
	var err error

	if strings.HasPrefix(listenAddress, "unix:") {
		path, _, err := parseUnixSocketAddr(listenAddress)
		if err != nil {
			return listener, fmt.Errorf("parse unix socket listen address %s fialed: %v", listenAddress, err)
		}
		listener, err = net.ListenUnix("unix", &net.UnixAddr{Name: path, Net: "unix"})
	} else {
		listener, err = net.Listen("tcp", listenAddress)
	}
	if err != nil {
		return listener, err
	}
	log.Printf("Listening on %s", listenAddress)
	return listener, nil
}

var (
	// Set during go build
	version string
	commit  string
	date    string

	// Default values
	defaultListenAddress = getEnv("GAE_LISTEN_ADDRESS", ":9889")
	defaultMetricsPath   = getEnv("GAE_TELEMETRY_PATH", "/metrics")

	// Command-line flags
	listenAddress = flag.String("web.listen-address", defaultListenAddress, "Address to listen on for telemetry")
	metricsPath   = flag.String("web.metrics-path", defaultMetricsPath, "Path under which to expose metrics")
)

var (
	buildInfoMetrics = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "graylog_alert_exporter_build_info",
			Help: "Exporter build information",
			ConstLabels: prometheus.Labels{
				"version": version,
				"commit":  commit,
				"date":    date,
			},
		},
	)
)

func main() {
	flag.Parse()

	fmt.Printf("Graylog Alert Exporter version=%v commit=%v date=%v", version, commit, date)

	r := prometheus.NewRegistry()
	r.MustRegister(buildInfoMetrics)

	buildInfoMetrics.Set(1)

	engine := html.New("./html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Link": *metricsPath,
		})
	})
	p := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	app.Get(*metricsPath, adaptor.HTTPHandler(p))
	app.Post("/store", client.UserServePayload)

	app.Listen(*listenAddress)
}
