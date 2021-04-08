// Package config provide environment and global variables in program
package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Init is function to initialize config variables
func Init() {
	pflag.StringP("listen", "l", "0.0.0.0:9889", "Host address to start service listen")
	pflag.StringP("path", "p", "metrics", "path for scape and push metrics")
	pflag.IntP("timeout", "t", 60, "timeout of alert to make it resolved in second")
	pflag.IntP("interval", "i", 5, "interval to check timeout (lower value consume more cpu) in second")
	pflag.Bool("debug", false, "enable debug log")
	pflag.Bool("caller", false, "enable log method caller in code")
	pflag.StringP("label_file", "f", "labels.yaml", "Map labels config file to dynamic label in Prometheus metrics")
	pflag.BoolP("version", "v", false, "print version")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	logrus.Error(err)

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("exporter")
	err = viper.BindEnv("listen")
	logrus.Debug(err)
	err = viper.BindEnv("path")
	logrus.Debug(err)
	err = viper.BindEnv("timeout")
	logrus.Debug(err)
	err = viper.BindEnv("interval")
	logrus.Debug(err)
	err = viper.BindEnv("debug")
	logrus.Debug(err)
	err = viper.BindEnv("caller")
	logrus.Debug(err)
	err = viper.BindEnv("label_file")
	logrus.Debug(err)

	if viper.IsSet("label_file") {
		file, err := os.Open(viper.GetString("label_file"))
		if err != nil {
			logrus.Panic(err)
		}
		viper.SetConfigType("yaml")
		if err := viper.ReadConfig(file); err != nil {
			logrus.Panic(err)
		}
	}
}
