package gofortiadc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
)

// SystemLocalCertificate represents a real server request/response
type SystemLocalCertificate struct {
	CaType      string        `json:"ca_type"`
	Comments    string        `json:"comments"`
	Extension   []interface{} `json:"extension"`
	Fingerprint string        `json:"fingerprint"`
	Hash        string        `json:"hash"`
	Issuer      string        `json:"issuer"`
	Mkey        string        `json:"mkey"`
	PinSha256   string        `json:"pin-sha256"`
	Sn          string        `json:"sn"`
	Status      string        `json:"status"`
	Subject     string        `json:"subject"`
	Type        string        `json:"type"`
	Validfrom   string        `json:"validfrom"`
	Validto     string        `json:"validto"`
	Version     int           `json:"version"`
}

// SystemGetLocalCertificates returns the list of all local certificates
func (c *Client) SystemGetLocalCertificates() ([]SystemLocalCertificate, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/system_certificate_local", c.Address), nil)
	if err != nil {
		return []SystemLocalCertificate{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []SystemLocalCertificate{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []SystemLocalCertificate{}, fmt.Errorf("failed to get local certificates list with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []SystemLocalCertificate{}, err
	}

	var SystemLocalCertificatePayload struct {
		Payload []SystemLocalCertificate
	}
	err = json.Unmarshal(body, &SystemLocalCertificatePayload)
	if err != nil {
		return []SystemLocalCertificate{}, err
	}

	return SystemLocalCertificatePayload.Payload, nil
}

// SystemGetLocalCertificate returns a real server by name
func (c *Client) SystemGetLocalCertificate(name string) (SystemLocalCertificate, error) {

	certificates, err := c.SystemGetLocalCertificates()
	if err != nil {
		return SystemLocalCertificate{}, err
	}

	for _, certificate := range certificates {
		if certificate.Mkey == name {
			return certificate, nil
		}
	}

	return SystemLocalCertificate{}, fmt.Errorf("local certificate %s not found: %w", name, ErrNotFound)
}

// SystemCreateLocalCertificate creates a new local certificate
func (c *Client) SystemCreateLocalCertificate(name, password string, cert, key []byte) error {

	form, contentType, err := createCertificateForm(name, password, cert, key)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s/api/upload/certificate_local", c.Address), form)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("local certificate creation failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("local certificate creation failed: %s ", getErrorMessage(resJSON.Payload))
	}

	return nil
}

func createCertificateForm(name, password string, cert, key []byte) (form *bytes.Buffer, contentType string, err error) {

	var b bytes.Buffer

	w := multipart.NewWriter(&b)
	defer w.Close()

	fields := map[string]string{
		"mkey":   name,
		"vdom":   "global",
		"type":   "CertKey",
		"passwd": password,
	}
	for field, value := range fields {
		err := w.WriteField(field, value)
		if err != nil {
			return &b, contentType, err
		}
	}

	certPart, err := w.CreateFormFile("cert", "tls.crt")
	if err != nil {
		return &b, contentType, err
	}

	_, err = certPart.Write(cert)
	if err != nil {
		return &b, contentType, err
	}

	keyPart, err := w.CreateFormFile("key", "tls.key")
	if err != nil {
		return &b, contentType, err
	}

	_, err = keyPart.Write(key)

	return &b, w.FormDataContentType(), err
}

// SystemDeleteLocalCertificate deletes an existing local certificate
func (c *Client) SystemDeleteLocalCertificate(name string) error {

	if len(name) == 0 {
		return errors.New("local certificate name cannot be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s/api/system_certificate_local?mkey=%s", c.Address, name), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("local certificate deletion failed with status code: %d", res.StatusCode)
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
		return fmt.Errorf("local certificate deletion failed: %s", getErrorMessage(resJSON.Payload))
	}

	return nil
}
