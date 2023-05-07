package zabbix

import "fmt"

const (
	// HostinterfaceMainNotDefault indicates that the interface is not used
	// as default on the host.
	HostinterfaceMainNotDefault = 0

	// HostinterfaceMainDefault indicates that the interface is used as
	// default on the host.
	HostinterfaceMainDefault = 1
)

const (
	// HostinterfaceTypeDefault is possible to returned value
	HostinterfaceTypeDefault = 0

	// HostinterfaceTypeAgent indicates that the interface type is agent.
	HostinterfaceTypeAgent = 1

	// HostinterfaceTypeSNMP indicates that the interface type is SNMP.
	HostinterfaceTypeSNMP = 2

	// HostinterfaceTypeIPMI indicates that the interface type is SNMP.
	HostinterfaceTypeIPMI = 3

	// HostinterfaceTypeJMX indicates that the interface type is SNMP.
	HostinterfaceTypeJMX = 4
)

const (
	// HostinterfaceUseipDNS indicates that connection using host DNS name.
	HostinterfaceUseipDNS = 0

	// HostinterfaceUseipAddress indeicates that connection using host IP
	// address.
	HostinterfaceUseipAddress = 1
)

// Hostinterface represents a Zabbix Hostinterface returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/hostinterface/object
type Hostinterface struct {
	// InterfaceID is ID of the interface.
	InterfaceID string

	// DNS is DNS name used by the interface.
	DNS string

	// HostID is ID of the host the interface belongs to.
	HostID string

	// IP is IP address used by the interface.
	IP string

	// Main shows that the interface is used as default on the host.
	Main int

	// Port is port number used by the interface. Can contain user macros.
	Port string

	// Type is interface type.
	Type int

	// Useip shows that the connection using host DNS name or IP address.
	Useip int
}

// HostinterfaceGetParams represent the parameters for a `hostinterface.get` API call.
//
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/hostinterface/get#parameters
type HostinterfaceGetParams struct {
	GetParameters

	// HostIDs filters search result to hostinterfaces that matched the
	// given Host IDs.
	HostIDs []string `json:"hostids,omitempty"`

	// InterfaceIDs filters search result to hostinterfaces that matched the
	// given Interface IDs.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// ItemIDs filters search result to hostinterfaces that matched the
	// given Item IDs.
	ItemIDs []string `json:"itemids,omitempty"`

	// TriggerIDs filters search result to hostinterfaces that matched the
	// given Trigger IDs.
	TriggerIDs []string `json:"triggerids,omitempty"`

	SelectItems SelectQuery `json:"selectItems,omitempty"`

	SelectHosts SelectQuery `json:"selectHosts,omitempty"`
}

// GetHostinterfaces queries the Zabbix API for Host Interfaces matching the
// given search parameter.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetHostinterfaces(params HostinterfaceGetParams) ([]Hostinterface, error) {
	hostinterfaces := make([]jHostinterface, 0)
	err := c.Get("hostinterface.get", params, &hostinterfaces)
	if err != nil {
		return nil, err
	}

	if len(hostinterfaces) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Events to Go Events
	out := make([]Hostinterface, len(hostinterfaces))
	for i, jhostinterface := range hostinterfaces {
		hostinterface, err := jhostinterface.Hostinterface()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Hostinterface %d in response: %v", i, err)
		}

		out[i] = *hostinterface
	}

	return out, nil
}
