package gofortiadc

import (
	"testing"
)

func TestSystemGlobal(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.SystemGlobal()
	if err != nil {
		t.Fatal(err)
	}
}
