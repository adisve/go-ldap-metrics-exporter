package common

import "github.com/prometheus/client_golang/prometheus"

var (
	UsersGauge  = newGaugeVec("users", "Number of user accounts", []string{"type"})
	GroupsGauge = newGaugeVec("groups", "Number of groups", nil)

	ReplicationConflictsGauge = newGaugeVec("replication_conflicts", "Number of ldap conflicts", nil)
	ReplicationStatusGauge    = newGaugeVec("replication_status", "Replication status by server", []string{"server"})
	ScrapeCounter             = newCounterVec("scrape_count", "successful vs unsuccessful ldap scrape attempts", []string{"result"})
	ScrapeDurationGauge       = newGaugeVec("scrape_duration_seconds", "time taken per scrape", nil)

	/* Requires cn=metrics */
	ReadWaitersGauge                = newGaugeVec("read_waiters", "Number of read waiters", nil)
	DTableSizeGauge                 = newGaugeVec("dtable_size", "Size of the dtable", nil)
	AnonymousBindsGauge             = newGaugeVec("anonymous_binds", "Number of anonymous binds", nil)
	UnauthBindsGauge                = newGaugeVec("unauth_binds", "Number of unauthenticated binds", nil)
	SimpleAuthBindsGauge            = newGaugeVec("simple_auth_binds", "Number of simple authenticated binds", nil)
	StrongAuthBindsGauge            = newGaugeVec("strong_auth_binds", "Number of strong authenticated binds", nil)
	BindSecurityErrorsGauge         = newGaugeVec("bind_security_errors", "Number of bind security errors", nil)
	InOpsGauge                      = newGaugeVec("in_ops", "Number of incoming operations", nil)
	ReadOpsGauge                    = newGaugeVec("read_ops", "Number of read operations", nil)
	CompareOpsGauge                 = newGaugeVec("compare_ops", "Number of compare operations", nil)
	AddEntryOpsGauge                = newGaugeVec("add_entry_ops", "Number of add entry operations", nil)
	ModifyEntryOpsGauge             = newGaugeVec("modify_entry_ops", "Number of modify entry operations", nil)
	RemoveEntryOpsGauge             = newGaugeVec("remove_entry_ops", "Number of remove entry operations", nil)
	ModifyRDNOpsGauge               = newGaugeVec("modify_rdn_ops", "Number of modify rdn operations", nil)
	SearchOpsGauge                  = newGaugeVec("search_ops", "Number of search operations", nil)
	OneLevelSearchOpsGauge          = newGaugeVec("one_level_search_ops", "Number of one level search operations", nil)
	WholeSubtreeSearchOpsGauge      = newGaugeVec("whole_subtree_search_ops", "Number of whole subtree search operations", nil)
	ReferralsGauge                  = newGaugeVec("referral_ops", "Number of referral operations", nil)
	SecurityErrorsGauge             = newGaugeVec("security_errors", "Number of security errors", nil)
	ErrorsGauge                     = newGaugeVec("errors", "Number of errors", nil)
	ConnectionsGauge                = newGaugeVec("connections", "Number of connections", nil)
	ConnectinosInMaxThreadsGauge    = newGaugeVec("connections_in_max_threads", "Number of connections in max threads", nil)
	ConnectionsEqGauge              = newGaugeVec("connections_eq", "Number of connections eq", nil)
	ConnectionsInMaxThreadsGauge    = newGaugeVec("connections_in_max_threads", "Number of connections in max threads", nil)
	ConnectionsMaxThreadsCountGauge = newGaugeVec("connections_max_threads_count", "Number of connections max threads count", nil)
	BytesRecvGauge                  = newGaugeVec("bytes_recv", "Number of bytes received", nil)
	EntriesReturnedGauge            = newGaugeVec("entries_returned", "Number of entries returned", nil)
	ReferralsReturnedGauge          = newGaugeVec("referrals_returned", "Number of referrals returned", nil)
	CacheEntriesGauge               = newGaugeVec("cache_entries", "Number of cache entries", nil)
	CacheHitsGauge                  = newGaugeVec("cache_hits", "Number of cache hits", nil)
	CurrentConnectionsGauge         = newGaugeVec("current_connections", "Current number of connections to the LDAP server", nil)
	TotalConnectionsGauge           = newGaugeVec("total_connections", "Total number of connections to the LDAP server", nil)
	EntriesGauge                    = newGaugeVec("entries", "Number of entries in the LDAP server", nil)
	OperationsCompletedGauge        = newGaugeVec("operations_completed", "Number of operations performed by the LDAP server", nil)
	OperationsInitiatedGauge        = newGaugeVec("operations_initiated", "Number of operations initiated by the LDAP server", nil)
	ThreadsGauge                    = newGaugeVec("threads", "Number of threads in the LDAP server", nil)
	BytesSentGauge                  = newGaugeVec("bytes_sent", "Number of bytes sent by the LDAP server", nil)
	VersionGauge                    = newGaugeVec("version", "LDAP server version", nil)
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
