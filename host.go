package zabbix

import (
	"errors"
	"fmt"
)

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

	SelectGroups SelectQuery `json:"selectGroups,omitempty"`

	SelectApplications SelectQuery `json:"selectApplications,omitempty"`

	SelectDiscoveries SelectQuery `json:"selectDiscoveries,omitempty"`

	SelectDiscoveryRule SelectQuery `json:"selectDiscoveryRule,omitempty"`

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

var ErrHostNotFound = errors.New("No Host was found matching the given search criteria")

func (c *Session) GetHosts(params HostGetParams) (*[]Host, error) {
	req := NewRequest("host.get", params)
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting hosts: %v", err)
	}

	hosts := make([]Host, 0)
	err = resp.Bind(&hosts)
	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON data for hosts %v", err)
	}

	if len(hosts) == 0 {
		return nil, ErrHostNotFound
	}

	return &hosts, nil
}
