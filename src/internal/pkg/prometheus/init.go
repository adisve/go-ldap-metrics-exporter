package prometheus

import (
	"go-ldap-metrics-exporter/internal/pkg/common"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

func Init() {
	log.Info("initializing Prometheus metrics")
	registerMetrics()
}

/**
 * Register all desired metrics with Prometheus.
 */
func registerMetrics() {
	prometheus.MustRegister(
		common.ReplicationConflictsGauge,
		common.ReplicationStatusGauge,
		common.ScrapeCounter,
		common.ScrapeDurationGauge,
		common.CurrentConnectionsGauge,
		common.TotalConnectionsGauge,
		common.CurrentConnectionsAtMaxThreadsGauge,
		common.MaxThreadsPerConnHitsGauge,
		common.DTableSizeGauge,
		common.ReadWaitersGauge,
		common.OpsInitiatedGauge,
		common.OpsCompletedGauge,
		common.EntriesSentGauge,
		common.BytesSentGauge,
		common.CurrentTimeGauge,
		common.StartTimeGauge,
		common.NBackendsGauge,
		common.AnonymousBindsGauge,
		common.UnauthBindsGauge,
		common.SimpleAuthBindsGauge,
		common.StrongAuthBindsGauge,
		common.BindSecurityErrorsGauge,
		common.InOpsGauge,
		common.ListOpsGauge,
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
		common.ChainingsGauge,
		common.SecurityErrorsGauge,
		common.ErrorsGauge,
		common.ConnectionsGauge,
		common.ConnectionsInMaxThreadsGauge,
		common.ConnectionsMaxThreadsCountGauge,
		common.ConnectionsEqGauge,
		common.BytesRecvGauge,
		common.EntriesReturnedGauge,
		common.ReferralsReturnedGauge,
		common.SupplierEntriesGauge,
		common.CopyEntriesGauge,
		common.CacheEntriesGauge,
		common.CacheHitsGauge,
		common.ConsumerHitsGauge,
		common.DsDiskGauge,
	)
}
