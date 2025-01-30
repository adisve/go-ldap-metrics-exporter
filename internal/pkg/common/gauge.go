package common

import "github.com/prometheus/client_golang/prometheus"

var (
	UsersGauge               = newGaugeVec("users", "Number of user accounts", []string{"type"})
	GroupsGauge              = newGaugeVec("groups", "Number of groups", nil)
	HostsGauge               = newGaugeVec("hosts", "Number of hosts", nil)
	HostGroupsGauge          = newGaugeVec("hostgroups", "Number of hostgroups", nil)
	HbacRulesGauge           = newGaugeVec("hbac_rules", "Number of hbac rules", nil)
	SudoRulesGauge           = newGaugeVec("sudo_rules", "Number of sudo rules", nil)
	DnsZonesGauge            = newGaugeVec("dns_zones", "Number of dns zones", nil)
	LdapConflictsGauge       = newGaugeVec("replication_conflicts", "Number of ldap conflicts", nil)
	ReplicationStatusGauge   = newGaugeVec("replication_status", "Replication status by server", []string{"server"})
	ScrapeCounter            = newCounterVec("scrape_count", "successful vs unsuccessful ldap scrape attempts", []string{"result"})
	ScrapeDurationGauge      = newGaugeVec("scrape_duration_seconds", "time taken per scrape", nil)
	CurrentConnectionsGauge  = newGaugeVec("current_connections", "Current number of connections to the LDAP server", nil)
	TotalConnectionsGauge    = newGaugeVec("total_connections", "Total number of connections to the LDAP server", nil)
	EntriesGauge             = newGaugeVec("entries", "Number of entries in the LDAP server", nil)
	OperationsCompletedGauge = newGaugeVec("operations_completed", "Number of operations performed by the LDAP server", nil)
	OperationsInitiatedGauge = newGaugeVec("operations_initiated", "Number of operations initiated by the LDAP server", nil)
	ThreadsGauge             = newGaugeVec("threads", "Number of threads in the LDAP server", nil)
	BytesSentGauge           = newGaugeVec("bytes_sent", "Number of bytes sent by the LDAP server", nil)
	VersionGauge             = newGaugeVec("version", "LDAP server version", nil)
)

const (
	subsystem = "ldap_389ds"
)

/**
 * Create a new gauge metric with labels.
 */
func newGaugeVec(name, help string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}

/**
 * Create a new counter metric with labels.
 */
func newCounterVec(name, help string, labels []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}
