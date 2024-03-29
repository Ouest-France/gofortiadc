package gofortiadc

import (
	"encoding/json"
	"fmt"
	"io"
)

// SystemGlobalRes represents a system global status respoonse
type SystemGlobalRes struct {
	HTTPPort                string `json:"http-port"`
	HTTPSPort               string `json:"https-port"`
	SSHPort                 string `json:"ssh-port"`
	TelnetPort              string `json:"telnet-port"`
	AdminIdleTimeout        string `json:"admin-idle-timeout"`
	SysGlobalLanguage       string `json:"sys-global-language"`
	Hostname                string `json:"hostname"`
	Theme                   string `json:"theme"`
	GuiSystem               string `json:"gui-system"`
	GuiRouter               string `json:"gui-router"`
	GuiLog                  string `json:"gui-log"`
	GuiLoadBalance          string `json:"gui-load-balance"`
	GuiGlobalDNSLoadBalance string `json:"gui-global-dns-load-balance"`
	GuiFirewall             string `json:"gui-firewall"`
	GuiLinkLoadBalance      string `json:"gui-link-load-balance"`
	GuiSecurity             string `json:"gui-security"`
	VdomAdmin               string `json:"vdom-admin"`
	IPPrimary               string `json:"ip_primary"`
	IPSecond                string `json:"ip_second"`
	IsSystemAdmin           bool   `json:"isSystemAdmin"`
}

// SystemGlobal returns system global status
func (c *Client) SystemGlobal() (SystemGlobalRes, error) {

	req, err := c.NewRequest("GET", fmt.Sprintf("%s/api/system_global", c.Address), nil)
	if err != nil {
		return SystemGlobalRes{}, fmt.Errorf("Failed create http request: %s", err)
	}

	get, err := c.Client.Do(req)
	if err != nil {
		return SystemGlobalRes{}, err
	}
	defer get.Body.Close()

	if get.StatusCode != 200 {
		return SystemGlobalRes{}, fmt.Errorf("Failed to get system endpoint with http code: %d", get.StatusCode)
	}

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return SystemGlobalRes{}, err
	}

	var systemGlobalRes struct{ Payload SystemGlobalRes }
	err = json.Unmarshal(body, &systemGlobalRes)
	if err != nil {
		return SystemGlobalRes{}, err
	}

	return systemGlobalRes.Payload, nil
}
