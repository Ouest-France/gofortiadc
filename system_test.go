package gofortiadc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystemGlobal(t *testing.T) {

	client, err := NewClientHelper()
	require.NoError(t, err)

	res, err := client.SystemGlobal()
	require.NoError(t, err)

	require.Equal(t, res.Hostname, os.Getenv("GOFORTI_HOSTNAME"))
}
