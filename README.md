# gofortiadc

A FortiADC API client enabling Go programs to interact with a Fortiadc server

[![GoDoc](https://godoc.org/github.com/Ouest-France/gofortiadc?status.svg)](https://godoc.org/github.com/Ouest-France/gofortiadc)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ouest-France/gofortiadc)](https://goreportcard.com/report/github.com/Ouest-France/gofortiadc)
[![GitHub issues](https://img.shields.io/github/issues/Ouest-France/gofortiadc.svg)](https://github.com/Ouest-France/gofortiadc/issues)

## NOTE

Tested with Fortiadc 5.1 API

## Example

```go
package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/Ouest-France/gofortiadc"
)

func main() {

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

	// Get the list of all virtual servers
	res, err := fortiClient.LoadbalanceGetVirtualServers()
	if err != nil {
		log.Print(err)
	}
}
```
