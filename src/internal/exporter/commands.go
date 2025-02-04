package exporter

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/app"
	"go-ldap-metrics-exporter/internal/pkg/structs"
	"time"

	log "github.com/sirupsen/logrus"
)

func Start(config *structs.Config) {
	log.Infof("starting go-ldap-metrics-exporter using LDAP address %s", config.LDAP.Address)

	if config.Server.Active {
		serverAddrFull := fmt.Sprintf("%s:%s", config.Server.Address, config.Server.Port)
		log.Infof("starting prometheus HTTP metrics server on %s", serverAddrFull)
		metricsServer := app.StartServer(serverAddrFull, "/metrics")
		defer app.StopServer(metricsServer)
	}

	go app.ExportMetrics(config.Export.File, time.Duration(config.Export.Interval))
	app.ScrapeMetrics(config)
}
