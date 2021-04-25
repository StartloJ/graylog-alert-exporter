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
	pflag.StringP("label_file", "f", "labels.yaml", "Map labels config file to dynamic label in Prometheus metrics")
	pflag.BoolP("version", "v", false, "print version")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		logrus.Error(err)
	}

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("exporter")
	viper.AutomaticEnv()

	if viper.IsSet("label_file") {
		file, err := os.Open(viper.GetString("label_file"))
		if err != nil {
			logrus.Fatalf("Unable to open file: ", err)
		}
		viper.SetConfigType("yaml")
		if err := viper.ReadConfig(file); err != nil {
			logrus.Fatalf("Unable to read file labels: ", err)
		}
	}
}
