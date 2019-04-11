package gofortiadc

import (
	"os"
	"testing"
)

func TestClient_SystemGetLocalCertificateGroupMembers(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.SystemGetLocalCertificateGroupMembers("goforti_group")
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_SystemGetLocalCertificateGroupMember(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.SystemGetLocalCertificateGroupMember("goforti_group", "1")
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_SystemCreateLocalCertificateGroupMember(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	member := SystemLocalCertificateGroupMember{
		Default:   "disable",
		LocalCert: "gofortiadc",
	}

	err = client.SystemCreateLocalCertificateGroupMember("goforti_group", member)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SystemUpdateLocalCertificateGroupMember(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	member := SystemLocalCertificateGroupMember{
		Default:         "enable",
		IntermediateCag: "Letsencrypt",
		LocalCert:       "gofortiadc",
		Mkey:            "1",
	}

	err = client.SystemUpdateLocalCertificateGroupMember("goforti_group", "1", member)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SystemDeleteLocalCertificateGroupMember(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.SystemDeleteLocalCertificateGroupMember("goforti_group", "1")
	if err != nil {
		t.Fatal(err)
	}
}
