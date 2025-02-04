package structs

type Config struct {
	LDAP struct {
		Address  string `json:"ldap.address"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"ldap"`
	IPA struct {
		Domain string `json:"domain"`
	} `json:"ipa"`
	Scrape struct {
		Interval int `json:"interval"`
	} `json:"scrape"`
	Server struct {
		Active  bool   `json:"active"`
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"server"`
	Log struct {
		Level string `json:"level"`
		JSON  bool   `json:"json"`
	} `json:"log"`
	Export struct {
		File     string `json:"file"`
		Interval int    `json:"interval"`
	} `json:"export"`
}
