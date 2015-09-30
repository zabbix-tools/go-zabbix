package zabbix

type Host struct {
	HostID      string `json:"hostid"`
	Hostname    string `json:"host"`
	DisplayName string `json:"name"`

	Macros []HostMacro `json:"macros,omitempty"`
}

// HostGetParams represent the parameters for a `host.get` API call.
//
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/host/get#parameters
type HostGetParams struct {
	// GroupIDs filters search results to hosts that are members of the given
	// Group IDs.
	GroupIDs []string `json:"groupids,omitempty"`

	// ApplicationIDs filters search results to hosts that have items in the
	// given Application IDs.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// DiscoveredServiceIDs filters search results to hosts that are related to
	// the given discovered service IDs.
	DiscoveredServiceIDs []string `json:"dserviceids,omitempty"`

	// GraphIDs filters search results to hosts that have the given graph IDs.
	GraphIDs []string `json:"graphids,omitempty"`

	// HostIDs filters search results to hosts that matched the given Host IDs.
	HostIDs []string `json:"hostids,omitempty"`

	// WebCheckIDs filters search results to hosts with the given Web Check IDs.
	WebCheckIDs []string `json:"httptestids,omitempty"`

	// InterfaceIDs filters search results to hosts that use the given Interface
	// IDs.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// ItemIDs filters search results to hosts with the given Item IDs.
	ItemIDs []string `json:"itemids,omitempty"`

	// MaintenanceIDs filters search results to hosts that are affected by the
	// given Maintenance IDs
	MaintenanceIDs []string `json:"maintenanceids,omitempty"`

	// MonitoredOnly filters search results to return only monitored hosts.
	MonitoredOnly bool `json:"monitored_hosts,omitempty"`

	// ProxyOnly filters search results to hosts which are Zabbix proxies.
	ProxiesOnly bool `json:"proxy_host,omitempty"`

	// ProxyIDs filters search results to hosts monitored by the given Proxy
	// IDs.
	ProxyIDs []string `json:"proxyids,omitempty"`

	// IncludeTemplates extends search results to include Templates.
	IncludeTemplates bool `json:"templated_hosts,omitempty"`

	// SelectGroups causes the Host Groups that each Host belongs to to be
	// attached in the search results.
	SelectGroups SelectQuery `json:"selectGroups,omitempty"`

	// SelectApplications causes the Applications from each Host to be attached
	// in the search results.
	SelectApplications SelectQuery `json:"selectApplications,omitempty"`

	// SelectDiscoveries causes the Low-Level Discoveries from each Host to be
	// attached in the search results.
	SelectDiscoveries SelectQuery `json:"selectDiscoveries,omitempty"`

	// SelectDiscoveryRule causes the Low-Level Discovery Rule that created each
	// Host to be attached in the search results.
	SelectDiscoveryRule SelectQuery `json:"selectDiscoveryRule,omitempty"`

	// SelectGraphs causes the Graphs from each Host to be attached in the
	// search results.
	SelectGraphs SelectQuery `json:"selectGraphs,omitempty"`

	SelectHostDiscovery SelectQuery `json:"selectHostDiscovery,omitempty"`

	SelectWebScenarios SelectQuery `json:"selectHttpTests,omitempty"`

	SelectInterfaces SelectQuery `json:"selectInterfaces,omitempty"`

	SelectInventory SelectQuery `json:"selectInventory,omitmepty"`

	SelectItems SelectQuery `json:"selectItems,omitempty"`

	SelectMacros SelectQuery `json:"selectMacros,omitempty"`

	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	SelectScreens         SelectQuery `json:"selectScreens,omitempty"`
	SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"`
}

func (c *Session) GetHosts(params HostGetParams) (*[]Host, error) {
	hosts := make([]Host, 0)
	err := c.Get("host.get", params, &hosts)
	if err != nil {
		return nil, err
	}

	if len(hosts) == 0 {
		return nil, ErrNotFound
	}

	return &hosts, nil
}
