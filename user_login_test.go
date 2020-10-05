package gofortiadc

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {

	tests := []struct {
		addr     string
		user     string
		password string
		valid    bool
	}{
		{"badaddr://feezf", "", "", false},
		{"http://fakehostname.bad", "", "", false},
		{os.Getenv("GOFORTI_ADDRESS"), "baduser", "badpassword", false},
		{os.Getenv("GOFORTI_ADDRESS"), os.Getenv("GOFORTI_USERNAME"), "badpassword", false},
		{os.Getenv("GOFORTI_ADDRESS"), os.Getenv("GOFORTI_USERNAME"), os.Getenv("GOFORTI_PASSWORD"), true},
	}

	for _, tc := range tests {
		cookieJar, _ := cookiejar.New(nil)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		httpClient := &http.Client{
			Jar:       cookieJar,
			Transport: tr,
		}

		fortiClient := Client{
			Client:   httpClient,
			Address:  tc.addr,
			Username: tc.user,
			Password: tc.password,
		}

		err := fortiClient.Login()
		if tc.valid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func NewClientHelper() (*Client, error) {
	cookieJar, _ := cookiejar.New(nil)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	fortiClient := &Client{
		Client:   httpClient,
		Address:  os.Getenv("GOFORTI_ADDRESS"),
		Username: os.Getenv("GOFORTI_USERNAME"),
		Password: os.Getenv("GOFORTI_PASSWORD"),
	}

	err := fortiClient.Login()
	if err != nil {
		return &Client{}, err
	}

	return fortiClient, nil
}
