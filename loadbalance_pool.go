package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// LoadbalancePoolReq represents a real server pool request
type LoadbalancePoolReq struct {
	Mkey                    string `json:"mkey"`
	PoolType                string `json:"pool_type"`
	HealthCheck             string `json:"health_check"`
	HealthCheckRelationship string `json:"health_check_relationship"`
	HealthCheckList         string `json:"health_check_list"`
	RsProfile               string `json:"rs_profile"`
}

// LoadbalancePoolRes represents a real server pool response
type LoadbalancePoolRes struct {
	Mkey                    string                  `json:"mkey"`
	PoolType                string                  `json:"pool_type"`
	HealthCheck             string                  `json:"health_check"`
	HealthCheckRelationship string                  `json:"health_check_relationship"`
	HealthCheckList         string                  `json:"health_check_list"`
	RsProfile               string                  `json:"rs_profile"`
	PoolMember              []LoadbalancePoolMember `json:"pool_member"`
}

// LoadbalanceGetPools returns the list of all real server pools
func (c *Client) LoadbalanceGetPools() ([]LoadbalancePoolRes, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_pool", c.Address))
	if err != nil {
		return []LoadbalancePoolRes{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return []LoadbalancePoolRes{}, fmt.Errorf("failed to get pools list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return []LoadbalancePoolRes{}, err
	}

	var LoadbalancePoolPayload struct {
		Payload []LoadbalancePoolRes
	}
	err = json.Unmarshal(body, &LoadbalancePoolPayload)
	if err != nil {
		return []LoadbalancePoolRes{}, err
	}

	return LoadbalancePoolPayload.Payload, nil
}

// LoadbalanceGetPool returns a real server pool by name
func (c *Client) LoadbalanceGetPool(name string) (LoadbalancePoolRes, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_pool", c.Address))
	if err != nil {
		return LoadbalancePoolRes{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return LoadbalancePoolRes{}, fmt.Errorf("failed to get pool with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return LoadbalancePoolRes{}, err
	}

	var LoadbalancePoolPayload struct {
		Payload []LoadbalancePoolRes
	}
	err = json.Unmarshal(body, &LoadbalancePoolPayload)
	if err != nil {
		return LoadbalancePoolRes{}, err
	}

	for _, pool := range LoadbalancePoolPayload.Payload {
		if pool.Mkey == name {
			return pool, nil
		}
	}

	return LoadbalancePoolRes{}, fmt.Errorf("pool %s not found", name)
}

// LoadbalanceCreatePool creates a new real server pool
func (c *Client) LoadbalanceCreatePool(req LoadbalancePoolReq) error {

	payloadJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.Client.Post(fmt.Sprintf("%s/api/load_balance_pool", c.Address), "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("virtual server pool creation failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("virtual server pool creation failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceUpdatePool updates an existing real server pool
func (c *Client) LoadbalanceUpdatePool(pool LoadbalancePoolReq) error {

	payloadJSON, err := json.Marshal(pool)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_pool/%s", c.Address, pool.Mkey), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("pool update failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("pool update failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceDeletePool deletes an existing real server pool
func (c *Client) LoadbalanceDeletePool(pool string) error {

	if len(pool) == 0 {
		return errors.New("pool name cannot be empty")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_pool/%s", c.Address, pool), nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("deletion failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("deletion failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}
