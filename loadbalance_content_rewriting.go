package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// LoadbalanceContentRewriting represents a content Rewriting request/response
type LoadbalanceContentRewriting struct {
	ActionType     string `json:"action_type"`
	URLStatus      string `json:"url_status"`
	URLContent     string `json:"url_content"`
	RefererStatus  string `json:"referer_status"`
	RefererContent string `json:"referer_content"`
	Redirect       string `json:"redirect"`
	Location       string `json:"location"`
	HeaderName     string `json:"header_name"`
	Comments       string `json:"comments"`
	Mkey           string `json:"mkey"`
	Action         string `json:"action"`
	HostStatus     string `json:"host_status"`
	HostContent    string `json:"host_content"`
}

// LoadbalanceGetContentRewritings returns the list of all content rewritings
func (c *Client) LoadbalanceGetContentRewritings() ([]LoadbalanceContentRewriting, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_rewriting", c.Address), nil)
	if err != nil {
		return []LoadbalanceContentRewriting{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []LoadbalanceContentRewriting{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []LoadbalanceContentRewriting{}, fmt.Errorf("failed to get content rewriting list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []LoadbalanceContentRewriting{}, err
	}

	var LoadbalanceContentRewritingPayload struct {
		Payload []LoadbalanceContentRewriting
	}
	err = json.Unmarshal(body, &LoadbalanceContentRewritingPayload)
	if err != nil {
		return []LoadbalanceContentRewriting{}, err
	}

	return LoadbalanceContentRewritingPayload.Payload, nil
}

// LoadbalanceGetContentRewriting returns a content rewriting by name
func (c *Client) LoadbalanceGetContentRewriting(name string) (LoadbalanceContentRewriting, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/load_balance_content_rewriting", c.Address), nil)
	if err != nil {
		return LoadbalanceContentRewriting{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return LoadbalanceContentRewriting{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LoadbalanceContentRewriting{}, fmt.Errorf("failed to get content rewriting list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LoadbalanceContentRewriting{}, err
	}

	var LoadbalanceContentRewritingPayload struct {
		Payload []LoadbalanceContentRewriting
	}
	err = json.Unmarshal(body, &LoadbalanceContentRewritingPayload)
	if err != nil {
		return LoadbalanceContentRewriting{}, err
	}

	for _, rs := range LoadbalanceContentRewritingPayload.Payload {
		if rs.Mkey == name {
			return rs, nil
		}
	}

	return LoadbalanceContentRewriting{}, fmt.Errorf("content rewriting %s not found", name)
}

// LoadbalanceCreateContentRewriting creates a new content rewriting
func (c *Client) LoadbalanceCreateContentRewriting(cr LoadbalanceContentRewriting) error {

	payloadJSON, err := json.Marshal(cr)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/load_balance_content_rewriting", c.Address), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("content rewriting creation failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content rewriting creation failed: %s ", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceUpdateContentRewriting updates an existing content rewriting
func (c *Client) LoadbalanceUpdateContentRewriting(cr LoadbalanceContentRewriting) error {

	payloadJSON, err := json.Marshal(cr)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s/api/load_balance_content_rewriting?mkey=%s", c.Address, cr.Mkey), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("content rewriting update failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content rewriting update failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// LoadbalanceDeleteContentRewriting deletes an existing content rewriting
func (c *Client) LoadbalanceDeleteContentRewriting(name string) error {

	if len(name) == 0 {
		return errors.New("content rewriting name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/load_balance_content_rewriting?mkey=%s", c.Address, name), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("content rewriting deletion failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("content rewriting deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
