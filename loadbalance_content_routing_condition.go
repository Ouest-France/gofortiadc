package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// LoadbalanceContentRoutingCondition represents a content routing condition request/response
type LoadbalanceContentRoutingCondition struct {
	Mkey    string `json:"mkey"`
	Object  string `json:"object"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Reverse string `json:"reverse"`
}

// LoadbalanceGetContentRoutingConditions returns the list of all content routing conditions
func (c *Client) LoadbalanceGetContentRoutingConditions(cr string) ([]LoadbalanceContentRoutingCondition, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?pkey=%s", c.Address, cr))
	if err != nil {
		return []LoadbalanceContentRoutingCondition{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return []LoadbalanceContentRoutingCondition{}, fmt.Errorf("failed to get content routing conditions list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return []LoadbalanceContentRoutingCondition{}, err
	}

	var LoadbalanceContentRoutingConditionPayload struct {
		Payload []LoadbalanceContentRoutingCondition
	}
	err = json.Unmarshal(body, &LoadbalanceContentRoutingConditionPayload)
	if err != nil {
		return []LoadbalanceContentRoutingCondition{}, err
	}

	return LoadbalanceContentRoutingConditionPayload.Payload, nil
}

// LoadbalanceGetContentRoutingCondition returns a content routing condition by name
func (c *Client) LoadbalanceGetContentRoutingCondition(cr, mkey string) (LoadbalanceContentRoutingCondition, error) {
	get, err := c.Client.Get(fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?pkey=%s", c.Address, cr))
	if err != nil {
		return LoadbalanceContentRoutingCondition{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return LoadbalanceContentRoutingCondition{}, fmt.Errorf("failed to get content routing conditions list with status code: %d", get.StatusCode)
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return LoadbalanceContentRoutingCondition{}, err
	}

	var LoadbalanceContentRoutingConditionPayload struct {
		Payload []LoadbalanceContentRoutingCondition
	}
	err = json.Unmarshal(body, &LoadbalanceContentRoutingConditionPayload)
	if err != nil {
		return LoadbalanceContentRoutingCondition{}, err
	}

	for _, rs := range LoadbalanceContentRoutingConditionPayload.Payload {
		if rs.Mkey == mkey {
			return rs, nil
		}
	}

	return LoadbalanceContentRoutingCondition{}, fmt.Errorf("content routing condition %s not found", mkey)
}

// LoadbalanceCreateContentRoutingCondition creates a new content routing condition
func (c *Client) LoadbalanceCreateContentRoutingCondition(cr string, req LoadbalanceContentRoutingCondition) error {

	payloadJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.Client.Post(
		fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?pkey=%s", c.Address, cr),
		"application/json",
		bytes.NewReader(payloadJSON),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("content routing condition creation failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("content routing condition creation failed: %s ", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceUpdateContentRoutingCondition updates an existing content routing condition
func (c *Client) LoadbalanceUpdateContentRoutingCondition(cr string, rs LoadbalanceContentRoutingCondition) error {

	payloadJSON, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition/%s?pkey=%s", c.Address, rs.Mkey, cr),
		bytes.NewReader(payloadJSON),
	)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("content routing condition update failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("content routing condition update failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}

// LoadbalanceDeleteContentRoutingCondition deletes an existing content routing
func (c *Client) LoadbalanceDeleteContentRoutingCondition(cr, mkey string) error {

	if len(mkey) == 0 {
		return errors.New("content routing condition mkey cannot be empty")
	}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition/%s?pkey=%s", c.Address, mkey, cr),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("content routing condition deletion failed with status code: %d", resp.StatusCode)
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
		return fmt.Errorf("content routing condition deletion failed: %s", getErrorMessage(res.Payload))
	}

	return nil
}
