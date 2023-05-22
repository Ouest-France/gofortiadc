package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SystemLocalCertificateGroup represents a local certificate request/response
type SystemLocalCertificateGroup struct {
	Mkey string `json:"mkey"`
}

// SystemGetLocalCertificateGroups returns the list of all local certificate groups
func (c *Client) SystemGetLocalCertificateGroups() ([]SystemLocalCertificateGroup, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/system_certificate_local_cert_group", c.Address), nil)
	if err != nil {
		return []SystemLocalCertificateGroup{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []SystemLocalCertificateGroup{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []SystemLocalCertificateGroup{}, fmt.Errorf("failed to get local certificate groups list with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []SystemLocalCertificateGroup{}, err
	}

	var SystemLocalCertificateGroupsPayload struct {
		Payload []SystemLocalCertificateGroup
	}
	err = json.Unmarshal(body, &SystemLocalCertificateGroupsPayload)
	if err != nil {
		return []SystemLocalCertificateGroup{}, err
	}

	return SystemLocalCertificateGroupsPayload.Payload, nil
}

// SystemGetLocalCertificateGroup returns a local certificate group by name
func (c *Client) SystemGetLocalCertificateGroup(name string) (SystemLocalCertificateGroup, error) {

	groups, err := c.SystemGetLocalCertificateGroups()
	if err != nil {
		return SystemLocalCertificateGroup{}, err
	}

	for _, group := range groups {
		if group.Mkey == name {
			return group, nil
		}
	}

	return SystemLocalCertificateGroup{}, fmt.Errorf("local certificate group %s not found: %w", name, ErrNotFound)
}

// SystemCreateLocalCertificateGroup creates a new local certificate group
func (c *Client) SystemCreateLocalCertificateGroup(lcg SystemLocalCertificateGroup) error {

	payloadJSON, err := json.Marshal(lcg)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/system_certificate_local_cert_group", c.Address), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("local certificate group creation failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resJSON := struct{ Payload int }{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return err
	}

	if resJSON.Payload != 0 {
		return fmt.Errorf("local certificate group creation failed: %s ", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// SystemDeleteLocalCertificateGroup deletes an existing local certificate group
func (c *Client) SystemDeleteLocalCertificateGroup(name string) error {

	if len(name) == 0 {
		return errors.New("local certificate group name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/system_certificate_local_cert_group?mkey=%s", c.Address, name), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("local certificate group deletion failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resJSON := struct{ Payload int }{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return err
	}

	if resJSON.Payload != 0 {
		return fmt.Errorf("local certificate group deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
