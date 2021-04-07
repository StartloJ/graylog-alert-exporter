// Package config provide environment and global variables in program
package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Init is function to initialize config variables
func Init() {
	pflag.String("listen", "0.0.0.0:9889", "Host address to start service listen")
	pflag.String("path", "metrics", "path for scape and push metrics")
	pflag.Int("timeout", 60, "timeout of alert to make it resolved")
	pflag.Int("interval", 5, "interval to check timeout (lower value consume more cpu)")
	pflag.Bool("debug", false, "enable debug log")
	pflag.Bool("caller", false, "enable log method caller in code")
	pflag.Bool("dashboard", false, "enable dashboard web service listen at path dashboard")
	pflag.Bool("version", false, "print version")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetEnvPrefix("app")
	viper.BindEnv("listen")
	viper.BindEnv("path")
	viper.BindEnv("timeout")
	viper.BindEnv("timeout")
	viper.BindEnv("debug")
	viper.BindEnv("caller")
	viper.BindEnv("dashboard")
}
