package gofortiadc

import (
	"os"
	"testing"
)

func TestClient_SystemGetLocalCertificateGroups(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.SystemGetLocalCertificateGroups()
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_SystemGetLocalCertificateGroup(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.SystemGetLocalCertificateGroup("goforti_group")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SystemCreateLocalCertificateGroup(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	group := SystemLocalCertificateGroup{
		Mkey: "goforti_group",
	}

	err = client.SystemCreateLocalCertificateGroup(group)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SystemDeleteLocalCertificateGroup(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.SystemDeleteLocalCertificateGroup("goforti_group")
	if err != nil {
		t.Fatal(err)
	}
}
