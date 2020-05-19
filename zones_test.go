package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	client "github.com/nl2go/hetzner-dns-go"
	. "gopkg.in/check.v1"
)

func (s *ClientSuite) TestZonesGetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/zones.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	hetznerDNSClient := client.NewAuthApiTokenClient("authToken123")
	hetznerDNSClient.SetBaseURL(ts.URL)

	servers, err := hetznerDNSClient.ZonesGet()
	c.Assert(err, IsNil)
	c.Assert(len(servers), Equals, 2)
	c.Assert(servers[0].Name, Equals, testDomain1)
	c.Assert(servers[1].Name, Equals, testDomain2)
}

func (s *ClientSuite) TestZonesGetInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	hetznerDNSClient := client.NewAuthApiTokenClient("authToken123")
	hetznerDNSClient.SetBaseURL(ts.URL)

	_, err := hetznerDNSClient.ZonesGet()
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestZonesGetServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	hetznerDNSClient := client.NewAuthApiTokenClient("authToken123")
	hetznerDNSClient.SetBaseURL(ts.URL)

	_, err := hetznerDNSClient.ZonesGet()
	c.Assert(err, Not(IsNil))
}
