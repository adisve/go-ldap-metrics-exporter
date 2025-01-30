package common

import "strings"

func IpaDomainToBaseDN(domain string) string {
	return "dc=" + strings.Replace(domain, ".", ",dc=", -1)
}

// function to use Ldap.User (for example "admin")
// and for example generate string "cn=admin,dc=example,dc=com"
func UserWithBaseDN(user string, domain string) string {
	return "cn=" + user + "," + IpaDomainToBaseDN(domain)
}