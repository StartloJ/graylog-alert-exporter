package main

import (
	"embed"
	"fmt"
	"graylog-alert-exporter/internal/config"
	"graylog-alert-exporter/internal/log"
	"graylog-alert-exporter/pkg/database"
	"graylog-alert-exporter/pkg/handlers"
	"graylog-alert-exporter/pkg/scheduler"
	"io/fs"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"

	//go:embed web/*
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

	app := fiber.New()
	app.Use(logger.New())

	webDir, _ := fs.Sub(resources, "web")
	if viper.GetBool("dashboard") {
		app.Use("/", filesystem.New(filesystem.Config{
			Root:   http.FS(webDir),
			Browse: true,
			Index:  "index.html",
		}))
	}

	app.Get(viper.GetString("path"), handlers.PrometheusHandler)
	app.Post(viper.GetString("path"), handlers.ReceiveGraylogAlertHandler)

	api := app.Group("/api")
	api.Get("/alerts", handlers.GetAlerts)
	api.Post("/alert", handlers.UpdateAlert)
	api.Delete("/alert/:id", handlers.DeleteAlert)

	scheduler.StartTimeoutScheduler(viper.GetInt("interval"))

	logrus.Info("Starting graylog alert exporter")
	logrus.Error(app.Listen(viper.GetString("listen")))
}
