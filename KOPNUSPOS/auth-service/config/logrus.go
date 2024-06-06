package config

import (
	"fmt"
	"strconv"
	"time"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/sirupsen/logrus"
)

type EnvLogrusConfig struct {
	LogLevel    string `mapstructure:"LOG_LEVEL" validate:"required"`
	GraylogHost string `mapstructure:"GRAYLOG_HOST" validate:"required"`
	GraylogPort string `mapstructure:"GRAYLOG_PORT" validate:"required"`
}

var (
	LogrusConfig EnvLogrusConfig
)

func InitLogrus(levelEnv, graylogHost, graylogPort string) *logrus.Logger {
	log := logrus.New()
	level, err := strconv.ParseUint(levelEnv, 10, 64)
	if err != nil {
		log.Println("Failed to setup logrus level log, error cause:", err)
	}

	if graylogHost != "" && graylogPort != "" {
		graylogAddress := fmt.Sprintf("%s:%s", graylogHost, graylogPort)
		hook := graylog.NewGraylogHook(graylogAddress, map[string]interface{}{
			"version":   "1.1",
			"host":      "auth-service",
			"facility":  "auth-service-facility",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		log.Hooks.Add(hook)

		log.Info("Graylog connection test - application startup")
	}

	logLevel := uint32(level)
	log.SetLevel(logrus.Level(logLevel))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
