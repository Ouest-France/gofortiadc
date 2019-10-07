package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// LoadbalanceVirtualServer represents a virtual server request/response
type LoadbalanceVirtualServer struct {
	AddrType             string `json:"addr-type"`
	Address              string `json:"address"`
	Address6             string `json:"address6"`
	AdfsPublishedService string `json:"adfs-published-service"`
	Alone                string `json:"alone"`
	AuthPolicy           string `json:"auth_policy"`
	AvProfile            string `json:"av-profile"`
	ClientSSLProfile     string `json:"client_ssl_profile"`
	ClonePool            string `json:"clone-pool"`
	CloneTrafficType     string `json:"clone-traffic-type"`
	Comments             string `json:"comments"`
	ConnectionLimit      string `json:"connection-limit"`
	ConnectionRateLimit  string `json:"connection-rate-limit"`
	ContentRewriting     string `json:"content-rewriting"`
	ContentRewritingList string `json:"content-rewriting-list"`
	ContentRouting       string `json:"content-routing"`
	ContentRoutingList   string `json:"content-routing-list"`
	ErrorMsg             string `json:"error-msg"`
	ErrorPage            string `json:"error-page"`
	Fortiview            string `json:"fortiview"`
	HTTP2HTTPS           string `json:"http2https"`
	HTTP2HTTPSPort       string `json:"http2https-port"`
	Interface            string `json:"interface"`
	L2ExceptionList      string `json:"l2-exception-list"`
	Method               string `json:"method"`
	Mkey                 string `json:"mkey"`
	PacketFwdMethod      string `json:"packet-fwd-method"`
	Pagespeed            string `json:"pagespeed"`
	Persistence          string `json:"persistence"`
	Pool                 string `json:"pool"`
	Port                 string `json:"port"`
	Profile              string `json:"profile"`
	Protocol             string `json:"protocol"`
	PublicIP             string `json:"public-ip"`
	PublicIPType         string `json:"public-ip-type"`
	PublicIP6            string `json:"public-ip6"`
	ScheduleList         string `json:"schedule-list"`
	SchedulePoolList     string `json:"schedule-pool-list"`
	ScriptingFlag        string `json:"scripting_flag"`
	ScriptingList        string `json:"scripting_list"`
	SourcePoolList       string `json:"source-pool-list"`
	SslMirror            string `json:"ssl-mirror"`
	SslMirrorIntf        string `json:"ssl-mirror-intf"`
	Status               string `json:"status"`
	TrafficGroup         string `json:"traffic-group"`
	TrafficLog           string `json:"traffic-log"`
	TransRateLimit       string `json:"trans-rate-limit"`
	Type                 string `json:"type"`
	WafProfile           string `json:"waf-profile"`
	Warmrate             string `json:"warmrate"`
	Warmup               string `json:"warmup"`
	Wccp                 string `json:"wccp"`
}

// LoadbalanceGetVirtualServers returns the list of all virtaul servers
func (c *Client) LoadbalanceGetVirtualServers() ([]LoadbalanceVirtualServer, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_virtual_server", c.Address), nil)
	if err != nil {
		return []LoadbalanceVirtualServer{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalanceVirtualServer{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalanceVirtualServer{}, fmt.Errorf("failed to get virtual servers list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []LoadbalanceVirtualServer{}, err
	}

	var loadbalanceVirtualServer struct {
		Payload []LoadbalanceVirtualServer
	}
	err = json.Unmarshal(body, &loadbalanceVirtualServer)
	if err != nil {
		return []LoadbalanceVirtualServer{}, err
	}

	return loadbalanceVirtualServer.Payload, nil
}

// LoadbalanceGetVirtualServer returns a virtual server by name
func (c *Client) LoadbalanceGetVirtualServer(name string) (LoadbalanceVirtualServer, error) {

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_virtual_server", c.Address), nil)
	if err != nil {
		return LoadbalanceVirtualServer{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalanceVirtualServer{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalanceVirtualServer{}, errors.New("Non 200 return code")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LoadbalanceVirtualServer{}, err
	}

	var loadbalanceVirtualServer struct {
		Payload []LoadbalanceVirtualServer
	}
	err = json.Unmarshal(body, &loadbalanceVirtualServer)
	if err != nil {
		return LoadbalanceVirtualServer{}, err
	}

	for _, lb := range loadbalanceVirtualServer.Payload {
		if lb.Mkey == name {
			return lb, nil
		}
	}

	return LoadbalanceVirtualServer{}, fmt.Errorf("virtual server %s not found", name)
}

// LoadbalanceCreateVirtualServer creates a new virtual server
func (c *Client) LoadbalanceCreateVirtualServer(vs LoadbalanceVirtualServer) error {

	payloadJSON, err := json.Marshal(vs)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/load_balance_virtual_server", c.Address), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("virtual server creation failed with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resJSON := struct{ Payload int }{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return err
	}

	if resJSON.Payload != 0 {
		return fmt.Errorf("virtual server creation failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceUpdateVirtualServer updates an existing virtual server
func (c *Client) LoadbalanceUpdateVirtualServer(vs LoadbalanceVirtualServer) error {

	payloadJSON, err := json.Marshal(vs)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_virtual_server?mkey=%s", c.Address, vs.Mkey), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("virtual server update failed with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	resJSON := struct{ Payload int }{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return err
	}

	if resJSON.Payload != 0 {
		return fmt.Errorf("virtual server update failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceDeleteVirtualServer deletes an existing virtual server
func (c *Client) LoadbalanceDeleteVirtualServer(vs string) error {

	if len(vs) == 0 {
		return errors.New("virtual server name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_virtual_server?mkey=%s", c.Address, vs), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("virtual server deletion failed with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resJSON := struct{ Payload int }{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return err
	}

	if resJSON.Payload != 0 {
		return fmt.Errorf("virtual server deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
