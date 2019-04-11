package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// SystemLocalCertificateGroup represents a local certificate request/response
type SystemLocalCertificateGroupMember struct {
	OCSPStapling         string `json:"OCSP_stapling"`
	Default              string `json:"default"`
	ExtraOCSPStapling    string `json:"extra_OCSP_stapling"`
	ExtraIntermediateCag string `json:"extra_intermediate_cag"`
	ExtraLocalCert       string `json:"extra_local_cert"`
	IntermediateCag      string `json:"intermediate_cag"`
	LocalCert            string `json:"local_cert"`
	Mkey                 string `json:"mkey"`
}

// SystemGetLocalCertificateGroupMembers returns the list of all local certificate group members
func (c *Client) SystemGetLocalCertificateGroupMembers(group string) ([]SystemLocalCertificateGroupMember, error) {

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/system_certificate_local_cert_group_child_group_member?pkey=%s", c.Address, group), nil)
	if err != nil {
		return []SystemLocalCertificateGroupMember{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []SystemLocalCertificateGroupMember{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []SystemLocalCertificateGroupMember{}, fmt.Errorf("failed to get local certificate group members list with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []SystemLocalCertificateGroupMember{}, err
	}

	var SystemLocalCertificateGroupMembersPayload struct {
		Payload []SystemLocalCertificateGroupMember
	}
	err = json.Unmarshal(body, &SystemLocalCertificateGroupMembersPayload)
	if err != nil {
		return []SystemLocalCertificateGroupMember{}, err
	}

	return SystemLocalCertificateGroupMembersPayload.Payload, nil
}

// SystemGetLocalCertificateGroupMember returns a local certificate group member by mkey
func (c *Client) SystemGetLocalCertificateGroupMember(group, mkey string) (SystemLocalCertificateGroupMember, error) {

	members, err := c.SystemGetLocalCertificateGroupMembers(group)
	if err != nil {
		return SystemLocalCertificateGroupMember{}, err
	}

	for _, member := range members {
		if member.Mkey == mkey {
			return member, nil
		}
	}

	return SystemLocalCertificateGroupMember{}, fmt.Errorf("local certificate group member %s not found", mkey)
}

// SystemCreateLocalCertificateGroupMember creates a new local certificate group member
func (c *Client) SystemCreateLocalCertificateGroupMember(group string, lcgm SystemLocalCertificateGroupMember) error {

	payloadJSON, err := json.Marshal(lcgm)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/system_certificate_local_cert_group_child_group_member?pkey=%s", c.Address, group), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("local certificate group member creation failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("local certificate group member creation failed: %s ", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// SystemUpdateLocalCertificateGroupMember updates an existing local certificate group member
func (c *Client) SystemUpdateLocalCertificateGroupMember(group, mkey string, lcgm SystemLocalCertificateGroupMember) error {

	payloadJSON, err := json.Marshal(lcgm)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s/api/system_certificate_local_cert_group_child_group_member?mkey=%s&pkey=%s", c.Address, mkey, group), bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("local certificate group member update failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("local certificate group member update failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}

// SystemDeleteLocalCertificateGroupMember deletes an existing local certificate group member
func (c *Client) SystemDeleteLocalCertificateGroupMember(group, mkey string) error {

	if len(group) == 0 {
		return errors.New("local certificate group name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/system_certificate_local_cert_group_child_group_member?mkey=%s&pkey=%s", c.Address, mkey, group), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("local certificate group member deletion failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("local certificate group member deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
