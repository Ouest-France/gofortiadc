package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?pkey=%s", c.Address, cr), nil)
	if err != nil {
		return []LoadbalanceContentRoutingCondition{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalanceContentRoutingCondition{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalanceContentRoutingCondition{}, fmt.Errorf("failed to get content routing conditions list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
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
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?pkey=%s", c.Address, cr), nil)
	if err != nil {
		return LoadbalanceContentRoutingCondition{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalanceContentRoutingCondition{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalanceContentRoutingCondition{}, fmt.Errorf("failed to get content routing conditions list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
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

	return LoadbalanceContentRoutingCondition{}, fmt.Errorf("content routing condition %s not found: %w", mkey, ErrNotFound)
}

// LoadbalanceGetContentRoutingConditionID returns a content routing condition ID by request
func (c *Client) LoadbalanceGetContentRoutingConditionID(cr string, obj LoadbalanceContentRoutingCondition) (string, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?pkey=%s", c.Address, cr), nil)
	if err != nil {
		return "", err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to get content routing conditions list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var LoadbalanceContentRoutingConditionPayload struct {
		Payload []LoadbalanceContentRoutingCondition
	}
	err = json.Unmarshal(body, &LoadbalanceContentRoutingConditionPayload)
	if err != nil {
		return "", err
	}

	for _, rs := range LoadbalanceContentRoutingConditionPayload.Payload {
		if rs.Content != obj.Content {
			continue
		}
		if rs.Object != obj.Object {
			continue
		}
		if rs.Type != obj.Type {
			continue
		}
		if rs.Reverse != obj.Reverse {
			continue
		}

		return rs.Mkey, nil
	}

	return "", fmt.Errorf("content routing condition ID %+v not found: %w", obj, ErrNotFound)
}

// LoadbalanceCreateContentRoutingCondition creates a new content routing condition
func (c *Client) LoadbalanceCreateContentRoutingCondition(cr string, rc LoadbalanceContentRoutingCondition) error {

	payloadJSON, err := json.Marshal(rc)
	if err != nil {
		return err
	}

	req, err := c.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?pkey=%s", c.Address, cr),
		bytes.NewReader(payloadJSON),
	)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("content routing condition creation failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content routing condition creation failed: %s ", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceUpdateContentRoutingCondition updates an existing content routing condition
func (c *Client) LoadbalanceUpdateContentRoutingCondition(cr string, rs LoadbalanceContentRoutingCondition) error {

	payloadJSON, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	req, err := c.NewRequest(
		"PUT",
		fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?mkey=%s&pkey=%s", c.Address, rs.Mkey, cr),
		bytes.NewReader(payloadJSON),
	)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("content routing condition update failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content routing condition update failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceDeleteContentRoutingCondition deletes an existing content routing
func (c *Client) LoadbalanceDeleteContentRoutingCondition(cr, mkey string) error {

	if len(mkey) == 0 {
		return errors.New("content routing condition mkey cannot be empty")
	}

	req, err := c.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/api/load_balance_content_routing_child_match_condition?mkey=%s&pkey=%s", c.Address, mkey, cr),
		nil,
	)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("content routing condition deletion failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content routing condition deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
