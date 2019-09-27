package gofortiadc

import (
	"fmt"
	"io"
	"net/http"
)

// Client represents a FortiADC API client instance
type Client struct {
	Client   *http.Client
	Address  string
	Username string
	Password string
	Token    string
}

// NewRequest create an http.Request with authorization header set
func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, url, body)

	if err == nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	return req, err
}
