// Package log provide initial function to setup log format and report level
package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Init is function to initialize log format and log level
func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		EnvironmentOverrideColors: false,
		DisableColors:             false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "01/02/2006 03:04:05 PM",
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          false,
		FieldMap: logrus.FieldMap{
			"time":  "Timestamp",
			"level": "Level",
			"msg":   "Message",
		},
	})
	logrus.SetReportCaller(viper.GetBool("caller"))

	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(os.Stdout)
}
