package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// AuthRes represents a login auth response
type AuthRes struct {
	Token string `json:"token"`
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
		return fmt.Errorf("Login failed with http code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var authRes AuthRes
	err = json.Unmarshal(body, &authRes)
	if err != nil {
		return err
	}

	c.Token = authRes.Token

	return nil
}
