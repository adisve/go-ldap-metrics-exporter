package prometheus

import (
	"go-ldap-metrics-exporter/internal/pkg/common"

	"github.com/prometheus/client_golang/prometheus"
)

func Init() {
	registerMetrics()
}

/**
 * Register all desired metrics with Prometheus.
 */
func registerMetrics() {
	prometheus.MustRegister(
		common.UsersGauge,
		common.GroupsGauge,
		common.ReplicationConflictsGauge,
		common.ReplicationStatusGauge,
		common.ScrapeCounter,
		common.ScrapeDurationGauge,
		common.ReadWaitersGauge,
		common.DTableSizeGauge,
		common.AnonymousBindsGauge,
		common.UnauthBindsGauge,
		common.SimpleAuthBindsGauge,
		common.StrongAuthBindsGauge,
		common.BindSecurityErrorsGauge,
		common.InOpsGauge,
		common.ReadOpsGauge,
		common.CompareOpsGauge,
		common.AddEntryOpsGauge,
		common.ModifyEntryOpsGauge,
		common.RemoveEntryOpsGauge,
		common.ModifyRDNOpsGauge,
		common.SearchOpsGauge,
		common.OneLevelSearchOpsGauge,
		common.WholeSubtreeSearchOpsGauge,
		common.ReferralsGauge,
		common.SecurityErrorsGauge,
		common.ErrorsGauge,
		common.ConnectionsGauge,
		common.BytesRecvGauge,
		common.EntriesReturnedGauge,
		common.ReferralsReturnedGauge,
		common.CacheEntriesGauge,
		common.CacheHitsGauge,
		common.CurrentConnectionsGauge,
		common.TotalConnectionsGauge,
		common.EntriesGauge,
		common.OperationsCompletedGauge,
		common.OperationsInitiatedGauge,
		common.ThreadsGauge,
		common.BytesSentGauge,
		common.VersionGauge,
	)
}
