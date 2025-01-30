package ldap

import (
	"fmt"
	"go-ldap-metrics-exporter/internal/pkg/common"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ldap.v2"
)

var ldapEscaper = strings.NewReplacer("=", "\\=", ",", "\\,")

/**
 * Query the LDAP server for the number of subordinates under a given base DN.
 */
func SubordinateQuery(l *ldap.Conn, baseDN, searchFilter string) (float64, error) {
	req := ldap.NewSearchRequest(
		baseDN, ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter, []string{"numSubordinates"}, nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		return -1, err
	}

	if len(sr.Entries) == 0 {
		return -1, fmt.Errorf("no entries contain numSubordinates for %s (%s)", baseDN, searchFilter)
	}

	val := sr.Entries[0].GetAttributeValue("numSubordinates")
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, err
	}

	return num, nil
}

/**
 * Query the LDAP server for the number of entries matching a given filter.
 */
func CountQuery(l *ldap.Conn, baseDN, searchFilter, attr string, scope int) (float64, error) {
	req := ldap.NewSearchRequest(
		baseDN, scope, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter, []string{attr}, nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		return -1, err
	}

	num := float64(len(sr.Entries))

	return num, nil
}

/**
 * Query the LDAP server for replication status.
 */
func ReplicationQuery(l *ldap.Conn, suffix string) error {
	escaped_suffix := ldapEscaper.Replace(suffix)
	base_dn := fmt.Sprintf("cn=replica,cn=%s,cn=mapping tree,cn=config", escaped_suffix)

	req := ldap.NewSearchRequest(
		base_dn, ldap.ScopeSingleLevel, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=nsds5replicationagreement)", []string{"nsDS5ReplicaHost", "nsds5replicaLastUpdateStatus"}, nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		return err
	}

	for _, entry := range sr.Entries {
		host := entry.GetAttributeValue("nsDS5ReplicaHost")
		status := entry.GetAttributeValue("nsds5replicaLastUpdateStatus")
		if strings.Contains(status, "Incremental update succeeded") {
			common.ReplicationStatusGauge.WithLabelValues(host).Set(1)
		} else if strings.Contains(status, "Problem connecting to replica") {
			common.ReplicationStatusGauge.WithLabelValues(host).Set(0)
		} else if strings.Contains(status, "Can't acquire busy replica") {
			common.ReplicationStatusGauge.WithLabelValues(host).Set(1)
		} else {
			log.Warnf("Unknown replication status host: %s, status: %s", host, status)
			common.ReplicationStatusGauge.WithLabelValues(host).Set(0)
		}
	}

	return nil
}

/**
 * Query the LDAP server for a specific monitor attribute.
 */
func MonitorAttributeQuery(l *ldap.Conn, baseDN, attribute string) (float64, error) {
	req := ldap.NewSearchRequest(
		baseDN, ldap.ScopeBaseObject,
		ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)", []string{attribute}, nil,
	)
	sr, err := l.Search(req)
	if err != nil {
		return -1, err
	}

	if len(sr.Entries) == 0 {
		return -1, fmt.Errorf("no entries found for %s", baseDN)
	}

	val := sr.Entries[0].GetAttributeValue(attribute)
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, err
	}

	return num, nil
}
