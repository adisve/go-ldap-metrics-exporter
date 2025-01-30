package prometheus

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"time"

	"go-ldap-metrics-exporter/internal/pkg/ldap"

	"github.com/hashicorp/go-multierror"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	ext_ldap "gopkg.in/ldap.v2"
)

type Metric struct {
	Name       string
	QueryFunc  func(*ext_ldap.Conn, string) (float64, error)
	Gauge      *prometheus.GaugeVec
	LabelValue string
}

/**
 * Metrics to scrape from the LDAP server.
 */
var metrics = []Metric{
	{"people", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.SubordinateQuery(l, fmt.Sprintf("ou=people,%s", suffix), "(objectClass=organizationalUnit)")
	}, common.UsersGauge, "active"},

	{"groups", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.SubordinateQuery(l, fmt.Sprintf("ou=groups,%s", suffix), "(objectClass=organizationalUnit)")
	}, common.GroupsGauge, ""},

	// LDAP server metrics queries, requires access to cn=monitor
	{"current_connections", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "currentconnections")
	}, common.CurrentConnectionsGauge, ""},

	{"total_connections", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "totalconnections")
	}, common.TotalConnectionsGauge, ""},

	{"entries", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "entries")
	}, common.EntriesGauge, ""},

	{"ops_initiated", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "opsinitiated")
	}, common.OperationsInitiatedGauge, ""},

	{"ops_completed", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "opscompleted")
	}, common.OperationsCompletedGauge, ""},

	{"threads", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "threads")
	}, common.ThreadsGauge, ""},

	{"version", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "version")
	}, common.VersionGauge, ""},

	{"bytessent", func(l *ext_ldap.Conn, suffix string) (float64, error) {
		return ldap.MonitorAttributeQuery(l, "cn=monitor", "bytessent")
	}, common.BytesSentGauge, ""},
}

/**
 * ScrapeMetrics scrapes all metrics from the LDAP server.
 */
func ScrapeMetrics(ldapAddr, ldapUser, ldapPass, ipaDomain string) {
	start := time.Now()
	if err := scrapeAllMetrics(ldapAddr, ldapUser, ldapPass, ipaDomain); err != nil {
		common.ScrapeCounter.WithLabelValues("fail").Inc()
		log.Error("scrape failed:", err)
	} else {
		common.ScrapeCounter.WithLabelValues("ok").Inc()
	}
	elapsed := time.Since(start).Seconds()
	common.ScrapeDurationGauge.WithLabelValues().Set(float64(elapsed))
	log.Infof("scrape completed in %f seconds", elapsed)
}

/**
 * Scrape all metrics from the LDAP server.
 */
func scrapeAllMetrics(ldapAddr, ldapUser, ldapPass, ipaDomain string) error {
	dnSuffix := common.IpaDomainToBaseDN(ipaDomain)
	userWithDn := common.UserWithBaseDN(ldapUser, dnSuffix)

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
		log.Debugf("getting %s", metric.Name)
		num, err := metric.QueryFunc(l, dnSuffix)
		if err != nil {
			errs = multierror.Append(errs, err)
		}
		metric.Gauge.WithLabelValues(metric.LabelValue).Set(num)
	}

	log.Debug("getting replication agreements")
	err = ldap.ReplicationQuery(l, dnSuffix)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	return errs
}
