package main

import (
	"komodorIntegrations/config"

	log "github.com/sirupsen/logrus"
)

func initializeLog() {
	setupLogFormat()
	setupLogLevel()
}

func setupLogFormat() {
	if config.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func setupLogLevel() {
	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
}
