package gofortiadc

import (
	"testing"
)

func TestClient_LoadbalanceGetVirtualServers(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetVirtualServers()
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreateVirtualServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceVirtualServerReq{
		Status:              "enable",
		Type:                "l4-load-balance",
		AddrType:            "ipv4",
		Address:             "128.1.201.35",
		PacketFwdMethod:     "NAT",
		Port:                "80",
		PortRange:           "0",
		ConnectionLimit:     "10000",
		ContentRouting:      "disable",
		Warmup:              "0",
		Warmrate:            "10",
		ConnectionRateLimit: "0",
		Log:                 "enable",
		Alone:               "enable",
		Mkey:                "GOFORTI-VS",
		Interface:           "port1",
		Profile:             "LB_PROF_TCP",
		Method:              "LB_METHOD_ROUND_ROBIN",
		Pool:                "GOFORTI_POOL",
	}

	err = client.LoadbalanceCreateVirtualServer(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceUpdateVirtualServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceVirtualServerReq{
		Status:              "enable",
		Type:                "l4-load-balance",
		AddrType:            "ipv4",
		Address:             "128.1.201.35",
		PacketFwdMethod:     "NAT",
		Port:                "80",
		PortRange:           "0",
		ConnectionLimit:     "10000",
		ContentRouting:      "disable",
		Warmup:              "0",
		Warmrate:            "10",
		ConnectionRateLimit: "0",
		Log:                 "enable",
		Alone:               "enable",
		Mkey:                "GOFORTI-VS",
		Interface:           "port1",
		Profile:             "LB_PROF_TCP",
		Method:              "LB_METHOD_FASTEST_RESPONSE",
		Pool:                "GOFORTI_POOL",
	}

	err = client.LoadbalanceUpdateVirtualServer(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetVirtualServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.LoadbalanceGetVirtualServer("GOFORTI-VS")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeleteVirtualServer(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeleteVirtualServer("GOFORTI-VS")
	if err != nil {
		t.Fatal(err)
	}
}
