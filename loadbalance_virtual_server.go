package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// LoadbalanceVirtualServerReq represents a virtual server request
type LoadbalanceVirtualServerReq struct {
	Status               string `json:"status"`
	Type                 string `json:"type"`
	AddrType             string `json:"addr-type"`
	Address              string `json:"address"`
	Address6             string `json:"address6"`
	PacketFwdMethod      string `json:"packet-fwd-method"`
	Port                 string `json:"port"`
	PortRange            string `json:"port-range"`
	ConnectionLimit      string `json:"connection-limit"`
	ContentRouting       string `json:"content-routing"`
	ContentRewriting     string `json:"content-rewriting"`
	ErrorMsg             string `json:"error-msg"`
	Warmup               string `json:"warmup"`
	Warmrate             string `json:"warmrate"`
	ConnectionRateLimit  string `json:"connection-rate-limit"`
	Log                  string `json:"log"`
	Alone                string `json:"alone"`
	TransRateLimit       string `json:"trans-rate-limit"`
	Mkey                 string `json:"mkey"`
	Interface            string `json:"interface"`
	ContentRoutingList   string `json:"content-routing-list"`
	ContentRewritingList string `json:"content-rewriting-list"`
	Profile              string `json:"profile"`
	Persistence          string `json:"persistence"`
	Method               string `json:"method"`
	Pool                 string `json:"pool"`
	SrcPool              string `json:"src_pool"`
	ErrorPage            string `json:"error-page"`
	WafProfile           string `json:"waf-profile"`
	AuthPolicy           string `json:"auth_policy"`
	Scripting            string `json:"scripting"`
}

// LoadbalanceVirtualServerRes represents a virtual server response
type LoadbalanceVirtualServerRes struct {
	Mkey                 string `json:"mkey"`
	Status               string `json:"status"`
	Type                 string `json:"type"`
	Interface            string `json:"interface"`
	AddrType             string `json:"addr-type"`
	Address              string `json:"address"`
	Address6             string `json:"address6"`
	PacketFwdMethod      string `json:"packet-fwd-method"`
	Port                 string `json:"port"`
	PortRange            string `json:"port-range"`
	ConnectionLimit      string `json:"connection-limit"`
	ContentRouting       string `json:"content-routing"`
	ContentRoutingList   string `json:"content-routing-list"`
	ContentRewriting     string `json:"content-rewriting"`
	ContentRewritingList string `json:"content-rewriting-list"`
	Profile              string `json:"profile"`
	Persistence          string `json:"persistence"`
	Method               string `json:"method"`
	ConnectionPool       string `json:"connection-pool"`
	Pool                 string `json:"pool"`
	SrcPool              string `json:"src_pool"`
	ErrorPage            string `json:"error-page"`
	ErrorMsg             string `json:"error-msg"`
	Warmup               string `json:"warmup"`
	Warmrate             string `json:"warmrate"`
	ConnectionRateLimit  string `json:"connection-rate-limit"`
	Log                  string `json:"log"`
	Alone                string `json:"alone"`
	TransRateLimit       string `json:"trans-rate-limit"`
	WafProfile           string `json:"waf-profile"`
	AuthPolicy           string `json:"auth_policy"`
	Scripting            string `json:"scripting"`
	Nondeletable         int    `json:"_nondeletable"`
	Noneditable          int    `json:"_noneditable"`
	CurrentStatus        int    `json:"current-status"`
}

// LoadbalanceGetVirtualServers returns the list of all virtaul servers
func (c *Client) LoadbalanceGetVirtualServers() ([]LoadbalanceVirtualServerRes, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_virtual_server", c.Address))
	if err != nil {
		return []LoadbalanceVirtualServerRes{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return []LoadbalanceVirtualServerRes{}, fmt.Errorf("failed to get virtual servers list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return []LoadbalanceVirtualServerRes{}, err
	}

	var loadbalanceVirtualServerRes struct {
		Payload []LoadbalanceVirtualServerRes
	}
	err = json.Unmarshal(body, &loadbalanceVirtualServerRes)
	if err != nil {
		return []LoadbalanceVirtualServerRes{}, err
	}

	return loadbalanceVirtualServerRes.Payload, nil
}

// LoadbalanceGetVirtualServer returns a virtual server by name
func (c *Client) LoadbalanceGetVirtualServer(name string) (LoadbalanceVirtualServerRes, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_virtual_server", c.Address))
	if err != nil {
		return LoadbalanceVirtualServerRes{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return LoadbalanceVirtualServerRes{}, errors.New("Non 200 return code")
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return LoadbalanceVirtualServerRes{}, err
	}

	var loadbalanceVirtualServerRes struct {
		Payload []LoadbalanceVirtualServerRes
	}
	err = json.Unmarshal(body, &loadbalanceVirtualServerRes)
	if err != nil {
		return LoadbalanceVirtualServerRes{}, err
	}

	for _, lb := range loadbalanceVirtualServerRes.Payload {
		if lb.Mkey == name {
			return lb, nil
		}
	}

	return LoadbalanceVirtualServerRes{}, fmt.Errorf("virtual server %s not found", name)
}

// LoadbalanceCreateVirtualServer creates a new virtual server
func (c *Client) LoadbalanceCreateVirtualServer(req LoadbalanceVirtualServerReq) error {

	payloadJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.Client.Post(fmt.Sprintf("%s/api/load_balance_virtual_server", c.Address), "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("virtual server creation failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("virtual server creation failed with result payload: %d", res.Payload)
	}

	return nil
}

// LoadbalanceUpdateVirtualServer updates an existing virtual server
func (c *Client) LoadbalanceUpdateVirtualServer(vs LoadbalanceVirtualServerReq) error {

	payloadJSON, err := json.Marshal(vs)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_virtual_server/%s", c.Address, vs.Mkey), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("virtual server update failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("virtual server update failed with result payload: %d", res.Payload)
	}

	return nil
}

// LoadbalanceDeleteVirtualServer deletes an existing virtual server
func (c *Client) LoadbalanceDeleteVirtualServer(vs string) error {

	if len(vs) == 0 {
		return errors.New("virtual server name cannot be empty")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_virtual_server/%s", c.Address, vs), nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("virtual server deletion failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("virtual server deletion failed with result payload: %d", res.Payload)
	}

	return nil
}
