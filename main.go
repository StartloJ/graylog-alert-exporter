package main

import (
	"flag"
	"fmt"
	"graylog-alert-exporter/pkg/client"
	"graylog-alert-exporter/pkg/meta"
	"graylog-alert-exporter/pkg/prom"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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
		if err != nil {
			return listener, fmt.Errorf("ListenUnix address %s fialed: %v", listenAddress, err)
		}
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
	// Default values
	defaultListenAddress = getEnv("GAE_LISTEN_ADDRESS", ":9889")
	defaultMetricsPath   = getEnv("GAE_TELEMETRY_PATH", "/metrics")

	// Command-line flags
	listenAddress = flag.String("web.listen-address", defaultListenAddress, "Address to listen on for telemetry")
	metricsPath   = flag.String("web.metrics-path", defaultMetricsPath, "Path under which to expose metrics")
)

func main() {
	flag.Parse()
	fmt.Printf("Graylog Alert Exporter version=%v commit=%v date=%v", meta.Version, meta.Commit, meta.Date)

	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Link": *metricsPath,
		})
	})

	app.Get(*metricsPath, prom.PrometheusHandler())
	app.Post("/store", client.UserServePayload)

	app.Listen(*listenAddress)
}
