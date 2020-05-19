package models

type ZonesResponse struct {
	Meta  Meta   `json:"meta"`
	Zones []Zone `json:"zones"`
}

type Zone struct {
	Created         string   `json:"created"`
	ID              string   `json:"id"`
	IsSecondaryDNS  bool     `json:"is_secondary_dns"`
	LegacyDNSHost   string   `json:"legacy_dns_host"`
	LegacyNs        []string `json:"legacy_ns"`
	Modified        string   `json:"modified"`
	Name            string   `json:"name"`
	Ns              []string `json:"ns"`
	Owner           string   `json:"owner"`
	Paused          bool     `json:"paused"`
	Permission      string   `json:"permission"`
	Project         string   `json:"project"`
	RecordsCount    int      `json:"records_count"`
	Registrar       string   `json:"registrar"`
	Status          string   `json:"status"`
	TTL             int      `json:"ttl"`
	TxtVerification struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	} `json:"txt_verification"`
	Verified string `json:"verified"`
	ZoneType struct {
		Description string      `json:"description"`
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		Prices      interface{} `json:"prices"`
	} `json:"zone_type"`
}
