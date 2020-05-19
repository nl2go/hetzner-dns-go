package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	client "github.com/nl2go/hetzner-dns-go"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ClientSuite struct{}

var _ = Suite(&ClientSuite{})

const testDomain1 = "ops-nl2go-test.de"
const testDomain2 = "hetzner-dns-test.de"

func (s *ClientSuite) TestSetDefaultUserAgent(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUserAgent := r.Header.Get("User-Agent")
		hetznerDNSClient := client.NewAuthApiTokenClient("authToken123")
		c.Assert(reqUserAgent, Equals, fmt.Sprintf("hetzner-dns-go-client/%s", hetznerDNSClient.GetVersion()))

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

	_, err := hetznerDNSClient.ZonesGet()
	c.Assert(err, IsNil)
}

func (s *ClientSuite) TestSetCustomUserAgent(c *C) {
	expectedUserAgent := "hetzner-dns-client-testsuite/0.0.1"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUserAgent := r.Header.Get("User-Agent")
		c.Assert(reqUserAgent, Equals, expectedUserAgent)

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
	hetznerDNSClient.SetUserAgent(expectedUserAgent)
	hetznerDNSClient.SetBaseURL(ts.URL)

	_, err := hetznerDNSClient.ZonesGet()
	c.Assert(err, IsNil)
}

func (s *ClientSuite) TestGetInvalidURL(c *C) {
	hetznerDNSClient := client.NewAuthApiTokenClient("authToken123")
	hetznerDNSClient.SetBaseURL("http://Not a valid URL")

	_, err := hetznerDNSClient.ZonesGet()
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestGetNonExistentURL(c *C) {
	hetznerDNSClient := client.NewAuthApiTokenClient("authToken123")
	hetznerDNSClient.SetBaseURL("http://DoesNotExist.nl2go")

	_, err := hetznerDNSClient.ZonesGet()
	c.Assert(err, Not(IsNil))
}
