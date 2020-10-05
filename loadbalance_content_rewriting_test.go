package gofortiadc

import (
	"os"
	"testing"
)

func TestClient_LoadbalanceGetContentRewritings(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetContentRewritings()
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetContentRewriting(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.LoadbalanceGetContentRewriting("gofortirw01")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreateContentRewriting(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRewriting{
		Mkey:           "gofortirw01",
		ActionType:     "request",
		URLStatus:      "enable",
		URLContent:     "/url",
		RefererStatus:  "enable",
		RefererContent: "http://",
		Redirect:       "redirect",
		Location:       "http://",
		HeaderName:     "header-name",
		Comments:       "",
		Action:         "rewrite_http_header",
		HostStatus:     "enable",
		HostContent:    "host",
	}

	err = client.LoadbalanceCreateContentRewriting(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceUpdateContentRewriting(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRewriting{
		Mkey:           "gofortirw01",
		ActionType:     "request",
		URLStatus:      "disable",
		URLContent:     "/url",
		RefererStatus:  "enable",
		RefererContent: "http://foo.bar",
		Redirect:       "redirect",
		Location:       "http://",
		HeaderName:     "header-name",
		Comments:       "",
		Action:         "rewrite_http_header",
		HostStatus:     "disable",
		HostContent:    "host",
	}

	err = client.LoadbalanceUpdateContentRewriting(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeleteContentRewriting(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeleteContentRewriting("gofortirw01")
	if err != nil {
		t.Fatal(err)
	}
}
