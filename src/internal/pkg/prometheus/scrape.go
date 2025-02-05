package prometheus

import (
	"errors"
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"go-ldap-metrics-exporter/internal/pkg/structs"
	"strings"
	"sync"
	"time"

	"net/url"

	log "github.com/sirupsen/logrus"
	ext_ldap "gopkg.in/ldap.v2"
)

func ScrapeMetrics(config *structs.Config) {
	start := time.Now()

	if err := scrapeAllMetrics(config); err != nil {
		common.ScrapeCounter.WithLabelValues("fail").Inc()
		log.Error("scrape failed: ", err)
	} else {
		common.ScrapeCounter.WithLabelValues("ok").Inc()
	}
	elapsed := time.Since(start).Seconds()
	common.ScrapeDurationGauge.WithLabelValues().Set(float64(elapsed))
	log.Debugf("scrape completed in %f seconds", elapsed)
}

func sanitizeLDAPAddress(ldapAddr string) (string, error) {
	if !strings.Contains(ldapAddr, "://") {
		return ldapAddr, nil
	}

	parsedURL, err := url.Parse(ldapAddr)
	if err != nil {
		return "", err
	}

	host := parsedURL.Host
	if host == "" {
		return "", fmt.Errorf("missing host in LDAP address")
	}

	if !strings.Contains(host, ":") {
		if parsedURL.Scheme == "ldaps" {
			host += ":636"
		} else {
			host += ":389"
		}
	}

	return host, nil
}

func scrapeAllMetrics(config *structs.Config) error {
	sanitizedLDAPAddress, err := sanitizeLDAPAddress(config.LDAP.Address)
	if err != nil {
		log.Errorf("Invalid LDAP address: %s", err)
		return err
	}

	l, err := ext_ldap.Dial("tcp", sanitizedLDAPAddress)
	if err != nil {
		log.Errorf("Failed to connect to LDAP server at %s: %v", sanitizedLDAPAddress, err)
		return err
	}
	defer l.Close()

	fullDn := fmt.Sprintf("uid=%s,%s", config.LDAP.Username, config.LDAP.UserBaseDN)
	log.Debugf("Connecting to %s as %s", config.LDAP.Address, fullDn)

	if err := l.Bind(fullDn, config.LDAP.Password); err != nil {
		log.Errorf("LDAP bind failed for user %s: %v", fullDn, err)
		return err
	}

	var errs []error
	errs = append(errs, scrapeMetrics(l, replication_metrics)...)
	errs = append(errs, scrapeMetrics(l, monitor_metrics)...)

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func scrapeMetrics(l *ext_ldap.Conn, metrics []Metric) []error {
	var errs []error
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, metric := range metrics {
		wg.Add(1)
		go func(metric Metric) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					log.Errorf("Panic while processing metric %s: %v", metric.Name, r)
					mu.Lock()
					errs = append(errs, fmt.Errorf("panic while scraping %s: %v", metric.Name, r))
					mu.Unlock()
				}
			}()

			log.Debugf("Getting '%s'", metric.Name)

			done := make(chan struct{})
			go func() {
				collectMetric(l, metric)
				close(done)
			}()

			select {
			case <-done:
				log.Debugf("Successfully scraped %s", metric.Name)
			case <-time.After(5 * time.Second):
				log.Errorf("Metric %s is taking too long!", metric.Name)
				mu.Lock()
				errs = append(errs, fmt.Errorf("timeout while scraping %s", metric.Name))
				mu.Unlock()
			}
		}(metric)
	}

	wg.Wait()
	return errs
}
