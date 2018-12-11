package gofortiadc

import (
	"testing"
)

func TestClient_LoadbalanceGetContentRoutingConditions(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetContentRoutingConditions("goforticr01")
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetContentRoutingCondition(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.LoadbalanceGetContentRoutingCondition("goforticr01", "2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreateContentRoutingCondition(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRoutingCondition{
		Mkey:    "",
		Object:  "http-host-header",
		Type:    "string",
		Content: "gofortiadc.fakedomain.local",
		Reverse: "disable",
	}

	err = client.LoadbalanceCreateContentRoutingCondition("goforticr01", req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceUpdateContentRoutingCondition(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRoutingCondition{
		Mkey:    "1",
		Object:  "http-request-url",
		Type:    "string",
		Content: "/goforti.html",
		Reverse: "disable",
	}

	err = client.LoadbalanceUpdateContentRoutingCondition("goforticr01", req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeleteContentRoutingCondition(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeleteContentRoutingCondition("goforticr01", "1")
	if err != nil {
		t.Fatal(err)
	}
}
