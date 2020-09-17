package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// LoadbalanceContentRewritingCondition represents a content rewriting condition request/response
type LoadbalanceContentRewritingCondition struct {
	Mkey       string `json:"mkey"`
	Content    string `json:"content"`
	Ignorecase string `json:"ignorecase"`
	Object     string `json:"object"`
	Reverse    string `json:"reverse"`
	Type       string `json:"type"`
}

// LoadbalanceGetContentRewritingConditions returns the list of all content rewriting conditions
func (c *Client) LoadbalanceGetContentRewritingConditions(cr string) ([]LoadbalanceContentRewritingCondition, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_rewriting_child_match_condition?pkey=%s", c.Address, cr), nil)
	if err != nil {
		return []LoadbalanceContentRewritingCondition{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalanceContentRewritingCondition{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalanceContentRewritingCondition{}, fmt.Errorf("failed to get content rewriting conditions list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []LoadbalanceContentRewritingCondition{}, err
	}

	var LoadbalanceContentRewritingConditionPayload struct {
		Payload []LoadbalanceContentRewritingCondition
	}
	err = json.Unmarshal(body, &LoadbalanceContentRewritingConditionPayload)
	if err != nil {
		return []LoadbalanceContentRewritingCondition{}, err
	}

	return LoadbalanceContentRewritingConditionPayload.Payload, nil
}

// LoadbalanceGetContentRewritingCondition returns a content rewriting condition by name
func (c *Client) LoadbalanceGetContentRewritingCondition(cr, mkey string) (LoadbalanceContentRewritingCondition, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_rewriting_child_match_condition?pkey=%s", c.Address, cr), nil)
	if err != nil {
		return LoadbalanceContentRewritingCondition{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalanceContentRewritingCondition{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalanceContentRewritingCondition{}, fmt.Errorf("failed to get content rewriting conditions list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LoadbalanceContentRewritingCondition{}, err
	}

	var LoadbalanceContentRewritingConditionPayload struct {
		Payload []LoadbalanceContentRewritingCondition
	}
	err = json.Unmarshal(body, &LoadbalanceContentRewritingConditionPayload)
	if err != nil {
		return LoadbalanceContentRewritingCondition{}, err
	}

	for _, cr := range LoadbalanceContentRewritingConditionPayload.Payload {
		if cr.Mkey == mkey {
			return cr, nil
		}
	}

	return LoadbalanceContentRewritingCondition{}, fmt.Errorf("content rewriting condition %s not found", mkey)
}

// LoadbalanceGetContentRewritingConditionID returns a content rewriting condition ID by request
func (c *Client) LoadbalanceGetContentRewritingConditionID(cr string, obj LoadbalanceContentRewritingCondition) (string, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_rewriting_child_match_condition?pkey=%s", c.Address, cr), nil)
	if err != nil {
		return "", err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to get content rewriting conditions list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var LoadbalanceContentRewritingConditionPayload struct {
		Payload []LoadbalanceContentRewritingCondition
	}
	err = json.Unmarshal(body, &LoadbalanceContentRewritingConditionPayload)
	if err != nil {
		return "", err
	}

	for _, cr := range LoadbalanceContentRewritingConditionPayload.Payload {
		if cr.Content != obj.Content {
			continue
		}
		if cr.Ignorecase != obj.Ignorecase {
			continue
		}
		if cr.Object != obj.Object {
			continue
		}
		if cr.Reverse != obj.Reverse {
			continue
		}
		if cr.Type != obj.Type {
			continue
		}

		return cr.Mkey, nil
	}

	return "", fmt.Errorf("content rewriting condition ID %+v not found", obj)
}

// LoadbalanceCreateContentRewritingCondition creates a new content rewriting condition
func (c *Client) LoadbalanceCreateContentRewritingCondition(cr string, rc LoadbalanceContentRewritingCondition) error {

	payloadJSON, err := json.Marshal(rc)
	if err != nil {
		return err
	}

	req, err := c.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/load_balance_content_rewriting_child_match_condition?pkey=%s", c.Address, cr),
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
		return fmt.Errorf("content rewriting condition creation failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content rewriting condition creation failed: %s ", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceUpdateContentRewritingCondition updates an existing content rewriting condition
func (c *Client) LoadbalanceUpdateContentRewritingCondition(cr string, rd LoadbalanceContentRewritingCondition) error {

	payloadJSON, err := json.Marshal(rd)
	if err != nil {
		return err
	}

	req, err := c.NewRequest(
		"PUT",
		fmt.Sprintf("%s/api/load_balance_content_rewriting_child_match_condition?mkey=%s&pkey=%s", c.Address, rd.Mkey, cr),
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
		return fmt.Errorf("content rewriting condition update failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content rewriting condition update failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceDeleteContentRewritingCondition deletes an existing content rewriting
func (c *Client) LoadbalanceDeleteContentRewritingCondition(cr, mkey string) error {

	if len(mkey) == 0 {
		return errors.New("content rewriting condition mkey cannot be empty")
	}

	req, err := c.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/api/load_balance_content_rewriting_child_match_condition?mkey=%s&pkey=%s", c.Address, mkey, cr),
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
		return fmt.Errorf("content rewriting condition deletion failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content rewriting condition deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
