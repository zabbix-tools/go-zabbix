package zabbix

import (
	"fmt"
)

// Item represents a Zabbix Item returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/item/object
type Item struct {
	// HostID is the unique ID of the Host.
	HostID int

	// ItemID is the unique ID of the Item.
	ItemID int

	// Itemname is the technical name of the Item.
	ItemName string

	// ItemDescr is the description of the Item.
	ItemDescr string

	// LastClock is the last Item epoh time.
	LastClock int

	// LastValue is the last value of the Item.
	LastValue string

	// LastValueType is the type of LastValue
	// 0 - float; 1 - text; 3 - int;
	LastValueType int
}

type ItemTagFilter struct {
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Operator int    `json:"operator"`
}

type ItemGetParams struct {
	GetParameters

	// ItemIDs filters search results to items with the given Item ID's.
	ItemIDs []string `json:"itemids,omitempty"`

	// GroupIDs filters search results to items belong to the hosts
	// of the given Group ID's.
	GroupIDs []string `json:"groupids,omitempty"`

	// TemplateIDs filters search results to items belong to the
	// given templates of the given Template ID's.
	TemplateIDs []string `json:"templateids,omitempty"`

	// HostIDs filters search results to items belong to the
	// given Host ID's.
	HostIDs []string `json:"hostids,omitempty"`

	// ProxyIDs filters search results to items that are
	// monitored by the given Proxy ID's.
	ProxyIDs []string `json:"proxyids,omitempty"`

	// InterfaceIDs filters search results to items that use
	// the given host Interface ID's.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// GraphIDs filters search results to items that are used
	// in the given graph ID's.
	GraphIDs []string `json:"graphids,omitempty"`

	// TriggerIDs filters search results to items that are used
	// in the given Trigger ID's.
	TriggerIDs []string `json:"triggerids,omitempty"`

	// ApplicationIDs filters search results to items that
	// belong to the given Applications ID's.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// WebItems flag includes web items in the result.
	WebItems bool `json:"webitems,omitempty"`

	// Inherited flag return only items inherited from a template
	// if set to 'true'.
	Inherited bool `json:"inherited,omitempty"`

	// Templated flag return only items that belong to templates
	// if set to 'true'.
	Templated bool `json:"templated,omitempty"`

	// Monitored flag return only enabled items that belong to
	// monitored hosts if set to 'true'.
	Monitored bool `json:"monitored,omitempty"`

	// Group filters search results to items belong to a group
	// with the given name.
	Group string `json:"group,omitempty"`

	// Host filters search results to items that belong to a host
	// with the given name.
	Host string `json:"host,omitempty"`

	// Application filters search results to items that belong to
	// an application with the given name.
	Application string `json:"application,omitempty"`

	// WithTriggers flag return only items that are used in triggers
	WithTriggers bool `json:"with_triggers,omitempty"`

	// Filter by tags
	Tags []ItemTagFilter `json:"tags,omitempty"`
}

// GetItems queries the Zabbix API for Items matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetItems(params ItemGetParams) ([]Item, error) {
	items := make([]jItem, 0)
	err := c.Get("item.get", params, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, ErrNotFound
	}
	// map JSON Events to Go Events
	out := make([]Item, len(items))
	for i, jitem := range items {
		item, err := jitem.Item()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Item %d in response: %v", i, err)
		}
		out[i] = *item
	}

	return out, nil
}
