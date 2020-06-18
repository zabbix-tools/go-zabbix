package zabbix

import "fmt"

// jHostinterface is a private map for the Zabbix API Hostinterface object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/hostinterface/object
type jHostinterface struct {
	InterfaceID string `json:"interfaceid"`
	DNS         string `json:"dns"`
	HostID      string `json:"hostid"`
	IP          string `json:"ip"`
	Main        int    `json:"main,string"`
	Port        string `json:"port"`
	Type        int    `json:"type,string"`
	Useip       int    `json:"useip,string"`
}

// Hostinterface returns a native Go Hostinterface struct mapped from the given
// JSON Hostinterface data.
func (c *jHostinterface) Hostinterface() (*Hostinterface, error) {
	hostinterface := &Hostinterface{}
	hostinterface.InterfaceID = c.InterfaceID
	hostinterface.DNS = c.DNS
	hostinterface.HostID = c.HostID
	hostinterface.IP = c.IP
	hostinterface.Main = c.Main
	hostinterface.Port = c.Port
	hostinterface.Type = c.Type
	hostinterface.Useip = c.Useip

	return hostinterface, nil
}

// jHostinterfaces is a slice of jHostinterface structs.
type jHostinterfaces []jHostinterface

func (c jHostinterfaces) Hostinterface() ([]Hostinterface, error) {
	if c != nil {
		hostinterfaces := make([]Hostinterface, len(c))
		for i, jhostinterface := range c {
			hostinterface, err := jhostinterface.Hostinterface()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Hostinterface %d in JSON data: %v", i, err)
			}

			hostinterfaces[i] = *hostinterface
		}

		return hostinterfaces, nil
	}

	return nil, nil
}
