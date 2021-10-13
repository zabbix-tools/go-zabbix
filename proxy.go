package zabbix

import (
	"fmt"
)

const (
	InterfaceUseIP  = 1
	InterfaceUseDNS = 0
)

const (
	PassiveProxy = 6
	ActiveProxy  = 5
)

type Proxy struct {
	ProxyID     string          `json:"proxyid"`
	ProxyName   string          `json:"host"`
	ProxyStatus int             `json:"status,string"`
	Hosts       []Host          `json:"hosts,omitempty"`
	Interface   *ProxyInterface `json:"interface,omitempty"`
}

type ProxyInterface struct {
	InterfaceID string `json:"interfaceid,omitempty"`
	HostID      string `json:"hostid,omitempty"`
	UseIP       int    `json:"useip,string"`
	IP          string `json:"ip"`
	DNS         string `json:"dns"`
	Port        string `json:"port"`
}

type ProxyResult struct {
	ProxyIDs []string `json:"proxyids"`
}

type ProxyGetParams struct {
	GetParameters

	ProxyIDs        []string    `json:"proxyids,omitempty"`
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectInterface SelectQuery `json:"selectInterface,omitempty"`
}

type ProxyUpdateParams struct {
	ProxyID string   `json:"proxyid"`
	Hosts   []string `json:"hosts"`
}

type ProxyCreateParams struct {
	ProxyName   string         `json:"host"`
	ProxyStatus int            `json:"status"`
	Hosts       []string       `json:"hosts,omitempty"`
	Interface   ProxyInterface `json:"interface,omitempty"`
}

type ProxyDeleteParams struct {
	ProxyID []string
}

func (c *Session) GetProxy(params ProxyGetParams) ([]Proxy, error) {
	proxy := make([]Proxy, 0)
	err := c.Get("proxy.get", params, &proxy)
	if err != nil {
		return nil, err
	}

	if len(proxy) == 0 {
		return nil, ErrNotFound
	}

	return proxy, nil
}

func (c *Session) UpdateProxy(params ProxyUpdateParams) (ProxyResult, error) {
	resp := ProxyResult{}
	err := c.Get("proxy.update", params, &resp)
	if err != nil {
		return ProxyResult{}, err
	}
	return resp, nil
}

func (c *Session) CreateProxy(params ProxyCreateParams) (ProxyResult, error) {
	if params.ProxyStatus != PassiveProxy && params.ProxyStatus != ActiveProxy {
		return ProxyResult{}, fmt.Errorf("Proxy status must be 5 or 6")
	}
	if params.ProxyStatus == PassiveProxy && params.Interface == (ProxyInterface{}) {
		return ProxyResult{}, fmt.Errorf("Passive proxy must settings interface")
	}
	resp := ProxyResult{}
	err := c.Get("proxy.create", params, &resp)
	if err != nil {
		return ProxyResult{}, err
	}
	return resp, nil
}

func (c *Session) DeleteProxy(params ProxyDeleteParams) (ProxyResult, error) {
	resp := ProxyResult{}
	err := c.Get("proxy.delete", params.ProxyID, &resp)
	if err != nil {
		return ProxyResult{}, err
	}
	return resp, nil
}
