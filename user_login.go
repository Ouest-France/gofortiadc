package goforti

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// AuthRes represents a login auth response
type AuthRes struct {
	Sid      int    `json:"sid"`
	Username string `json:"username"`
}

// Login authenticates goforti client
func (c *Client) Login() error {

	if len(c.Address) == 0 {
		return errors.New("FortiADC address cannot be empty")
	}

	payload := map[string]string{
		"username": c.Username,
		"password": c.Password,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := c.Client.Post(fmt.Sprintf("%s/api/user/login", c.Address), "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Login failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var authRes struct{ payload AuthRes }
	err = json.Unmarshal(body, &authRes)
	if err != nil {
		return err
	}

	return nil
}
