package gofortiadc

import (
	"testing"
)

func TestLoadbalance(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	// Create real server
	reqCreateRealServer := LoadbalanceRealServer{
		Status:   "enable",
		Address:  "128.1.52.12",
		Address6: "::",
		Mkey:     "gofortirs01",
	}
	t.Logf("reqCreateRealServer: %+v", reqCreateRealServer)

	err = client.LoadbalanceCreateRealServer(reqCreateRealServer)
	if err != nil {
		t.Fatalf("LoadbalanceCreateRealServer failed with error: %s", err)
	}

	// Update real server
	reqUpdateRealServer := LoadbalanceRealServer{
		Status:   "enable",
		Address:  "128.1.52.13",
		Address6: "::",
		Mkey:     "gofortirs01",
	}
	t.Logf("reqUpdateRealServer: %+v", reqUpdateRealServer)

	err = client.LoadbalanceUpdateRealServer(reqUpdateRealServer)
	if err != nil {
		t.Fatalf("LoadbalanceUpdateRealServer failed with error: %s", err)
	}

	// Create real server pool
	reqCreateRealServerPool := LoadbalancePoolReq{
		Mkey:                    "GOFORTI_POOL",
		PoolType:                "ipv4",
		HealthCheck:             "enable",
		HealthCheckRelationship: "AND",
		HealthCheckList:         "LB_HLTHCK_HTTP LB_HLTHCK_HTTPS",
		RsProfile:               "NONE",
	}
	t.Logf("reqCreateRealServerPool: %+v", reqCreateRealServerPool)

	err = client.LoadbalanceCreatePool(reqCreateRealServerPool)
	if err != nil {
		t.Fatalf("LoadbalanceCreatePool failed with error: %s", err)
	}

	// Update real server pool
	reqUpdateRealServerPool := LoadbalancePoolReq{
		Mkey:                    "GOFORTI_POOL",
		PoolType:                "ipv4",
		HealthCheck:             "enable",
		HealthCheckRelationship: "AND",
		HealthCheckList:         "LB_HLTHCK_HTTP",
		RsProfile:               "NONE",
	}
	t.Logf("reqUpdateRealServerPool: %+v", reqUpdateRealServerPool)

	err = client.LoadbalanceUpdatePool(reqUpdateRealServerPool)
	if err != nil {
		t.Fatalf("LoadbalanceUpdatePool failed with error: %s", err)
	}

	// Create real server pool member
	reqCreateRealServerPoolMember := LoadbalancePoolMember{
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
	t.Logf("reqCreateRealServerPoolMember: %+v", reqCreateRealServerPoolMember)

	err = client.LoadbalanceCreatePoolMember("GOFORTI_POOL", reqCreateRealServerPoolMember)
	if err != nil {
		t.Fatalf("LoadbalanceCreatePoolMember failed with error: %s", err)
	}

	// Update real server pool member
	reqUpdateRealServerPoolMember := LoadbalancePoolMember{
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
	t.Logf("reqUpdateRealServerPoolMember: %+v", reqUpdateRealServerPoolMember)

	err = client.LoadbalanceUpdatePoolMember("GOFORTI_POOL", "1", reqUpdateRealServerPoolMember)
	if err != nil {
		t.Fatalf("LoadbalanceUpdatePoolMember failed with error: %s", err)
	}

	// Create virtual server
	reqCreateVirtualServer := LoadbalanceVirtualServerReq{
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
		HTTP2HTTPS:          "enable",
	}
	t.Logf("reqCreateVirtualServer: %+v", reqCreateVirtualServer)

	err = client.LoadbalanceCreateVirtualServer(reqCreateVirtualServer)
	if err != nil {
		t.Fatalf("LoadbalanceCreateVirtualServer failed with error: %s", err)
	}

	// Update virtual server
	reqUpdateVirtualServer := LoadbalanceVirtualServerReq{
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
	t.Logf("reqUpdateVirtualServer: %+v", reqUpdateVirtualServer)

	err = client.LoadbalanceUpdateVirtualServer(reqUpdateVirtualServer)
	if err != nil {
		t.Fatalf("LoadbalanceUpdateVirtualServer failed with error: %s", err)
	}

	// Delete virtual server
	err = client.LoadbalanceDeleteVirtualServer("GOFORTI-VS")
	if err != nil {
		t.Fatalf("LoadbalanceDeleteVirtualServer failed with error: %s", err)
	}

	// Delete real server pool member
	err = client.LoadbalanceDeletePoolMember("GOFORTI_POOL", "1")
	if err != nil {
		t.Fatalf("LoadbalanceDeletePoolMember failed with error: %s", err)
	}

	// Delete real server pool
	err = client.LoadbalanceDeletePool("GOFORTI_POOL")
	if err != nil {
		t.Fatalf("LoadbalanceDeletePool failed with error: %s", err)
	}

	// Delete real server
	err = client.LoadbalanceDeleteRealServer("gofortirs01")
	if err != nil {
		t.Fatalf("LoadbalanceDeleteRealServer failed with error: %s", err)
	}
}
