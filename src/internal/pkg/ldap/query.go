package ldap

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ldap.v2"
)

var ldapEscaper = strings.NewReplacer("=", "\\=", ",", "\\,")

/**
 * Query the LDAP server for the number of subordinates under a given base DN.
 */
func CollectSubordinateMetrics(l *ldap.Conn, baseDN, searchFilter string, gauge *prometheus.GaugeVec, label string) {
	req := ldap.NewSearchRequest(
		baseDN, ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0, 0, false,
		searchFilter, []string{"numSubordinates"}, nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		log.Errorf("LDAP query failed for %s: %v", baseDN, err)
		return
	}

	if len(sr.Entries) == 0 {
		log.Warnf("No entries found for %s (%s)", baseDN, searchFilter)
		return
	}

	val := sr.Entries[0].GetAttributeValue("numSubordinates")
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Errorf("Failed to parse numSubordinates for %s: %v", baseDN, err)
		return
	}

	gauge.WithLabelValues(label).Set(num)
}

/**
 * Query the LDAP server for the number of entries matching a given filter.
 */
func CollectEntryCountMetrics(l *ldap.Conn, baseDN, searchFilter, attr string, scope int, gauge *prometheus.GaugeVec, label string) {
	req := ldap.NewSearchRequest(
		baseDN, scope, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter, []string{attr}, nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		log.Errorf("LDAP CountQuery failed for %s (%s): %v", baseDN, searchFilter, err)
		return
	}

	num := float64(len(sr.Entries))
	gauge.WithLabelValues(label).Set(num)
}

/**
 * Query the LDAP server for replication status.
 */
func CollectReplicationMetrics(l *ldap.Conn, suffix string, gauge *prometheus.GaugeVec) error {
	escaped_suffix := ldapEscaper.Replace(suffix)
	base_dn := fmt.Sprintf("cn=replica,cn=%s,cn=mapping tree,cn=config", escaped_suffix)

	req := ldap.NewSearchRequest(
		base_dn, ldap.ScopeSingleLevel,
		ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=nsds5replicationagreement)",
		[]string{"nsDS5ReplicaHost", "nsds5replicaLastUpdateStatus", "nsds5replicaConflictCount"},
		nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		return err
	}

	for _, entry := range sr.Entries {
		host := entry.GetAttributeValue("nsDS5ReplicaHost")
		status := entry.GetAttributeValue("nsds5replicaLastUpdateStatus")
		conflictStr := entry.GetAttributeValue("nsds5replicaConflictCount")

		if strings.Contains(status, "Incremental update succeeded") {
			gauge.WithLabelValues(host).Set(1)
		} else {
			log.Warnf("Replication issue detected on host: %s, status: %s", host, status)
			gauge.WithLabelValues(host).Set(0)
		}

		conflictCount, err := strconv.ParseFloat(conflictStr, 64)
		if err != nil {
			log.Warnf("Failed to parse replication conflict count for host: %s, error: %v", host, err)
			conflictCount = -1
		}
		common.ReplicationConflictsGauge.WithLabelValues(host).Set(conflictCount)
	}

	return nil
}

/**
 * Query the LDAP server for a specific monitor attribute.
 */
func CollectMonitorMetrics(l *ldap.Conn, attribute string, gauge *prometheus.GaugeVec) {
	req := ldap.NewSearchRequest(
		"cn=monitor",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)",
		[]string{attribute},
		nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		log.Errorf("Failed to query monitor attribute %s: %v", attribute, err)
		return
	}

	if len(sr.Entries) == 0 {
		log.Warnf("No entries found for monitor attribute %s", attribute)
		return
	}

	val := sr.Entries[0].GetAttributeValue(attribute)
	if val == "" {
		log.Warnf("No value found for monitor attribute %s", attribute)
		return
	}
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Errorf("Failed to parse monitor attribute %s: %v", attribute, err)
		return
	}

	gauge.WithLabelValues().Set(num)
}
