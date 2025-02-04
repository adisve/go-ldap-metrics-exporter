package prometheus

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"go-ldap-metrics-exporter/internal/pkg/ldap"

	"github.com/prometheus/client_golang/prometheus"
	ext_ldap "gopkg.in/ldap.v2"
)

type Metric struct {
	Name       string
	QueryFunc  func(*ext_ldap.Conn, string, *prometheus.GaugeVec, string)
	Gauge      *prometheus.GaugeVec
	LabelValue string
}

/**
 * Metrics to scrape from the LDAP server.
 */
var metrics = []Metric{
	{"people", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectSubordinateMetrics(l, fmt.Sprintf("ou=people,%s", suffix), "(objectClass=organizationalUnit)", gauge, label)
	}, common.UsersGauge, "active"},

	{"groups", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectSubordinateMetrics(l, fmt.Sprintf("ou=groups,%s", suffix), "(objectClass=organizationalUnit)", gauge, label)
	}, common.GroupsGauge, ""},

	{"replication_conflicts", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectReplicationMetrics(l, suffix, gauge)
	}, common.ReplicationConflictsGauge, ""},

	{"replication_status", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectReplicationMetrics(l, suffix, gauge)
	}, common.ReplicationStatusGauge, ""},

	{"readwaiters", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "readwaiters", gauge)
	}, common.ReadWaitersGauge, ""},

	{"dtablesize", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "dtablesize", gauge)
	}, common.DTableSizeGauge, ""},

	{"anonymousbinds", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "anonymousbinds", gauge)
	}, common.AnonymousBindsGauge, ""},

	{"unauthbinds", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "unauthbinds", gauge)
	}, common.UnauthBindsGauge, ""},

	{"simpleauthbinds", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "simpleauthbinds", gauge)
	}, common.SimpleAuthBindsGauge, ""},

	{"strongauthbinds", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "strongauthbinds", gauge)
	}, common.StrongAuthBindsGauge, ""},

	{"securityerrors", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "securityerrors", gauge)
	}, common.SecurityErrorsGauge, ""},

	{"errors", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "errors", gauge)
	}, common.ErrorsGauge, ""},

	{"connections", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "connections", gauge)
	}, common.ConnectionsGauge, ""},

	{"bytessent", func(l *ext_ldap.Conn, suffix string, gauge *prometheus.GaugeVec, label string) {
		ldap.CollectMonitorMetrics(l, "bytessent", gauge)
	}, common.BytesSentGauge, ""},
}
