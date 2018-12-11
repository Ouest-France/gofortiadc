package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// LoadbalanceContentRouting represents a content routing request/response
type LoadbalanceContentRouting struct {
	Mkey                  string `json:"mkey"`
	Type                  string `json:"type"`
	PacketFwdMethod       string `json:"packet-fwd-method"`
	SourcePoolList        string `json:"source-pool-list"`
	Persistence           string `json:"persistence"`
	PersistenceInherit    string `json:"persistence_inherit"`
	Method                string `json:"method"`
	MethodInherit         string `json:"method_inherit"`
	ConnectionPool        string `json:"connection-pool"`
	ConnectionPoolInherit string `json:"connection_pool_inherit"`
	Pool                  string `json:"pool"`
	IP                    string `json:"ip"`
	IP6                   string `json:"ip6"`
	Comments              string `json:"comments"`
	ScheduleList          string `json:"schedule-list"`
	SchedulePoolList      string `json:"schedule-pool-list"`
}

// LoadbalanceGetContentRoutings returns the list of all content routings
func (c *Client) LoadbalanceGetContentRoutings() ([]LoadbalanceContentRouting, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_content_routing", c.Address))
	if err != nil {
		return []LoadbalanceContentRouting{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return []LoadbalanceContentRouting{}, fmt.Errorf("failed to get content routing list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return []LoadbalanceContentRouting{}, err
	}

	var LoadbalanceContentRoutingPayload struct {
		Payload []LoadbalanceContentRouting
	}
	err = json.Unmarshal(body, &LoadbalanceContentRoutingPayload)
	if err != nil {
		return []LoadbalanceContentRouting{}, err
	}

	return LoadbalanceContentRoutingPayload.Payload, nil
}

// LoadbalanceGetContentRouting returns a content routing by name
func (c *Client) LoadbalanceGetContentRouting(name string) (LoadbalanceContentRouting, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_content_routing", c.Address))
	if err != nil {
		return LoadbalanceContentRouting{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return LoadbalanceContentRouting{}, fmt.Errorf("failed to get content routing list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return LoadbalanceContentRouting{}, err
	}

	var LoadbalanceContentRoutingPayload struct {
		Payload []LoadbalanceContentRouting
	}
	err = json.Unmarshal(body, &LoadbalanceContentRoutingPayload)
	if err != nil {
		return LoadbalanceContentRouting{}, err
	}

	for _, rs := range LoadbalanceContentRoutingPayload.Payload {
		if rs.Mkey == name {
			return rs, nil
		}
	}

	return LoadbalanceContentRouting{}, fmt.Errorf("content routing %s not found", name)
}

// LoadbalanceCreateContentRouting creates a new content routing
func (c *Client) LoadbalanceCreateContentRouting(req LoadbalanceContentRouting) error {

	payloadJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.Client.Post(fmt.Sprintf("%s/api/load_balance_content_routing", c.Address), "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("content routing creation failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("content routing creation failed: %s ", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceUpdateContentRouting updates an existing content routing
func (c *Client) LoadbalanceUpdateContentRouting(rs LoadbalanceContentRouting) error {

	payloadJSON, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_content_routing/%s", c.Address, rs.Mkey), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("content routing update failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("content routing update failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceDeleteContentRouting deletes an existing content routing
func (c *Client) LoadbalanceDeleteContentRouting(name string) error {

	if len(name) == 0 {
		return errors.New("content routing name cannot be empty")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_content_routing/%s", c.Address, name), nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("content routing deletion failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("content routing deletion failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}
