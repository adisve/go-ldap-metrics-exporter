package exporter

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/app"
	"go-ldap-metrics-exporter/internal/pkg/structs"

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

	scrapeDone := make(chan struct{})

	go app.ScrapeMetrics(config, scrapeDone)
	go app.ExportMetrics(config.Export.File, scrapeDone)

	select {}
}
