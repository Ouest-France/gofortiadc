package gofortiadc

import (
	"net/http"
)

// Client represents a Forti API client instance
type Client struct {
	Client   *http.Client
	Address  string
	Username string
	Password string
}
