package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://dns.hetzner.com/api/v1"
const version = "0.0.1"
const userAgent = "hetzner-dns-go-client/" + version

type Client struct {
	AuthAPIToken string
	baseURL      string
	userAgent    string
}

func NewAuthApiTokenClient(authAPIToken string) HetznerDNSClient {
	return &Client{
		AuthAPIToken: authAPIToken,
		baseURL:      baseURL,
		userAgent:    userAgent,
	}
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) GetVersion() string {
	return version
}

func (c *Client) doGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Auth-API-Token", c.AuthAPIToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
