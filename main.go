package main

import (
	"embed"
	"fmt"
	"graylog-alert-exporter/internal/config"
	"graylog-alert-exporter/internal/log"
	"graylog-alert-exporter/pkg/database"
	"graylog-alert-exporter/pkg/handlers"
	"graylog-alert-exporter/pkg/scheduler"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"

	//go:embed index.html
	resources embed.FS
)

func init() {
	viper.SetDefault("versionNumber", version)
	viper.SetDefault("commit", commit)
	viper.SetDefault("date", date)
	viper.SetDefault("builtBy", builtBy)
}

func main() {
	config.Init()
	if viper.GetBool("version") {
		fmt.Println(viper.GetString("versionNumber"))
		os.Exit(0)
	}

	log.Init()
	database.Init()

	app := fiber.New(fiber.Config{Views: html.NewFileSystem(http.FS(resources), ".html")})
	app.Use(logger.New())
	app.Use(etag.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"link": viper.GetString("path"), "dashboard": viper.GetBool("dashboard")})
	})

	app.Get(viper.GetString("path"), handlers.PrometheusHandler)
	app.Post(viper.GetString("path"), handlers.GetGraylogOutputHandler)

	scheduler.StartTimeoutScheduler(viper.GetInt("interval"))

	logrus.Info("Starting graylog alert exporter")
	logrus.Panic(app.Listen(viper.GetString("listen")))
}
