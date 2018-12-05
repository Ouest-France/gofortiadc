package goforti

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// LoadbalanceRealServer represents a real server request/response
type LoadbalanceRealServer struct {
	Status   string `json:"status"`
	Address  string `json:"address"`
	Address6 string `json:"address6"`
	Mkey     string `json:"mkey"`
}

// LoadbalanceGetRealServers returns the list of all real servers
func (c *Client) LoadbalanceGetRealServers() ([]LoadbalanceRealServer, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_real_server", c.Address))
	if err != nil {
		return []LoadbalanceRealServer{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return []LoadbalanceRealServer{}, fmt.Errorf("failed to get real servers list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
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
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_real_server", c.Address))
	if err != nil {
		return LoadbalanceRealServer{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return LoadbalanceRealServer{}, fmt.Errorf("failed to get real servers list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
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

	return LoadbalanceRealServer{}, fmt.Errorf("real server %s not found", name)
}

// LoadbalanceCreateRealServer creates a new real server
func (c *Client) LoadbalanceCreateRealServer(req LoadbalanceRealServer) error {

	payloadJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.Client.Post(fmt.Sprintf("%s/api/load_balance_real_server", c.Address), "application/json", bytes.NewReader(payloadJSON))
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
		return fmt.Errorf("real server creation failed with result payload: %d", res.Payload)
	}

	return nil
}

// LoadbalanceUpdateRealServer updates an existing real server
func (c *Client) LoadbalanceUpdateRealServer(rs LoadbalanceRealServer) error {

	payloadJSON, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_real_server/%s", c.Address, rs.Mkey), bytes.NewReader(payloadJSON))
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
	fmt.Println(string(body))

	res := struct{ Payload int }{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if res.Payload != 0 {
		return fmt.Errorf("real server update failed with result payload: %d", res.Payload)
	}

	return nil
}

// LoadbalanceDeleteRealServer deletes an existing real server
func (c *Client) LoadbalanceDeleteRealServer(rs string) error {

	if len(rs) == 0 {
		return errors.New("real server name cannot be empty")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_real_server/%s", c.Address, rs), nil)
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
		return fmt.Errorf("real server deletion failed with result payload: %d", res.Payload)
	}

	return nil
}
