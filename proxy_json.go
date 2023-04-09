package zabbix

import (
	"fmt"
	"strconv"
	"time"
)

// jProxy is a private map for the Zabbix API Proxy object.
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/proxy/object
type jProxy struct {
	ProxyID        string `json:"proxyid,omitempty"`
	Hostname       string `json:"host"`
	Status         int    `json:"status,string"`
	Description    string `json:"description"`
	LastAccess     string `json:"lastaccess"`
	TLSConnect     int    `json:"tls_connect,string"`
	TLSAccept      int    `json:"tls_accept,string"`
	TLSIssuer      string `json:"tls_issuer"`
	TLSSubject     string `json:"tls_subject"`
	TLSPSKIdentity string `json:"tls_psk_identity"`
	TLSPSK         string `json:"tls_psk"`
	ProxyAddress   string `json:"proxy_address"`
	AutoCompress   int    `json:"auto_compress,string"`
}

// Proxy returns a native Go Proxy struct mapped from the given JSON Proxy data.
func (c *jProxy) Proxy() (*Proxy, error) {
	var err error
	var sec int64
	proxy := &Proxy{}
	proxy.ProxyID = c.ProxyID
	proxy.Hostname = c.Hostname
	proxy.Status = c.Status
	proxy.Description = c.Description

	// parse timestamp
	sec, err = strconv.ParseInt(c.LastAccess, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Proxy timestamp: %v", err)
	}

	proxy.LastAccess = time.Unix(sec, 0)

	proxy.TLSConnect = c.TLSConnect
	proxy.TLSAccept = c.TLSAccept
	proxy.TLSIssuer = c.TLSIssuer
	proxy.TLSSubject = c.TLSSubject
	proxy.TLSPSKIdentity = c.TLSPSKIdentity
	proxy.TLSPSK = c.TLSPSK
	proxy.ProxyAddress = c.ProxyAddress
	proxy.AutoCompress = c.AutoCompress

	return proxy, nil
}

type jProxies []jProxy

// Proxies returns a native Go slice of Proxies mapped from the given JSON
// Proxies data.
func (c jProxies) Proxies() ([]Proxy, error) {
	if c != nil {
		proxies := make([]Proxy, len(c))
		for i, jproxy := range c {
			proxy, err := jproxy.Proxy()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Proxy %d in JSON data: %v", i, err)
			}

			proxies[i] = *proxy

			return proxies, nil
		}
	}

	return nil, nil
}
