package goforti

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
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
		Address:  os.Getenv("GOFORTI_ADDRESS"),
		Username: os.Getenv("GOFORTI_USERNAME"),
		Password: os.Getenv("GOFORTI_PASSWORD"),
	}

	err := fortiClient.Login()
	if err != nil {
		t.Fatal(err)
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
