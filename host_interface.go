package zabbix

const (
	// HostInterfaceAvailabilityUnknown Unknown availability of host, never has come online
	HostInterfaceAvailabilityUnknown = 0
	// HostInterfaceAvailabilityAvailable Host is available
	HostInterfaceAvailabilityAvailable = 1
	// HostInterfaceAvailabilityUnavailable Host is NOT available
	HostInterfaceAvailabilityUnavailable = 2

	// HostInterfaceTypeAgent Host interface type agent
	HostInterfaceTypeAgent = 1
	// HostInterfaceTypeSNMP Host interface type SNMP
	HostInterfaceTypeSNMP = 2
	// HostInterfaceTypeIPMI Host interface type IPMI
	HostInterfaceTypeIPMI = 3
	// HostInterfaceTypeJMX Host interface type JMX
	HostInterfaceTypeJMX = 4
)

// HostInterface  This class is designed to work with host interfaces.
//
// See https://www.zabbix.com/documentation/current/manual/api/reference/hostinterface/object#host_interface
type HostInterface struct {
	// (readonly) ID of the interface.
	InterfaceID string `json:"interfaceid"`

	// (readonly) Availability of host interface.
	Available int `json:"available,string,omitempty"`

	// DNS name used by the interface.
	DNS string `json:"dns"`

	// IP address used by the interface.
	IP string `json:"ip"`

	// (readonly) Error text if host interface is unavailable.
	Error string `json:"error,omitempty"`

	// (readonly) Time when host interface became unavailable.
	ErrorsFrom *UnixTimestamp `json:"errors_from,string,omitempty"`

	// ID of the host the interface belongs to.
	HostID string `json:"hostid"`

	// Whether the interface is used as default on the host. Only one interface of some type can be set as default on a host.
	Main ZBXBoolean `json:"main,string"`

	// Interface type.
	Type int `json:"type,string"`

	// Whether the connection should be made via IP.
	UseIP ZBXBoolean `json:"useip,string"`
}

type HostInterfaceGetParams struct {
	GetParameters

	// Return only host interfaces used by the given hosts.
	HostIDs []string `json:"hostids,omitempty"`

	// Return only host interfaces with the given IDs.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// Return only host interfaces used by the given items.
	ItemIDs []string `json:"itemids,omitempty"`

	// Return only host interfaces used by items in the given triggers.
	TriggerIDs []string `json:"triggerids,omitempty"`
}

// GetHostInterfaces queries the Zabbix API for Hosts interfaces matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetHostInterfaces(params HostInterfaceGetParams) ([]HostInterface, error) {
	hostInterfaces := make([]HostInterface, 0)
	err := c.Get("hostinterface.get", params, &hostInterfaces)
	if err != nil {
		return nil, err
	}

	if len(hostInterfaces) == 0 {
		return nil, ErrNotFound
	}

	return hostInterfaces, nil
}
