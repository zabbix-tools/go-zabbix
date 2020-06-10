package zabbix

import (
	"fmt"
	"time"
)

const (
	// ProxyStatusActiveProxy indicates that a Proxy type is active.
	ProxyStatusActiveProxy = 5

	// ProxyStatusPassiveProxy indicates that a Proxy type is passive.
	ProxyStatusPassiveProxy = 6
)

const (
	// ProxyTLSConnectDefault indicates that unencrypted connections to host.
	ProxyTLSConnectDefault = 1

	// ProxyTLSConnectPSK indicates that PSK encryption connections to host.
	ProxyTLSConnectPSK = 2

	// ProxyTLSConnectCert indicates that certificate encryption connections to host.
	ProxyTLSConnectCert = 4
)

const (
	// ProxyTLSAcceptDefault indicates that unencrypted connections from host.
	ProxyTLSAcceptDefault = iota

	// ProxyTLSAcceptPSK indicates that PSK encryption connections from host.
	ProxyTLSAcceptPSK

	// ProxyTLSAcceptCert indicates that certificate connections from host.
	ProxyTLSAcceptCert
)

const (
	// ProxyAutoCompressDisabled indicates that communication is not
	// compressed between Zabbix server and proxy.
	ProxyAutoCompressDisabled = 0

	// ProxyAutoCompressEnabled indicates that communication is compressed
	// between Zabbix server and proxy.
	ProxyAutoCompressEnabled = 1
)

// Proxy represents a Zabbix Proxy returned from the Zabbix API
//
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/proxy/object
type Proxy struct {
	// ProxyID is the unique ID of the Proxy.
	ProxyID string

	// Hostname is the name of Proxy.
	Hostname string

	// Status is type of Proxy.
	Status int

	// Description of the Proxy.
	Description string

	// LastAccess is the UTC timestamp at which the Proxy last connected
	// to the server.
	LastAccess time.Time

	// TLSConnect shows connection to host.
	TLSConnect int

	// TLSAccept shows connection from host.
	TLSAccept int

	// TLSIssuer is certificate issuer.
	TLSIssuer string

	// TLSSubject is certificate subject.
	TLSSubject string

	// TLSPSKIdentity is PSK identity.
	TLSPSKIdentity string

	// TLSPSK is the preshared key.
	TLSPSK string

	// ProxyAddress is IP address or DNS names of active proxy.
	ProxyAddress string

	// AutoCompress indicates if communication between Zabbix server and proxy
	// is compressed.
	AutoCompress int
}

// ProxyGetParams represent the parameters for a ``proxy.get API call.
//
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/proxy/get#parameters
type ProxyGetParams struct {
	GetParameters

	// ProxyIDs filters search results to proxies that are members of the given
	// Proxy IDs.
	ProxyIDs []string `json:"proxyids,omitempty"`

	SelectHosts SelectQuery `json:"selectHosts"`

	SelectInterface SelectQuery `json:"selectInterface"`
}

// GetProxies queries the Zabbix API for Proxy matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetProxies(params ProxyGetParams) ([]Proxy, error) {
	proxies := make([]jProxy, 0)
	err := c.Get("proxy.get", params, &proxies)
	if err != nil {
		return nil, err
	}

	if len(proxies) == 0 {
		return nil, ErrNotFound
	}

	out := make([]Proxy, len(proxies))
	for i, jproxy := range proxies {
		proxy, err := jproxy.Proxy()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Proxy %d in response: %v", i, err)
		}

		out[i] = *proxy
	}

	return out, nil
}
