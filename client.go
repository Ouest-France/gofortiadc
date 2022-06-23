package gofortiadc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client represents a FortiADC API client instance
type Client struct {
	Client   *http.Client
	Address  string
	Username string
	Password string
	Token    string
	VDom     string
}

// NewRequest create an http.Request with authorization header set
func (c *Client) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
	uri, err := c.getUrl(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, uri, body)

	if err == nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	return req, err
}

func (c *Client) getUrl(path string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	q := u.Query()
	if c.VDom != "" {
		q.Set("vdom", c.VDom)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
