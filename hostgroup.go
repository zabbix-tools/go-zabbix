package zabbix

import (
	"fmt"
)

const (
	// HostgroupSourcePlain indicates that a Hostgroup was created in the normal way.
	HostgroupSourcePlain = 0

	// HostgroupSourceDiscovery indicates that a Hostgroup was created by Host discovery.
	HostgroupSourceDiscovery = 4

	// HostgroupInternalNo indicates that a Hostgroup is used not internally by the system.
	HostgroupInternalNo = 0

	// HostgroupInternalYes indicates that a Hostgroup is used internally by the system.
	HostgroupInternalYes = 1
)

// Hostgroup represents a Zabbix Hostgroup Object returned from the Zabbix API (see zabbix documentation).
type Hostgroup struct {
	GroupID  string `json:"groupid"`
	Name     string `json:"name"`
	Flags    string `json:"flags"`
	Internal string `json:"internal"`
	Hosts    []Host `json:"hosts,omitempty"`
}

// HostgroupGetParams represent the parameters for a `hostgroup.get` API call (see zabbix documentation).
type HostgroupGetParams struct {
	GetParameters

	// Return only host groups that contain hosts or templates with the given graphs
	GraphIDs []string `json:"graphids,omitempty"`

	// Return only host groups with the given host group IDs
	GroupIDs []string `json:"groupids,omitempty"`

	// Return only host groups that contain the given hosts
	HostIDs []string `json:"hostids,omitempty"`

	// Return only host groups that are affected by the given maintenances
	MaintenanceIDs []string `json:"maintenanceids,omitempty"`

	// Return only host groups that contain monitored hosts
	MonitoredHosts int `json:"monitored_hosts,omitempty"`

	// Return only host groups that contain hosts
	RealHosts int `json:"real_hosts,omitempty"`

	// Return only host groups that contain templates
	TemplatedHosts int `json:"templated_hosts,omitempty"`

	// Return only host groups that contain the given templates
	TemplateIDs []string `json:"templateids,omitempty"`

	// Return only host groups that contain hosts or templates with the given triggers
	TriggerIDs []string `json:"triggerids,omitempty"`

	// Return only host groups that contain hosts with applications
	WithApplications int `json:"with_applications,omitempty"`

	// Return only host groups that contain hosts with graphs
	WithGraphs int `json:"with_graphs,omitempty"`

	// Return only host groups that contain hosts or templates
	WithHostsAndTemplates int `json:"with_hosts_and_templates,omitempty"`

	// Return only host groups that contain hosts with web checks
	WithHttptests int `json:"with_httptests,omitempty"`

	// Return only host groups that contain hosts or templates with items
	WithItems int `json:"with_items,omitempty"`

	// Return only host groups that contain hosts with enabled web checks
	WithMonitoredHttptests int `json:"with_monitored_httptests,omitempty"`

	// Return only host groups that contain hosts or templates with enabled items
	WithMonitoredItems int `json:"with_monitored_items,omitempty"`

	// Return only host groups that contain hosts with enabled triggers
	WithMonitoredTriggers int `json:"with_monitored_triggers,omitempty"`

	// Return only host groups that contain hosts with numeric items
	WithSimpleGraphItems int `json:"with_simple_graph_items,omitempty"`

	// Return only host groups that contain hosts with triggers
	WithTriggers int `json:"with_triggers,omitempty"`

	// Return the LLD rule that created the host group in the discoveryRule property
	SelectDiscoveryRule SelectQuery `json:"selectDiscoveryRule,omitempty"`

	// Return the host group discovery object in the groupDiscovery property
	SelectGroupDiscovery SelectQuery `json:"selectGroupDiscovery,omitempty"`

	// Return the hosts that belong to the host group in the hosts property
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`

	// Return the templates that belong to the host group in the templates property
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`

	// Limits the number of records returned by subselects
	LimitSelects int `json:"limitSelects,omitempty"`

	// Return only host groups that contain the given templates
	Sortfield []string `json:"sortfield,omitempty"`
}

// GetHostgroups queries the Zabbix API for Hostgroups matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetHostgroups(params HostgroupGetParams) ([]Hostgroup, error) {
	hostgroups := make([]jHostgroup, 0)
	err := c.Get("hostgroup.get", params, &hostgroups)
	if err != nil {
		return nil, err
	}

	if len(hostgroups) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Events to Go Events
	out := make([]Hostgroup, len(hostgroups))
	for i, jhostgroup := range hostgroups {
		hostgroup, err := jhostgroup.Hostgroup()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Hostgroup %d in response: %v", i, err)
		}

		out[i] = *hostgroup
	}

	return out, nil
}
