package prometheus

import (
	"go-ldap-metrics-exporter/internal/pkg/common"
	"go-ldap-metrics-exporter/internal/pkg/ldap"

	"github.com/prometheus/client_golang/prometheus"
	ext_ldap "gopkg.in/ldap.v2"
)

type Metric struct {
	Name   string
	Gauge  *prometheus.GaugeVec
	Suffix string
}

func collectMetric(l *ext_ldap.Conn, metric Metric) {
	ldap.CollectMonitorMetrics(l, metric.Suffix, metric.Name, metric.Gauge)
}

var replication_metrics = []Metric{

	/* cn=replication,cn=monitor */
	{"replicationconflicts", common.ReplicationConflictsGauge, "cn=replication,cn=monitor"},
	{"replicationstatus", common.ReplicationStatusGauge, "cn=replication,cn=monitor"},
}

var monitor_metrics = []Metric{

	/* cn=monitor */
	{"currentconnections", common.CurrentConnectionsGauge, "cn=monitor"},
	{"totalconnections", common.TotalConnectionsGauge, "cn=monitor"},
	{"currentconnectionsatmaxthreads", common.CurrentConnectionsAtMaxThreadsGauge, "cn=monitor"},
	{"maxthreadsperconnhits", common.MaxThreadsPerConnHitsGauge, "cn=monitor"},
	{"dtablesize", common.DTableSizeGauge, "cn=monitor"},
	{"readwaiters", common.ReadWaitersGauge, "cn=monitor"},
	{"opsinitiated", common.OpsInitiatedGauge, "cn=monitor"},
	{"opscompleted", common.OpsCompletedGauge, "cn=monitor"},
	{"entriessent", common.EntriesSentGauge, "cn=monitor"},
	{"bytessent", common.BytesSentGauge, "cn=monitor"},
	{"currenttime", common.CurrentTimeGauge, "cn=monitor"},
	{"starttime", common.StartTimeGauge, "cn=monitor"},
	{"nbackends", common.NBackendsGauge, "cn=monitor"},

	/* cn=snmp,cn=monitor */
	{"anonymousbinds", common.AnonymousBindsGauge, "cn=snmp,cn=monitor"},
	{"unauthbinds", common.UnauthBindsGauge, "cn=snmp,cn=monitor"},
	{"simpleauthbinds", common.SimpleAuthBindsGauge, "cn=snmp,cn=monitor"},
	{"strongauthbinds", common.StrongAuthBindsGauge, "cn=snmp,cn=monitor"},
	{"bindsecurityerrors", common.BindSecurityErrorsGauge, "cn=snmp,cn=monitor"},
	{"inops", common.InOpsGauge, "cn=snmp,cn=monitor"},
	{"listops", common.ListOpsGauge, "cn=snmp,cn=monitor"},
	{"readops", common.ReadOpsGauge, "cn=snmp,cn=monitor"},
	{"compareops", common.CompareOpsGauge, "cn=snmp,cn=monitor"},
	{"addentryops", common.AddEntryOpsGauge, "cn=snmp,cn=monitor"},
	{"modifyentryops", common.ModifyEntryOpsGauge, "cn=snmp,cn=monitor"},
	{"removeentryops", common.RemoveEntryOpsGauge, "cn=snmp,cn=monitor"},
	{"modifyrdnops", common.ModifyRDNOpsGauge, "cn=snmp,cn=monitor"},
	{"searchops", common.SearchOpsGauge, "cn=snmp,cn=monitor"},
	{"onelevelsearchops", common.OneLevelSearchOpsGauge, "cn=snmp,cn=monitor"},
	{"wholesubtreesearchops", common.WholeSubtreeSearchOpsGauge, "cn=snmp,cn=monitor"},
	{"referrals", common.ReferralsGauge, "cn=snmp,cn=monitor"},
	{"chainings", common.ChainingsGauge, "cn=snmp,cn=monitor"},
	{"securityerrors", common.SecurityErrorsGauge, "cn=snmp,cn=monitor"},
	{"errors", common.ErrorsGauge, "cn=snmp,cn=monitor"},
	{"connections", common.ConnectionsGauge, "cn=snmp,cn=monitor"},
	{"connectionsinmaxthreads", common.ConnectionsInMaxThreadsGauge, "cn=snmp,cn=monitor"},
	{"connectionsmaxthreadscount", common.ConnectionsMaxThreadsCountGauge, "cn=snmp,cn=monitor"},
	{"connectionseq", common.ConnectionsEqGauge, "cn=snmp,cn=monitor"},
	{"bytesrecv", common.BytesRecvGauge, "cn=snmp,cn=monitor"},
	{"entriesreturned", common.EntriesReturnedGauge, "cn=snmp,cn=monitor"},
	{"referralsreturned", common.ReferralsReturnedGauge, "cn=snmp,cn=monitor"},
	{"supplierentries", common.SupplierEntriesGauge, "cn=snmp,cn=monitor"},
	{"copyentries", common.CopyEntriesGauge, "cn=snmp,cn=monitor"},
	{"cacheentries", common.CacheEntriesGauge, "cn=snmp,cn=monitor"},
	{"cachehits", common.CacheHitsGauge, "cn=snmp,cn=monitor"},
	{"consumerhits", common.ConsumerHitsGauge, "cn=snmp,cn=monitor"},

	/* cn=disk space,cn=monitor */
	{"dsdisk", common.DsDiskGauge, "cn=disk space,cn=monitor"},
}
