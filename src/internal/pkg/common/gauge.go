package common

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	/* dn: cn=replication,cn=monitor */
	ReplicationConflictsGauge = newGaugeVec("replicationconflicts", "Number of ldap conflicts", []string{})
	ReplicationStatusGauge    = newGaugeVec("replicationstatus", "Replication status by server", []string{"server"})
	ScrapeCounter             = newGaugeVec("scrapecount", "successful vs unsuccessful ldap scrape attempts", []string{"result"})
	ScrapeDurationGauge       = newGaugeVec("scrapedurationseconds", "time taken per scrape", []string{})

	/* dn: cn=monitor */
	CurrentConnectionsGauge             = newGaugeVec("currentconnections", "Current number of connections to the LDAP server", []string{})
	TotalConnectionsGauge               = newGaugeVec("totalconnections", "Total number of connections to the LDAP server", []string{})
	CurrentConnectionsAtMaxThreadsGauge = newGaugeVec("currentconnectionsatmaxthreads", "Current number of connections at max threads", []string{})
	MaxThreadsPerConnHitsGauge          = newGaugeVec("maxthreadsperconnhits", "Max threads per connection", []string{})
	DTableSizeGauge                     = newGaugeVec("dtablesize", "Size of the dtable", []string{})
	ReadWaitersGauge                    = newGaugeVec("readwaiters", "Number of read waiters", []string{})
	OpsInitiatedGauge                   = newGaugeVec("opsinitiated", "Number of operations initiated", []string{})
	OpsCompletedGauge                   = newGaugeVec("opscompleted", "Number of operations completed", []string{})
	EntriesSentGauge                    = newGaugeVec("entriessent", "Number of entries sent", []string{})
	BytesSentGauge                      = newGaugeVec("bytessent", "Number of bytes sent", []string{})
	CurrentTimeGauge                    = newGaugeVec("currenttime", "Current time", []string{})
	StartTimeGauge                      = newGaugeVec("starttime", "Start time", []string{})
	NBackendsGauge                      = newGaugeVec("nbackends", "Number of backends", []string{})

	/* dn: cn=snmp,cn=monitor */
	AnonymousBindsGauge             = newGaugeVec("anonymousbinds", "Number of anonymous binds", []string{})
	UnauthBindsGauge                = newGaugeVec("unauthbinds", "Number of unauthenticated binds", []string{})
	SimpleAuthBindsGauge            = newGaugeVec("simpleauthbinds", "Number of simple authenticated binds", []string{})
	StrongAuthBindsGauge            = newGaugeVec("strongauthbinds", "Number of strong authenticated binds", []string{})
	BindSecurityErrorsGauge         = newGaugeVec("bindsecurityerrors", "Number of bind security errors", []string{})
	InOpsGauge                      = newGaugeVec("inops", "Number of incoming operations", []string{})
	ListOpsGauge                    = newGaugeVec("listops", "Number of list operations", []string{})
	ReadOpsGauge                    = newGaugeVec("readops", "Number of read operations", []string{})
	CompareOpsGauge                 = newGaugeVec("compareops", "Number of compare operations", []string{})
	AddEntryOpsGauge                = newGaugeVec("addentryops", "Number of add entry operations", []string{})
	ModifyEntryOpsGauge             = newGaugeVec("modifyentryops", "Number of modify entry operations", []string{})
	RemoveEntryOpsGauge             = newGaugeVec("removeentryops", "Number of remove entry operations", []string{})
	ModifyRDNOpsGauge               = newGaugeVec("modifyrdnops", "Number of modify rdn operations", []string{})
	SearchOpsGauge                  = newGaugeVec("searchops", "Number of search operations", []string{})
	OneLevelSearchOpsGauge          = newGaugeVec("onelevelsearchops", "Number of one level search operations", []string{})
	WholeSubtreeSearchOpsGauge      = newGaugeVec("wholesubtreesearchops", "Number of whole subtree search operations", []string{})
	ReferralsGauge                  = newGaugeVec("referrals", "Number of referral operations", []string{})
	ChainingsGauge                  = newGaugeVec("chainings", "Number of chaining operations", []string{})
	SecurityErrorsGauge             = newGaugeVec("securityerrors", "Number of security errors", []string{})
	ErrorsGauge                     = newGaugeVec("errors", "Number of errors", []string{})
	ConnectionsGauge                = newGaugeVec("connections", "Number of connections", []string{})
	ConnectionsInMaxThreadsGauge    = newGaugeVec("connectionsinmaxthreads", "Number of connections at max threads", []string{})
	ConnectionsMaxThreadsCountGauge = newGaugeVec("connectionsmaxthreadsount", "Max number of connections", []string{})
	ConnectionsEqGauge              = newGaugeVec("connectionseq", "Number of connections equal", []string{})
	BytesRecvGauge                  = newGaugeVec("bytesrecv", "Number of bytes received", []string{})
	EntriesReturnedGauge            = newGaugeVec("entriesreturned", "Number of entries returned", []string{})
	ReferralsReturnedGauge          = newGaugeVec("referralsreturned", "Number of referrals returned", []string{})
	SupplierEntriesGauge            = newGaugeVec("supplierentries", "Number of supplier entries", []string{})
	CopyEntriesGauge                = newGaugeVec("copyentries", "Number of copy entries", []string{})
	CacheEntriesGauge               = newGaugeVec("cacheentries", "Number of cache entries", []string{})
	CacheHitsGauge                  = newGaugeVec("cachehits", "Number of cache hits", []string{})
	ConsumerHitsGauge               = newGaugeVec("consumerhits", "Number of consumer hits", []string{})

	/* dn: cn=disk space,cn=monitor */
	DsDiskGauge = newGaugeVec("dsdisk", "Disk space used", []string{"partition", "metric_type"})
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
