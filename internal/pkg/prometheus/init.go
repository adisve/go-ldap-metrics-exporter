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
		common.HostsGauge,
		common.HostGroupsGauge,
		common.HbacRulesGauge,
		common.SudoRulesGauge,
		common.DnsZonesGauge,
		common.LdapConflictsGauge,
		common.ReplicationStatusGauge,
		common.ScrapeCounter,
		common.ScrapeDurationGauge,
		common.CurrentConnectionsGauge,
		common.TotalConnectionsGauge,
		common.EntriesGauge,
		common.OperationsCompletedGauge,
		common.OperationsInitiatedGauge,
		common.ThreadsGauge,
	)
}
