package zabbix

import (
	"fmt"
	//"strconv"
)

// jHost is a private map for the Zabbix API Host object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/host/object
type jHost struct {
	HostID   string      `json:"hostid"`
	Hostname string      `json:"host"`
	Flags    int         `json:"flags,string,omitempty"`
	Name     string      `json:"name,omitempty"`
	Macros   []HostMacro `json:"macros,omitempty"`
	Groups   []Hostgroup `json:"groups,omitempty"`
}

// Host returns a native Go Host struct mapped from the given JSON Host data.
func (c *jHost) Host() (*Host, error) {
	//var err error

	host := &Host{}
	host.HostID = c.HostID
	host.Hostname = c.Hostname
	host.DisplayName = c.Name
	host.Macros = c.Macros
	host.Groups = c.Groups
	/*
		host.Source, err = strconv.Atoi(c.Flags)
		if err != nil {
			return nil, fmt.Errorf("Error parsing Host Flags: %v", err)
		}
	*/
	host.Source = c.Flags
	return host, nil
}

// jHosts is a slice of jHost structs.
type jHosts []jHost

// Hosts returns a native Go slice of Hosts mapped from the given JSON Hosts
// data.
func (c jHosts) Hosts() ([]Host, error) {
	if c != nil {
		hosts := make([]Host, len(c))
		for i, jhost := range c {
			host, err := jhost.Host()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Host %d in JSON data: %v", i, err)
			}

			hosts[i] = *host
		}

		return hosts, nil
	}

	return nil, nil
}
