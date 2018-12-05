package gofortiadc

import "testing"

func TestClient_LoadbalanceGetPoolMembers(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetPoolMembers("GOFORTI_POOL")
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceGetPoolMember(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.LoadbalanceGetPoolMember("GOFORTI_POOL", "1")
	if err != nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceCreatePoolMember(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalancePoolMember{
		Address:                  "0.0.0.0",
		Address6:                 "::",
		Backup:                   "disable",
		ConnectionRateLimit:      "0",
		Connlimit:                "0",
		Cookie:                   "",
		HcStatus:                 "1",
		HealthCheckInherit:       "enable",
		MHealthCheck:             "disable",
		MHealthCheckRelationship: "AND",
		MHealthCheckList:         "",
		MysqlGroupID:             "0",
		MysqlReadOnly:            "disable",
		Port:                     "80",
		RealServerID:             "gofortirs01",
		Recover:                  "0",
		RsProfileInherit:         "enable",
		Ssl:                      "disable",
		Status:                   "enable",
		Weight:                   "1",
		Warmup:                   "0",
		Warmrate:                 "100",
	}

	err = client.LoadbalanceCreatePoolMember("GOFORTI_POOL", req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LoadbalanceUpdatePoolMember(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	req := LoadbalancePoolMember{
		Mkey:                     "1",
		Address:                  "0.0.0.0",
		Address6:                 "::",
		Backup:                   "disable",
		ConnectionRateLimit:      "0",
		Connlimit:                "0",
		Cookie:                   "",
		HcStatus:                 "1",
		HealthCheckInherit:       "enable",
		MHealthCheck:             "disable",
		MHealthCheckRelationship: "AND",
		MHealthCheckList:         "",
		MysqlGroupID:             "0",
		MysqlReadOnly:            "disable",
		Port:                     "8080",
		RealServerID:             "gofortirs01",
		Recover:                  "0",
		RsProfileInherit:         "enable",
		Ssl:                      "disable",
		Status:                   "enable",
		Weight:                   "1",
		Warmup:                   "0",
		Warmrate:                 "100",
	}

	err = client.LoadbalanceUpdatePoolMember("GOFORTI_POOL", "1", req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Client_LoadbalanceDeletePoolMember(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.LoadbalanceDeletePoolMember("GOFORTI_POOL", "1")
	if err != nil {
		t.Fatal(err)
	}
}
