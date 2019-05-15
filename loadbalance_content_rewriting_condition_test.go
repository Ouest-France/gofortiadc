package gofortiadc

import (
	"fmt"
	"os"
	"testing"
)

func TestClient_LoadbalanceGetContentRewritingConditions(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetContentRewritingConditions("gofortirw01")
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetContentRewritingCondition(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.LoadbalanceGetContentRewritingCondition("gofortirw01", "1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreateContentRewritingCondition(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRewritingCondition{
		Mkey:       "",
		Content:    "match",
		Ignorecase: "enable",
		Object:     "http-host-header",
		Reverse:    "disable",
		Type:       "string",
	}

	err = client.LoadbalanceCreateContentRewritingCondition("gofortirw01", req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetContentRewritingConditionID(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRewritingCondition{
		Mkey:       "",
		Content:    "match",
		Ignorecase: "enable",
		Object:     "http-host-header",
		Reverse:    "disable",
		Type:       "string",
	}

	id, err := client.LoadbalanceGetContentRewritingConditionID("gofortirw01", req)
	if err != nil {
		t.Fatal(err)
	}

	if id != "1" {
		t.Fatal(fmt.Errorf("id should be 1 but has a value of %s", id))
	}
}

func TestClient_LoadbalanceUpdateContentRewritingCondition(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRewritingCondition{
		Mkey:       "1",
		Content:    "127.0.0.1",
		Ignorecase: "disable",
		Object:     "ip-source-address",
		Reverse:    "disable",
		Type:       "string",
	}

	err = client.LoadbalanceUpdateContentRewritingCondition("gofortirw01", req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeleteContentRewritingCondition(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeleteContentRewritingCondition("gofortirw01", "1")
	if err != nil {
		t.Fatal(err)
	}
}
