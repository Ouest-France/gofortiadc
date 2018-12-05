package goforti

import (
	"testing"
)

func TestClient_LoadbalanceGetRealServers(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetRealServers()
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetRealServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.LoadbalanceGetRealServer("gofortirs01")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreateRealServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceRealServer{
		Status:   "enable",
		Address:  "128.1.52.12",
		Address6: "::",
		Mkey:     "gofortirs01",
	}

	err = client.LoadbalanceCreateRealServer(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceUpdateRealServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceRealServer{
		Status:   "enable",
		Address:  "128.1.52.13",
		Address6: "::",
		Mkey:     "gofortirs01",
	}

	err = client.LoadbalanceUpdateRealServer(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeleteRealServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeleteRealServer("gofortirs01")
	if err != nil {
		t.Fatal(err)
	}
}
