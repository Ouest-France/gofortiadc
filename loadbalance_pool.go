package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_pool", c.Address), nil)
	if err != nil {
		return []LoadbalancePoolRes{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalancePoolRes{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalancePoolRes{}, fmt.Errorf("failed to get pools list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
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

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_pool", c.Address), nil)
	if err != nil {
		return LoadbalancePoolRes{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalancePoolRes{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalancePoolRes{}, fmt.Errorf("failed to get pool with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
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
func (c *Client) LoadbalanceCreatePool(pool LoadbalancePoolReq) error {

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
func (c *Client) LoadbalanceUpdatePool(pool LoadbalancePoolReq) error {

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
	fmt.Println(string(body))

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
