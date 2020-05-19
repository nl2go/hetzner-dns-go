package client

import "github.com/nl2go/hetzner-dns-go/models"

type HetznerDNSClient interface {
	SetBaseURL(baseURL string)
	SetUserAgent(userAgent string)
	GetVersion() string
	ZonesGet() ([]models.Zone, error)
}
