package prometheus

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"strings"
	"time"

	"net/url"

	log "github.com/sirupsen/logrus"
	ext_ldap "gopkg.in/ldap.v2"
)

func ScrapeMetrics(ldapAddr, ldapUser, ldapPass, ipaDomain string) {
	start := time.Now()

	parsedAddr, err := sanitizeLDAPAddress(ldapAddr)
	if err != nil {
		log.Errorf("Invalid LDAP address: %s", err)
		return
	}

	if err := scrapeAllMetrics(parsedAddr, ldapUser, ldapPass, ipaDomain); err != nil {
		common.ScrapeCounter.WithLabelValues("fail").Inc()
		log.Error("scrape failed: ", err)
	} else {
		common.ScrapeCounter.WithLabelValues("ok").Inc()
	}
	elapsed := time.Since(start).Seconds()
	common.ScrapeDurationGauge.WithLabelValues().Set(float64(elapsed))
	log.Infof("scrape completed in %f seconds", elapsed)
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

func scrapeAllMetrics(ldapAddr, ldapUser, ldapPass, ipaDomain string) error {
	dnSuffix := common.IpaDomainToBaseDN(ipaDomain)
	userWithDn := common.UserWithBaseDN(ldapUser, dnSuffix)

	log.Infof("Connecting to %s as %s", ldapAddr, userWithDn)

	l, err := ext_ldap.Dial("tcp", ldapAddr)
	if err != nil {
		return err
	}
	defer l.Close()

	err = l.Bind(userWithDn, ldapPass)
	if err != nil {
		return err
	}

	var errs error
	for _, metric := range metrics {
		log.Debugf("Getting %s", metric.Name)
		metric.QueryFunc(l, dnSuffix, metric.Gauge, metric.LabelValue)
	}

	return errs
}
