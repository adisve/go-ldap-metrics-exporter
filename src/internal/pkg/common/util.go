package common

import (
	"go-ldap-metrics-exporter/internal/pkg/structs"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ViperLoadConfig(configPath string, config *structs.Config) {
	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath(".")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func SetLogLevel(level string, json bool) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	if json {
		log.SetFormatter(&log.JSONFormatter{})
	}
}
