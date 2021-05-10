package main

import (
	"os"

	"github.com/seggga/csvquery/config"
	"github.com/sirupsen/logrus"
)

// initLogInfo - initializes info logger
func initLogInfo(log *logrus.Logger, file *os.File, conf *config.ConfigType) {

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(file)
	log.Info("logging started")
	log.Info("command-line parameters:")
	log.Infof("timeout: %d", conf.Timeout)
}

// initLogErr - initializes error logger
func initLogErr(log *logrus.Logger, file *os.File, conf *config.ConfigType) {
	log.SetLevel(logrus.ErrorLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(file)
}
