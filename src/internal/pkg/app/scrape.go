package app

import (
	"go-ldap-metrics-exporter/internal/pkg/prometheus"
	"go-ldap-metrics-exporter/internal/pkg/structs"
	"os"
	"time"

	ext_prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	log "github.com/sirupsen/logrus"
)

func ScrapeMetrics(config *structs.Config) {
	ticker := time.NewTicker(time.Duration(config.Scrape.Interval))
	defer ticker.Stop()

	prometheus.Init()

	log.Infof("scraping metrics from LDAP server %s every %d seconds", config.LDAP.Address, config.Scrape.Interval)

	for range ticker.C {
		log.Debug("starting metrics scrape")
		prometheus.ScrapeMetrics(config)
	}
}

func ExportMetrics(exportFile string, exportInterval time.Duration) {
	ticker := time.NewTicker(exportInterval)
	defer ticker.Stop()

	for range ticker.C {
		log.Debug("exporting metrics to file")
		if err := writeMetrics(exportFile); err != nil {
			log.Errorf("failed to export metrics: %v", err)
		}
	}
}

func writeMetrics(filePath string) error {
	gatherers := ext_prom.Gatherers{
		ext_prom.DefaultGatherer,
	}

	mfs, err := gatherers.Gather()
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, mf := range mfs {
		if _, err := expfmt.MetricFamilyToText(f, mf); err != nil {
			return err
		}
	}

	return nil
}
