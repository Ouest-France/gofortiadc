package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// LoadbalancePoolMember represents a real server pool member request/response
type LoadbalancePoolMember struct {
	Mkey                     string `json:"mkey"`
	Address                  string `json:"address"`
	Address6                 string `json:"address6"`
	HealthCheckInherit       string `json:"health_check_inherit"`
	MHealthCheck             string `json:"m_health_check"`
	MHealthCheckRelationship string `json:"m_health_check_relationship"`
	MHealthCheckList         string `json:"m_health_check_list"`
	Port                     string `json:"port"`
	Weight                   string `json:"weight"`
	Connlimit                string `json:"connlimit"`
	Recover                  string `json:"recover"`
	Warmup                   string `json:"warmup"`
	Warmrate                 string `json:"warmrate"`
	ConnectionRateLimit      string `json:"connection-rate-limit"`
	Cookie                   string `json:"cookie"`
	Status                   string `json:"status"`
	Ssl                      string `json:"ssl"`
	RsProfile                string `json:"rs_profile"`
	RsProfileInherit         string `json:"rs_profile_inherit"`
	Backup                   string `json:"backup"`
	HcStatus                 string `json:"hc_status"`
	MysqlGroupID             string `json:"mysql_group_id"`
	MysqlReadOnly            string `json:"mysql_read_only"`
	RealServerID             string `json:"real_server_id"`
}

// LoadbalanceGetPoolMembers returns the list of all real server pool members
func (c *Client) LoadbalanceGetPoolMembers(pool string) ([]LoadbalancePoolMember, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_pool_child_pool_member?pkey=%s", c.Address, pool), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalancePoolMember{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalancePoolMember{}, fmt.Errorf("failed to get pool members list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []LoadbalancePoolMember{}, err
	}

	var LoadbalancePoolPayload struct {
		Payload []LoadbalancePoolMember
	}
	err = json.Unmarshal(body, &LoadbalancePoolPayload)
	if err != nil {
		return []LoadbalancePoolMember{}, err
	}

	return LoadbalancePoolPayload.Payload, nil
}

// LoadbalanceGetPoolMember returns a real server pool member by mkey
func (c *Client) LoadbalanceGetPoolMember(pool, mkey string) (LoadbalancePoolMember, error) {

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_pool_child_pool_member?pkey=%s", c.Address, pool), nil)
	if err != nil {
		return LoadbalancePoolMember{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalancePoolMember{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalancePoolMember{}, fmt.Errorf("failed to get pool member with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LoadbalancePoolMember{}, err
	}

	var LoadbalancePoolPayload struct {
		Payload []LoadbalancePoolMember
	}
	err = json.Unmarshal(body, &LoadbalancePoolPayload)
	if err != nil {
		return LoadbalancePoolMember{}, err
	}

	for _, member := range LoadbalancePoolPayload.Payload {
		if member.Mkey == mkey {
			return member, nil
		}
	}

	return LoadbalancePoolMember{}, fmt.Errorf("pool member %s in pool %s not found", mkey, pool)
}

// LoadbalanceGetPoolMemberID returns a real server pool member id by name
func (c *Client) LoadbalanceGetPoolMemberID(pool, name string) (string, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_pool_child_pool_member?pkey=%s", c.Address, pool), nil)
	if err != nil {
		return "", err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to get pool member with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var LoadbalancePoolPayload struct {
		Payload []LoadbalancePoolMember
	}
	err = json.Unmarshal(body, &LoadbalancePoolPayload)
	if err != nil {
		return "", err
	}

	for _, member := range LoadbalancePoolPayload.Payload {
		if member.RealServerID == name {
			return member.Mkey, nil
		}
	}

	return "", fmt.Errorf("pool member %s in pool %s not found", name, pool)
}

// LoadbalanceCreatePoolMember creates a new real server pool member
func (c *Client) LoadbalanceCreatePoolMember(pool string, member LoadbalancePoolMember) error {

	payloadJSON, err := json.Marshal(member)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/load_balance_pool_child_pool_member?pkey=%s", c.Address, pool), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("pool member creation failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("pool member creation failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceUpdatePoolMember updates an existing real server pool member
func (c *Client) LoadbalanceUpdatePoolMember(pool, mkey string, member LoadbalancePoolMember) error {

	payloadJSON, err := json.Marshal(member)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_pool_child_pool_member?mkey=%s&pkey=%s", c.Address, mkey, pool), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("pool member update failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("pool member update failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceDeletePoolMember deletes an existing real server pool member
func (c *Client) LoadbalanceDeletePoolMember(pool, mkey string) error {

	if len(pool) == 0 {
		return errors.New("pool member name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_pool_child_pool_member?mkey=%s&pkey=%s", c.Address, mkey, pool), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("pool member deletion failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
