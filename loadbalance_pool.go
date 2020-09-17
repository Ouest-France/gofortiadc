package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// LoadbalancePool represents a real server pool request/response
type LoadbalancePool struct {
	HealthCheck             string                  `json:"health_check"`
	HealthCheckList         string                  `json:"health_check_list"`
	HealthCheckRelationship string                  `json:"health_check_relationship"`
	Mkey                    string                  `json:"mkey"`
	PoolMember              []LoadbalancePoolMember `json:"pool_member,omitempty"`
	PoolType                string                  `json:"pool_type"`
	RsProfile               string                  `json:"rs_profile"`
}

// LoadbalanceGetPools returns the list of all real server pools
func (c *Client) LoadbalanceGetPools() ([]LoadbalancePool, error) {

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_pool", c.Address), nil)
	if err != nil {
		return []LoadbalancePool{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalancePool{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalancePool{}, fmt.Errorf("failed to get pools list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []LoadbalancePool{}, err
	}

	var LoadbalancePoolPayload struct {
		Payload []LoadbalancePool
	}
	err = json.Unmarshal(body, &LoadbalancePoolPayload)
	if err != nil {
		return []LoadbalancePool{}, err
	}

	return LoadbalancePoolPayload.Payload, nil
}

// LoadbalanceGetPool returns a real server pool by name
func (c *Client) LoadbalanceGetPool(name string) (LoadbalancePool, error) {

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_pool", c.Address), nil)
	if err != nil {
		return LoadbalancePool{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalancePool{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalancePool{}, fmt.Errorf("failed to get pool with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LoadbalancePool{}, err
	}

	var LoadbalancePoolPayload struct {
		Payload []LoadbalancePool
	}
	err = json.Unmarshal(body, &LoadbalancePoolPayload)
	if err != nil {
		return LoadbalancePool{}, err
	}

	for _, pool := range LoadbalancePoolPayload.Payload {
		if pool.Mkey == name {
			return pool, nil
		}
	}

	return LoadbalancePool{}, fmt.Errorf("pool %s not found", name)
}

// LoadbalanceCreatePool creates a new real server pool
func (c *Client) LoadbalanceCreatePool(pool LoadbalancePool) error {

	payloadJSON, err := json.Marshal(pool)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/load_balance_pool", c.Address), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("virtual server pool creation failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("virtual server pool creation failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceUpdatePool updates an existing real server pool
func (c *Client) LoadbalanceUpdatePool(pool LoadbalancePool) error {

	payloadJSON, err := json.Marshal(pool)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_pool?mkey=%s", c.Address, pool.Mkey), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("pool update failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("pool update failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceDeletePool deletes an existing real server pool
func (c *Client) LoadbalanceDeletePool(pool string) error {

	if len(pool) == 0 {
		return errors.New("pool name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_pool?mkey=%s", c.Address, pool), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("deletion failed with status code: %d", res.StatusCode)
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
