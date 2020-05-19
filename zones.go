package client

import (
	"encoding/json"

	"github.com/nl2go/hetzner-dns-go/models"
)

func (c *Client) ZonesGet() ([]models.Zone, error) {
	url := c.baseURL + "/server"
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var zonesResp models.ZonesResponse
	err = json.Unmarshal(bytes, &zonesResp)
	if err != nil {
		return nil, err
	}

	var data []models.Zone
	for _, zone := range zonesResp.Zones {
		data = append(data, zone)
	}

	return data, nil
}
