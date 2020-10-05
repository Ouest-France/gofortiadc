package gofortiadc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadbalance(t *testing.T) {
	client, err := NewClientHelper()
	require.NoError(t, err, "NewClientHelper")

	// Create real server
	reqCreateRealServer := LoadbalanceRealServer{
		Status:   "enable",
		Address:  "128.1.52.12",
		Address6: "::",
		Mkey:     "gofortirs01",
	}

	err = client.LoadbalanceCreateRealServer(reqCreateRealServer)
	require.NoError(t, err, "LoadbalanceCreateRealServer")

	defer func() {
		// Delete real server
		err = client.LoadbalanceDeleteRealServer("gofortirs01")
		require.NoError(t, err, "LoadbalanceDeleteRealServer")
	}()

	// Update real server
	reqUpdateRealServer := LoadbalanceRealServer{
		Status:   "enable",
		Address:  "128.1.52.13",
		Address6: "::",
		Mkey:     "gofortirs01",
	}

	err = client.LoadbalanceUpdateRealServer(reqUpdateRealServer)
	require.NoError(t, err, "LoadbalanceUpdateRealServer")

	// Create real server pool
	reqCreateRealServerPool := LoadbalancePool{
		Mkey:                    "GOFORTI_POOL",
		PoolType:                "ipv4",
		HealthCheck:             "enable",
		HealthCheckRelationship: "AND",
		HealthCheckList:         "LB_HLTHCK_HTTP LB_HLTHCK_HTTPS",
		RsProfile:               "NONE",
	}

	err = client.LoadbalanceCreatePool(reqCreateRealServerPool)
	require.NoError(t, err, "LoadbalanceCreatePool")

	defer func() {
		// Delete real server pool
		err = client.LoadbalanceDeletePool("GOFORTI_POOL")
		require.NoError(t, err, "LoadbalanceDeletePool")
	}()

	// Update real server pool
	reqUpdateRealServerPool := LoadbalancePool{
		Mkey:                    "GOFORTI_POOL",
		PoolType:                "ipv4",
		HealthCheck:             "enable",
		HealthCheckRelationship: "AND",
		HealthCheckList:         "LB_HLTHCK_HTTP",
		RsProfile:               "NONE",
	}

	err = client.LoadbalanceUpdatePool(reqUpdateRealServerPool)
	require.NoError(t, err, "LoadbalanceUpdatePool")

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

	err = client.LoadbalanceCreatePoolMember("GOFORTI_POOL", reqCreateRealServerPoolMember)
	require.NoError(t, err, "LoadbalanceCreatePoolMember")

	defer func() {
		// Delete real server pool member
		err = client.LoadbalanceDeletePoolMember("GOFORTI_POOL", "1")
		require.NoError(t, err, "LoadbalanceDeletePoolMember")
	}()

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

	err = client.LoadbalanceUpdatePoolMember("GOFORTI_POOL", "1", reqUpdateRealServerPoolMember)
	require.NoError(t, err, "LoadbalanceUpdatePoolMember")

	// Create virtual server
	reqCreateVirtualServer := LoadbalanceVirtualServer{
		AddrType:            "ipv4",
		Address:             "128.1.201.35",
		Alone:               "enable",
		ConnectionLimit:     "10000",
		ConnectionRateLimit: "0",
		ContentRewriting:    "disable",
		ContentRouting:      "disable",
		HTTP2HTTPS:          "disable",
		Interface:           "port1",
		Method:              "LB_METHOD_ROUND_ROBIN",
		Mkey:                "GOFORTI-VS",
		PacketFwdMethod:     "NAT",
		Pool:                "GOFORTI_POOL",
		Port:                "80",
		Profile:             "LB_PROF_TCP",
		Status:              "enable",
		TrafficLog:          "enable",
		Type:                "l4-load-balance",
		Warmrate:            "10",
		Warmup:              "0",
	}

	err = client.LoadbalanceCreateVirtualServer(reqCreateVirtualServer)
	require.NoError(t, err, "LoadbalanceCreateVirtualServer")

	defer func() {
		// Delete virtual server
		err = client.LoadbalanceDeleteVirtualServer("GOFORTI-VS")
		require.NoError(t, err, "LoadbalanceDeleteVirtualServer")
	}()

	// Update virtual server
	reqUpdateVirtualServer := LoadbalanceVirtualServer{
		AddrType:            "ipv4",
		Address:             "128.1.201.35",
		Alone:               "enable",
		ConnectionLimit:     "10000",
		ConnectionRateLimit: "0",
		ContentRewriting:    "disable",
		ContentRouting:      "disable",
		Interface:           "port1",
		Method:              "LB_METHOD_FASTEST_RESPONSE",
		Mkey:                "GOFORTI-VS",
		PacketFwdMethod:     "NAT",
		Pool:                "GOFORTI_POOL",
		Port:                "80",
		Profile:             "LB_PROF_TCP",
		Status:              "enable",
		TrafficLog:          "enable",
		Type:                "l4-load-balance",
		Warmrate:            "10",
		Warmup:              "0",
	}

	err = client.LoadbalanceUpdateVirtualServer(reqUpdateVirtualServer)
	require.NoError(t, err, "LoadbalanceUpdateVirtualServer")

	// List content rewritings
	_, err = client.LoadbalanceGetContentRewritings()
	require.NoError(t, err, "LoadbalanceGetContentRewritings")

	// Create content rewriting
	reqCW := LoadbalanceContentRewriting{
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

	err = client.LoadbalanceCreateContentRewriting(reqCW)
	require.NoError(t, err, "LoadbalanceCreateContentRewriting")

	defer func() {
		// Delete content rewriting
		err = client.LoadbalanceDeleteContentRewriting("gofortirw01")
		require.NoError(t, err, "LoadbalanceDeleteContentRewriting")
	}()

	// Get content rewriting
	_, err = client.LoadbalanceGetContentRewriting("gofortirw01")
	require.NoError(t, err, "LoadbalanceGetContentRewriting")

	// Update content rewriting
	reqCW = LoadbalanceContentRewriting{
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

	err = client.LoadbalanceUpdateContentRewriting(reqCW)
	require.NoError(t, err, "LoadbalanceUpdateContentRewriting")

	// Get content rewriting conditions
	_, err = client.LoadbalanceGetContentRewritingConditions("gofortirw01")
	require.NoError(t, err, "LoadbalanceGetContentRewritingConditions")

	// Create content rewriting condition
	condReqCW := LoadbalanceContentRewritingCondition{
		Mkey:       "",
		Content:    "match",
		Ignorecase: "enable",
		Object:     "http-host-header",
		Reverse:    "disable",
		Type:       "string",
	}

	err = client.LoadbalanceCreateContentRewritingCondition("gofortirw01", condReqCW)
	require.NoError(t, err, "LoadbalanceCreateContentRewritingCondition")

	defer func() {
		err = client.LoadbalanceDeleteContentRewritingCondition("gofortirw01", "1")
		require.NoError(t, err, "LoadbalanceDeleteContentRewritingCondition")
	}()

	// Get content rewriting condition
	_, err = client.LoadbalanceGetContentRewritingCondition("gofortirw01", "1")
	require.NoError(t, err, "LoadbalanceGetContentRewritingCondition")

	// Get content rewriting condition ID
	condReqCW = LoadbalanceContentRewritingCondition{
		Mkey:       "",
		Content:    "match",
		Ignorecase: "enable",
		Object:     "http-host-header",
		Reverse:    "disable",
		Type:       "string",
	}

	id, err := client.LoadbalanceGetContentRewritingConditionID("gofortirw01", condReqCW)
	require.NoError(t, err, "LoadbalanceGetContentRewritingConditionID")
	require.Equal(t, "1", id, "LoadbalanceGetContentRewritingConditionID")

	// Update content rewriting condition
	condReqCW = LoadbalanceContentRewritingCondition{
		Mkey:       "1",
		Content:    "127.0.0.1",
		Ignorecase: "disable",
		Object:     "ip-source-address",
		Reverse:    "disable",
		Type:       "string",
	}

	err = client.LoadbalanceUpdateContentRewritingCondition("gofortirw01", condReqCW)

	// Get content routings
	_, err = client.LoadbalanceGetContentRoutings()
	require.NoError(t, err, "LoadbalanceGetContentRoutings")

	// Create content routing
	reqCR := LoadbalanceContentRouting{
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

	err = client.LoadbalanceCreateContentRouting(reqCR)
	require.NoError(t, err, "LoadbalanceCreateContentRouting")

	defer func() {
		// Delete content routing
		err = client.LoadbalanceDeleteContentRouting("goforticr01")
		require.NoError(t, err, "LoadbalanceDeleteContentRouting")
	}()

	// Get content routing
	_, err = client.LoadbalanceGetContentRouting("goforticr01")
	require.NoError(t, err, "LoadbalanceGetContentRouting")

	// Update content routing
	reqCR = LoadbalanceContentRouting{
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

	err = client.LoadbalanceUpdateContentRouting(reqCR)
	require.NoError(t, err, "LoadbalanceUpdateContentRouting")

	// Get content routing conditions
	_, err = client.LoadbalanceGetContentRoutingConditions("goforticr01")
	require.NoError(t, err, "LoadbalanceGetContentRoutingConditions")

	// Create content routing condition
	condReqCR := LoadbalanceContentRoutingCondition{
		Mkey:    "",
		Object:  "http-host-header",
		Type:    "string",
		Content: "gofortiadc.fakedomain.local",
		Reverse: "disable",
	}

	err = client.LoadbalanceCreateContentRoutingCondition("goforticr01", condReqCR)
	require.NoError(t, err, "LoadbalanceCreateContentRoutingCondition")

	defer func() {
		// Delete content routing condition
		err = client.LoadbalanceDeleteContentRoutingCondition("goforticr01", "1")
		require.NoError(t, err, "LoadbalanceDeleteContentRoutingCondition")
	}()

	// Get content routing condition
	_, err = client.LoadbalanceGetContentRoutingCondition("goforticr01", "1")
	require.NoError(t, err, "LoadbalanceGetContentRoutingCondition")

	// Update content routing condition
	condReqCR = LoadbalanceContentRoutingCondition{
		Mkey:    "1",
		Object:  "http-request-url",
		Type:    "string",
		Content: "/goforti.html",
		Reverse: "disable",
	}

	err = client.LoadbalanceUpdateContentRoutingCondition("goforticr01", condReqCR)
}
