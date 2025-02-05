package structs

type Config struct {
	LDAP struct {
		Address    string `json:"address"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		BaseDN     string `json:"baseDn"`
		UserBaseDN string `json:"userBaseDn"`
		Port       string `json:"port"`
	} `json:"ldap"`
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
