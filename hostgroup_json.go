package zabbix

import (
	"fmt"
)

// jHostgroup is a private map for the Hostgroup Zabbix API object (see zabbix documentation).
type jHostgroup struct {
	GroupID  string `json:"groupid"`
	Name     string `json:"name"`
	Flags    string `json:"flags"`
	Internal string `json:"internal"`
	Hosts    jHosts `json:"hosts,omitempty"`
}

// Hostgroup returns a native Go Hostgroup struct mapped from the given JSON Hostgroup data.
func (c *jHostgroup) Hostgroup() (*Hostgroup, error) {
	hostgroup := &Hostgroup{}
	hostgroup.GroupID = c.GroupID
	hostgroup.Name = c.Name
	hostgroup.Flags = c.Flags
	hostgroup.Internal = c.Internal

	if len(c.Hosts) > 0 {
		if hosts, err := c.Hosts.Hosts(); err == nil {
			hostgroup.Hosts = hosts
		}

	}

	return hostgroup, nil
}

// jHostgroups is a slice of jHostgroup structs.
type jHostgroups []jHostgroup

// Hostgroups returns a native Go slice of Hostgroups mapped from the given JSON Hostgroups
// data.
func (c jHostgroups) Hostgroups() ([]Hostgroup, error) {
	if c != nil {
		hosts := make([]Hostgroup, len(c))
		for i, jhost := range c {
			host, err := jhost.Hostgroup()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Hostgroup %d in JSON data: %v", i, err)
			}

			hosts[i] = *host
		}

		return hosts, nil
	}

	return nil, nil
}
