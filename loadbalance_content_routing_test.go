package gofortiadc

import (
	"testing"
)

func TestClient_LoadbalanceGetContentRoutings(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetContentRoutings()
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetContentRouting(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.LoadbalanceGetContentRouting("goforticr01")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreateContentRouting(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRouting{
		Mkey:                  "goforticr01",
		Type:                  "l7-content-routing",
		PacketFwdMethod:       "inherit",
		SourcePoolList:        "",
		Persistence:           "",
		PersistenceInherit:    "enable",
		Method:                "",
		MethodInherit:         "enable",
		ConnectionPool:        "",
		ConnectionPoolInherit: "enable",
		Pool:                  "GOFORTI_POOL",
		IP:                    "0.0.0.0/0",
		IP6:                   "::/0",
		Comments:              "",
		ScheduleList:          "disable",
		SchedulePoolList:      "",
	}

	err = client.LoadbalanceCreateContentRouting(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceUpdateContentRouting(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalanceContentRouting{
		Mkey:                  "goforticr01",
		Type:                  "l7-content-routing",
		PacketFwdMethod:       "inherit",
		SourcePoolList:        "",
		Persistence:           "",
		PersistenceInherit:    "enable",
		Method:                "LB_METHOD_LEAST_CONNECTION",
		MethodInherit:         "disable",
		ConnectionPool:        "",
		ConnectionPoolInherit: "enable",
		Pool:                  "GOFORTI_POOL",
		IP:                    "0.0.0.0/0",
		IP6:                   "::/0",
		Comments:              "",
		ScheduleList:          "disable",
		SchedulePoolList:      "",
	}

	err = client.LoadbalanceUpdateContentRouting(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeleteContentRouting(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeleteContentRouting("goforticr01")
	if err != nil {
		t.Fatal(err)
	}
}
