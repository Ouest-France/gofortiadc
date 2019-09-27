package gofortiadc

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func Example_login() {

	// Construct an http client with cookies
	cookieJar, _ := cookiejar.New(nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	httpClient := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	// Construct new forti Client instance
	fortiClient := Client{
		Client:   httpClient,
		Address:  "https://myfortiadc.example.com",
		Username: "fortiuser",
		Password: "fortipassword",
	}

	// Send auth request to get authorization token
	err := fortiClient.Login()
	if err != nil {
		log.Fatal(err)
	}
}
