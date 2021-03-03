package zabbix

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type MaintenanceType int
type TagsEvaltype int

var ErrMaintenanceHostNotFound = errors.New("Failed to find ID by host name")

const (
	withDataCollection MaintenanceType = iota
	withoutDataCollection

	and TagsEvaltype = iota * 2
	or

	Once int = iota
	EveryDay
	EveryWeek
	EveryMonth
)

type Maintenance struct {
	MaintenanceID string
	Name          string
	ActiveSince   time.Time
	ActiveTill    time.Time
	Description   string
	// Service period in hours
	ServicePeriod       int
	Type                MaintenanceType
	ActionEvalTypeAndOr TagsEvaltype
}

type MaintenanceGetParams struct {
	GetParameters

	// Sort the result by the given properties.
	// Possible values are: maintenanceid, name and maintenance_type.
	SortField []string `json:"sortfield,omitempty"`

	// Return the maintenance's time periods in the timeperiods property.
	SelectTimeperiods SelectQuery `json:"selectTimeperiods,omitempty"`

	// Return hosts assigned to the maintenance in the hosts property.
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`

	// Return host groups assigned to the maintenance in the groups property.
	SelectGroups SelectQuery `json:"selectGroups,omitempty"`

	// Return only maintenances with the given IDs.
	Maintenanceids []string `json:"maintenanceids,omitempty"`

	// Return only maintenances that are assigned to the given hosts.
	Hostids []string `json:"hostids,omitempty"`

	// Return only maintenances that are assigned to the given host groups.
	Groupids []string `json:"groupids,omitempty"`
}

type MaintenanceCreateParams struct {
	JMaintenance

	Groupids []string `json:"groupids,omitempty"`
	// Hosts name
	HostNames   []string      `json:"-"`
	HostIDs     []string      `json:"hostids"`
	Timeperiods []Timeperiods `json:"timeperiods"`
	Tags        []string      `json:"tags,omitempty"`
}

type MaintenanceCreateResponse struct {
	IDs []string `json:"maintenanceids"`
}

// GetMaintenance queries the Zabbix API for Maintenance matching the given search
// parameters.
func (s *Session) GetMaintenance(params *MaintenanceGetParams) ([]Maintenance, error) {
	jmaintenance := make([]JMaintenance, 0)
	err := s.Get("maintenance.get", params, &jmaintenance)
	if err != nil {
		return nil, err
	}

	if len(jmaintenance) == 0 {
		return nil, ErrNotFound
	}

	out := make([]Maintenance, len(jmaintenance))
	for i, jaction := range jmaintenance {
		maintenance, err := jaction.Maintenance()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Maintenance %d in response: %v", i, err)
		}

		out[i] = *maintenance
	}

	return out, nil
}

func (s *Session) CreateMaintenance(params *MaintenanceCreateParams) (response MaintenanceCreateResponse, err error) {
	if err = params.FillHostIDs(s); err != nil {
		return
	}

	err = s.Get("maintenance.create", params, &response)
	return
}

func (m *Maintenance) Delete(session *Session) error {
	ID := []string{m.MaintenanceID}
	response := make(map[string]interface{})
	if err := session.Get("maintenance.delete", ID, &response); err != nil {
		return err
	}
	return nil
}

func (m *MaintenanceCreateParams) FillHostIDs(session *Session) error {
	hosts, err := session.GetHosts(HostGetParams{})
	if err != nil {
		return err
	}

	err = ErrMaintenanceHostNotFound
	for _, name := range m.HostNames {
		for _, host := range hosts {
			if strings.ToUpper(strings.Trim(host.Hostname, " ")) == strings.ToUpper(strings.Trim(name, " ")) {
				m.HostIDs = append(m.HostIDs, host.HostID)

				err = nil
			}
		}
	}

	return err
}

func (c *MaintenanceCreateParams) FillFields(Object *Maintenance) *MaintenanceCreateParams {
	c.ActiveSince = Object.ActiveSince.Unix()
	c.ActiveTill = Object.ActiveSince.Add(time.Hour * time.Duration(Object.ServicePeriod)).Unix()
	c.Description = Object.Description
	c.MaintenanceID = Object.MaintenanceID
	c.Name = Object.Name
	c.TagsEvaltype = int(Object.Type)
	c.MaintenanceType = int(Object.ActionEvalTypeAndOr)

	return c
}
