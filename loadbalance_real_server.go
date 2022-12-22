package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// LoadbalanceRealServer represents a real server request/response
type LoadbalanceRealServer struct {
	Type     string `json:"type"`
	Status   string `json:"status"`
	FQDN     string `json:"FQDN"`
	Address  string `json:"address"`
	Address6 string `json:"address6"`
	Mkey     string `json:"mkey"`
}

// LoadbalanceGetRealServers returns the list of all real servers
func (c *Client) LoadbalanceGetRealServers() ([]LoadbalanceRealServer, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_real_server", c.Address), nil)
	if err != nil {
		return []LoadbalanceRealServer{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalanceRealServer{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalanceRealServer{}, fmt.Errorf("failed to get real servers list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []LoadbalanceRealServer{}, err
	}

	var LoadbalanceRealServerPayload struct {
		Payload []LoadbalanceRealServer
	}
	err = json.Unmarshal(body, &LoadbalanceRealServerPayload)
	if err != nil {
		return []LoadbalanceRealServer{}, err
	}

	return LoadbalanceRealServerPayload.Payload, nil
}

// LoadbalanceGetRealServer returns a real server by name
func (c *Client) LoadbalanceGetRealServer(name string) (LoadbalanceRealServer, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_real_server", c.Address), nil)
	if err != nil {
		return LoadbalanceRealServer{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalanceRealServer{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalanceRealServer{}, fmt.Errorf("failed to get real servers list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LoadbalanceRealServer{}, err
	}

	var LoadbalanceRealServerPayload struct {
		Payload []LoadbalanceRealServer
	}
	err = json.Unmarshal(body, &LoadbalanceRealServerPayload)
	if err != nil {
		return LoadbalanceRealServer{}, err
	}

	for _, rs := range LoadbalanceRealServerPayload.Payload {
		if rs.Mkey == name {
			return rs, nil
		}
	}

	return LoadbalanceRealServer{}, fmt.Errorf("real server %s not found: %w", name, ErrNotFound)
}

// LoadbalanceCreateRealServer creates a new real server
func (c *Client) LoadbalanceCreateRealServer(rs LoadbalanceRealServer) error {

	payloadJSON, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/load_balance_real_server", c.Address), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("real server creation failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := struct{ Payload int }{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if res.Payload != 0 {
		return fmt.Errorf("real server creation failed: %s ", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceUpdateRealServer updates an existing real server
func (c *Client) LoadbalanceUpdateRealServer(rs LoadbalanceRealServer) error {

	payloadJSON, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_real_server?mkey=%s", c.Address, rs.Mkey), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("real server update failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := struct{ Payload int }{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if res.Payload != 0 {
		return fmt.Errorf("real server update failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceDeleteRealServer deletes an existing real server
func (c *Client) LoadbalanceDeleteRealServer(name string) error {

	if len(name) == 0 {
		return errors.New("real server name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_real_server?mkey=%s", c.Address, name), nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("real server deletion failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := struct{ Payload int }{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if res.Payload != 0 {
		return fmt.Errorf("real server deletion failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}
