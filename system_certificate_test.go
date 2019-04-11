package gofortiadc

import (
	"testing"
)

func TestSystemCertficate(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	// Create local certificate
	cert, key, err := generateCertificate()
	if err != nil {
		t.Fatal(err)
	}

	err = client.SystemCreateLocalCertificate("gofortiadc", "", cert, key)
	if err != nil {
		t.Fatalf("SystemCreateLocalCertificate failed with error: %s", err)
	}

	// Create local certificate group
	group := SystemLocalCertificateGroup{
		Mkey: "goforti_group",
	}

	err = client.SystemCreateLocalCertificateGroup(group)
	if err != nil {
		t.Fatalf("SystemCreateLocalCertificateGroup failed with error: %s", err)
	}

	// Create local certificate group member
	member := SystemLocalCertificateGroupMember{
		Default:   "disable",
		LocalCert: "gofortiadc",
	}

	err = client.SystemCreateLocalCertificateGroupMember("goforti_group", member)
	if err != nil {
		t.Fatalf("SystemCreateLocalCertificateGroupMember failed with error: %s", err)
	}

	// Update local certificate group member
	member = SystemLocalCertificateGroupMember{
		Default:         "enable",
		IntermediateCag: "Letsencrypt",
		LocalCert:       "gofortiadc",
		Mkey:            "1",
	}

	err = client.SystemUpdateLocalCertificateGroupMember("goforti_group", "1", member)
	if err != nil {
		t.Fatalf("SystemUpdateLocalCertificateGroupMember failed with error: %s", err)
	}

	// Delete local certificate group member
	err = client.SystemDeleteLocalCertificateGroupMember("goforti_group", "1")
	if err != nil {
		t.Fatalf("SystemDeleteLocalCertificateGroupMember failed with error: %s", err)
	}

	// Delete local certificate group
	err = client.SystemDeleteLocalCertificateGroup("goforti_group")
	if err != nil {
		t.Fatalf("SystemDeleteLocalCertificateGroup failed with error: %s", err)
	}

	// Delete local certificate
	err = client.SystemDeleteLocalCertificate("gofortiadc")
	if err != nil {
		t.Fatalf("SystemDeleteLocalCertificate failed with error: %s", err)
	}
}
