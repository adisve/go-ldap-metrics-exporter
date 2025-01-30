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

func NewConfig(
	ldapAddr string,
	ldapUser string,
	ldapPass string,
	ipaDomain string,
	scrapeInterval int,
	serverActive bool,
	serverAddr string,
	serverPort string,
	logLevel string,
	logJSON bool,
	exportFile string,
	exportInterval int,
) *Config {
	return &Config{
		LDAP: struct {
			Address  string `json:"ldap.address"`
			User     string `json:"user"`
			Password string `json:"password"`
		}{
			Address:  ldapAddr,
			User:     ldapUser,
			Password: ldapPass,
		},
		IPA: struct {
			Domain string `json:"domain"`
		}{
			Domain: ipaDomain,
		},
		Scrape: struct {
			Interval int `json:"interval"`
		}{
			Interval: scrapeInterval,
		},
		Server: struct {
			Active  bool   `json:"active"`
			Address string `json:"address"`
			Port    string `json:"port"`
		}{
			Active:  serverActive,
			Address: serverAddr,
			Port:    serverPort,
		},
		Log: struct {
			Level string `json:"level"`
			JSON  bool   `json:"json"`
		}{
			Level: logLevel,
			JSON:  logJSON,
		},
		Export: struct {
			File     string `json:"file"`
			Interval int    `json:"interval"`
		}{
			File:     exportFile,
			Interval: exportInterval,
		},
	}
}
