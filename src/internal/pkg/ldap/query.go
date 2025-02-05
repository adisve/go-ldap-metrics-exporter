package ldap

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ldap.v2"
)

/**
 * Query the LDAP server for replication status.
 */
func CollectReplicationMetrics(l *ldap.Conn, suffix string, gauge *prometheus.GaugeVec) error {
	base_dn := fmt.Sprintf("cn=replica,cn=%s,cn=mapping tree,cn=config", suffix)

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
		conflict_str := entry.GetAttributeValue("nsds5replicaConflictCount")

		if strings.Contains(status, "Incremental update succeeded") {
			gauge.WithLabelValues(host).Set(1)
		} else {
			log.Warnf("Replication issue detected on host: %s, status: %s", host, status)
			gauge.WithLabelValues(host).Set(0)
		}

		conflict_count, err := strconv.ParseFloat(conflict_str, 64)
		if err != nil {
			log.Warnf("Failed to parse replication conflict count for host: %s, error: %v", host, err)
			conflict_count = -1
		}
		common.ReplicationConflictsGauge.WithLabelValues(host).Set(conflict_count)
	}

	return nil
}

func parseMonitorAttributes(attrValue string) (map[string]string, error) {
	result := make(map[string]string)
	parts := strings.Fields(attrValue)

	for _, part := range parts {
		keyVal := strings.SplitN(part, "=", 2)
		if len(keyVal) != 2 {
			return nil, fmt.Errorf("invalid attribute format: %s", part)
		}
		key := keyVal[0]
		value := strings.Trim(keyVal[1], "\"")
		result[key] = value
	}
	return result, nil
}

func parseLDAPTime(ldapTime string) (float64, error) {
	layout := "20060102150405Z"
	t, err := time.Parse(layout, ldapTime)
	if err != nil {
		return 0, fmt.Errorf("failed to parse LDAP time: %s", ldapTime)
	}
	return float64(t.Unix()), nil
}

/**
 * Query the LDAP server for a specific monitor attribute.
 */
func CollectMonitorMetrics(l *ldap.Conn, base_dn string, attribute string, gauge *prometheus.GaugeVec) {
	req := ldap.NewSearchRequest(
		base_dn,
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

	// special cases for attributes that need custom parsing
	switch attribute {
	case "dsdisk":
		parsedAttributes, err := parseMonitorAttributes(val)
		if err != nil {
			log.Errorf("Failed to parse monitor attribute %s: %v", attribute, err)
			return
		}

		partition := parsedAttributes["partition"]
		size, _ := strconv.ParseFloat(parsedAttributes["size"], 64)
		used, _ := strconv.ParseFloat(parsedAttributes["used"], 64)
		available, _ := strconv.ParseFloat(parsedAttributes["available"], 64)
		usePercent, _ := strconv.ParseFloat(parsedAttributes["use%"], 64)

		gauge.WithLabelValues(partition, "size").Set(size)
		gauge.WithLabelValues(partition, "used").Set(used)
		gauge.WithLabelValues(partition, "available").Set(available)
		gauge.WithLabelValues(partition, "use_percent").Set(usePercent)
		log.Debugf("Collected disk metrics for partition %s: size=%f, used=%f, available=%f, use%%=%f", partition, size, used, available, usePercent)
		return

	case "starttime", "currenttime":
		timestamp, err := parseLDAPTime(val)
		if err != nil {
			log.Errorf("Failed to parse monitor attribute %s: %v", attribute, err)
			return
		}
		gauge.WithLabelValues().Set(timestamp)
		log.Debugf("Collected timestamp metric %s: %f", attribute, timestamp)
		return
	}

	// directly parse as float
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Errorf("Failed to parse monitor attribute %s: %v", attribute, err)
		return
	}

	gauge.WithLabelValues().Set(num)
	log.Debugf("Collected metric %s: %f", attribute, num)
}
