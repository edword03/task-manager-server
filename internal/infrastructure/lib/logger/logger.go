package logger

import (
	log "github.com/sirupsen/logrus"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string) {
	switch env {
	case envLocal:
		log.SetFormatter(&log.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	case envDev:
		log.SetFormatter(&log.JSONFormatter{})
	case envProd:
		log.SetFormatter(&log.JSONFormatter{})
	}
}
