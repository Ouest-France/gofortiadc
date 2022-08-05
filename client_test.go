package gofortiadc

import (
	"testing"
)

// TestClient_NewRequest creates an http.Request with authorization header set
func TestClient_NewRequest(t *testing.T) {

	// Create Table test case
	type testCase struct {
		method string
		url    string
		host   string
		path   string
		vdom   string
		query  string
	}

	// Create test cases
	testCases := []testCase{
		{
			method: "GET",
			url:    "http://localhost:8080/api/v1/status/system",
			host:   "localhost:8080",
			path:   "/api/v1/status/system",
			vdom:   "",
			query:  "",
		},
		{
			method: "POST",
			url:    "http://localhost:8080/api/v1/status/system",
			host:   "localhost:8080",
			path:   "/api/v1/status/system",
			vdom:   "root",
			query:  "vdom=root",
		},
		{
			method: "POST",
			url:    "http://localhost:8080/api/load_balance_virtual_server?mkey=test",
			host:   "localhost:8080",
			path:   "/api/load_balance_virtual_server",
			vdom:   "root",
			query:  "mkey=test&vdom=root",
		},
	}
	// Iterate over test cases
	for _, tc := range testCases {

		// Create client
		client, err := NewClientHelper()
		if err != nil {
			t.Fatal(err)
		}

		// Set vdom parameter on client
		client.VDom = tc.vdom

		// Create request
		req, err := client.NewRequest(tc.method, tc.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Test if req.Host is set with the correct value
		if req.Host != tc.host {
			t.Errorf("req.Host is not set with the correct value, expected %s, got %s", tc.host, req.Host)
		}

		// Test if req.Path is set with the correct value
		if req.URL.Path != tc.path {
			t.Errorf("req.Path is not set with the correct value, expected: %s, got: %s", tc.path, req.URL.Path)
		}

		// Test if authorization header is set with the correct value
		if req.Header.Get("Authorization") != "Bearer "+client.Token {
			t.Errorf("Authorization header is not set")
		}

		// Test if vdom parameter is set when vdom is set on client
		if tc.vdom != "" && req.URL.Query().Get("vdom") != tc.vdom {
			t.Errorf("vdom parameter is not set")
		}

		// Test if vdom parameter is not set when vdom is not set on client
		if tc.vdom == "" && req.URL.Query().Get("vdom") != "" {
			t.Errorf("vdom parameter is set")
		}

		// Test if query parameter is set with the correct value
		if tc.query != "" && req.URL.RawQuery != tc.query {
			t.Errorf("query parameter is not set with the correct value, expected: %s, got: %s", tc.query, req.URL.RawQuery)
		}
	}
}
