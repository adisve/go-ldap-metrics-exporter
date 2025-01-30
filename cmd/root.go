package cmd

import (
	"go-ldap-metrics-exporter/internal/exporter"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"go-ldap-metrics-exporter/internal/pkg/structs"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	config structs.Config
)

var rootCmd = &cobra.Command{
	Use:   "go-ldap-metrics-exporter",
	Short: "LDAP Exporter for Prometheus",
	Run: func(cmd *cobra.Command, args []string) {
		exporter.Start(&config)
	},
}

func Execute() {

	var configPath string
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to the config file. (--config/-c)")

	common.ViperLoadConfig(configPath, &config)
	common.SetLogLevel(config.Log.Level, config.Log.JSON)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}
