package gofortiadc

import (
	"os"
	"testing"
)

func TestClient_LoadbalanceGetPools(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetPools()
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetPool(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetPool("GOFORTI_POOL")
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreatePool(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalancePool{
		Mkey:                    "GOFORTI_POOL",
		PoolType:                "ipv4",
		HealthCheck:             "enable",
		HealthCheckRelationship: "AND",
		HealthCheckList:         "LB_HLTHCK_HTTP LB_HLTHCK_HTTPS",
		RsProfile:               "NONE",
	}

	err = client.LoadbalanceCreatePool(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceUpdatePool(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalancePool{
		Mkey:                    "GOFORTI_POOL",
		PoolType:                "ipv4",
		HealthCheck:             "enable",
		HealthCheckRelationship: "AND",
		HealthCheckList:         "LB_HLTHCK_HTTP",
		RsProfile:               "NONE",
	}

	err = client.LoadbalanceUpdatePool(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeletePool(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeletePool("GOFORTI_POOL")
	if err != nil {
		t.Fatal(err)
	}
}
