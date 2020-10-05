package gofortiadc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystemCertficate(t *testing.T) {
	client, err := NewClientHelper()
	require.NoError(t, err, "NewClientHelper")

	// Create local certificate
	cert, key, err := generateCertificate()
	require.NoError(t, err, "generateCertificate")

	err = client.SystemCreateLocalCertificate("gofortiadc", "", cert, key)
	require.NoError(t, err, "SystemCreateLocalCertificate")

	defer func() {
		// Delete local certificate
		err = client.SystemDeleteLocalCertificate("gofortiadc")
		require.NoError(t, err, "SystemDeleteLocalCertificate")
	}()

	// Get local certificate
	certRes, err := client.SystemGetLocalCertificate("gofortiadc")
	require.NoError(t, err, "SystemGetLocalCertificate")
	require.Equal(t, "/O=Gofortiadc Test", certRes.Subject)

	// Create local certificate group
	group := SystemLocalCertificateGroup{
		Mkey: "goforti_group",
	}

	err = client.SystemCreateLocalCertificateGroup(group)
	require.NoError(t, err, "SystemCreateLocalCertificateGroup")

	defer func() {
		// Delete local certificate group
		err = client.SystemDeleteLocalCertificateGroup("goforti_group")
		require.NoError(t, err, "SystemDeleteLocalCertificateGroup")
	}()

	// Get local certificate group
	_, err = client.SystemGetLocalCertificateGroup(group.Mkey)
	require.NoError(t, err, "SystemGetLocalCertificateGroup")

	// Create local certificate group member
	member := SystemLocalCertificateGroupMember{
		Default:   "disable",
		LocalCert: "gofortiadc",
	}

	err = client.SystemCreateLocalCertificateGroupMember("goforti_group", member)
	require.NoError(t, err, "SystemCreateLocalCertificateGroupMember")

	defer func() {
		// Delete local certificate group member
		err = client.SystemDeleteLocalCertificateGroupMember("goforti_group", "1")
		require.NoError(t, err, "SystemDeleteLocalCertificateGroupMember")
	}()

	// Update local certificate group member
	member = SystemLocalCertificateGroupMember{
		Default:         "enable",
		IntermediateCag: "Letsencrypt",
		LocalCert:       "gofortiadc",
		Mkey:            "1",
	}

	err = client.SystemUpdateLocalCertificateGroupMember("goforti_group", "1", member)
	require.NoError(t, err, "SystemUpdateLocalCertificateGroupMember")
}
