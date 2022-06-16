package zabbix

type jProxyInterface struct {
	InterfaceID string `json:"interfaceid,omitempty"`
	HostID      string `json:"hostid,omitempty"`
	UseIP       int    `json:"useip,string"`
	IP          string `json:"ip,omitempty"`
	DNS         string `json:"dns,omitempty"`
	Port        string `json:"port"`
}

func (c *jProxyInterface) ProxyInterface() (*ProxyInterface, error) {

	proxy_int := &ProxyInterface{}
	proxy_int.InterfaceID = c.InterfaceID
	proxy_int.HostID = c.HostID
	proxy_int.UseIP = c.UseIP
	proxy_int.IP = c.IP
	proxy_int.DNS = c.DNS
	proxy_int.Port = c.Port
	return proxy_int, nil
}
