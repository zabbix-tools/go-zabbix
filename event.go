package zabbix

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	// EventSourceTrigger indicates that an Event was created by a Trigger.
	EventSourceTrigger = iota

	// EventSourceDiscoveryRule indicates than an Event was created by a
	// Discovery Rule.
	EventSourceDiscoveryRule

	// EventSourceAutoRegistration indicates that an Event was created by an
	// active Host registration rule.
	EventSourceAutoRegistration

	// EventSourceInternal indicates that an Event was created by an Internal
	// Event.
	EventSourceInternal
)

const (
	// EventObjectTypeTrigger indicates that an Event with Source type
	// EventSourceTrigger or EventSourceInternal is related to a Trigger.
	EventObjectTypeTrigger = iota

	// EventObjectTypeDiscoveredHost indicates that an Event with Source type
	// EventSourceDiscoveryRule is related to a discovered Host.
	EventObjectTypeDiscoveredHost

	// EventObjectTypeDiscoveredService indicates that an Event with Source type
	// EventSourceDiscoveryRule is related to a discovered Service.
	EventObjectTypeDiscoveredService

	// EventObjectTypeAutoRegisteredHost indicates that an Event with Source
	// type EventSourceAutoRegistration is related to an auto-registered Host.
	EventObjectTypeAutoRegisteredHost

	// EventObjectTypeItem indicates that an Event with Source type
	// EventSourceInternal is related to an Item.
	EventObjectTypeItem

	// EventObjectTypeItem indicates that an Event with Source type
	// EventSourceInternal is related to a low-level Discovery Rule.
	EventObjectTypeLLDRule
)

const (
	// TriggerEventValueOK indicates that the Object related to an Event with
	// Source type EventSourceTrigger is in an "OK" state.
	TriggerEventValueOK = iota

	// TriggerEventValueProblem indicates that the Object related to an Event with
	// Source type EventSourceTrigger is in a "Problem" state.
	TriggerEventValueProblem
)

const (
	// DiscoveryEventValueUp indicates that the Host or Service related to an
	// Event with Source type EventSourceDiscoveryRule is in an "Up" state.
	DiscoveryEventValueUp = iota

	// DiscoveryEventValueDown indicates that the Host or Service related to an
	// Event with Source type EventSourceDiscoveryRule is in a "Down" state.
	DiscoveryEventValueDown

	// DiscoveryEventValueDiscovered indicates that the Host or Service related
	// to an Event with Source type EventSourceDiscoveryRule is in a
	// "Discovered" state.
	DiscoveryEventValueDiscovered

	// DiscoveryEventValueLost indicates that the Host or Service related to an
	// Event with Source type EventSourceDiscoveryRule is in a "Lost" state.
	DiscoveryEventValueLost
)

const (
	// InternalEventValueNormal indicates that the Object related to an Event
	// with Source type EventSourceInternal is in a "Normal" state.
	InternalEventValueNormal = iota

	// InternalEventValueNotSupported indicates that the Object related to an
	// Event with Source type EventSourceInternal is in an "Unknown" or
	// "Not supported" state.
	InternalEventValueNotSupported
)

// jEvent is a private map for the Zabbix API Event object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/event/object
type jEvent struct {
	EventID      string `json:"eventid"`
	Acknowledged string `json:"acknowledged"`
	Clock        string `json:"clock"`
	Nanoseconds  string `json:"ns"`
	ObjectType   string `json:"object"`
	ObjectId     string `json:"objectid"`
	Source       string `json:"source"`
	Value        string `json:"value"`
	ValueChanged string `json:"value_changed"`
}

// Event returns a native Go Event struct mapped from the given JSON Event data.
func (c *jEvent) Event() (*Event, error) {
	event := &Event{}
	event.EventID = c.EventID
	event.Acknowledged = (c.Acknowledged == "1")

	// parse timestamp
	sec, err := strconv.ParseInt(c.Clock, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event timestamp: %v", err)
	}

	nsec, err := strconv.ParseInt(c.Nanoseconds, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event timestamp nanoseconds: %v", err)
	}

	event.Timestamp = time.Unix(sec, nsec)

	event.ObjectType, err = strconv.Atoi(c.ObjectType)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event Object Type: %v", err)
	}

	event.ObjectID, err = strconv.Atoi(c.ObjectId)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event Object ID: %v", err)
	}

	event.Source, err = strconv.Atoi(c.Source)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event Source: %v", err)
	}

	event.Value, err = strconv.Atoi(c.Value)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event Source: %v", err)
	}

	event.ValueChanged = (c.ValueChanged == "1")

	return event, nil
}

// Event represents a Zabbix Event returned from the Zabbix API. Events are
// readonly as they may only be created by the Zabbix server.
//
// See: https://www.zabbix.com/documentation/2.2/manual/config/events
type Event struct {
	// EventID is the ID of the Event.
	EventID string

	// Acknowledged indicates if the Event has been acknowledged by an operator.
	Acknowledged bool

	// Timestamp is the UTC timestamp at which the Event occurred.
	Timestamp time.Time

	// Source is the type of the Event source. Source must be one of the
	// EventSource constants.
	Source int

	// ObjectType is the type of the Object that is related to the Event.
	// ObjectType must be one of the EventObjectType constants.
	ObjectType int

	// ObjectID is the unique identifier of the Object that caused this Event.
	ObjectID int

	// Value is the state of the related Object. Value must be one of the
	// EventValue constants, according to the Event's Source type.
	Value int

	// ValueChanges indicates if the state of the related Object has changed
	// since the previous Event.
	ValueChanged bool
}

type EventGetParams struct {
	// EventIDs filters search results to Events that matched the given Event
	// IDs.
	EventIDs []string `json:"eventids,omitempty"`

	// GroupIDs filters search results to events for hosts that are members of
	// the given Group IDs.
	GroupIDs []string `json:"groupids,omitempty"`

	// HostIDs filters search results to events for hosts that matched the given
	// Host IDs.
	HostIDs []string `json:"hostids,omitempty"`

	// ObjectIDs filters search results to events for Objects that matched
	// the given Object IDs.
	ObjectIDs []string `json:"objectids,omitempty"`

	// ObjectType filters search results to events created by the given Object
	// Type. Must be one of the EventObjectType constants.
	//
	// Default: EventObjectTypeTrigger
	ObjectType int `json:"object"`

	// AcknowledgedOnly filters search results to event which have been
	// acknowledged.
	AcknowledgedOnly bool `json:"acknowledged"`

	// MinEventID filters search results to Events with an ID greater or equal
	// to the given ID.
	MinEventID string `json:"eventid_from,omitempty"`

	// MaxEventID filters search results to Events with an ID lesser or equal
	// to the given ID.
	MaxEventID string `json:"eventid_till,omitempty"`

	// Value filters search results to Events with the given values. Each value
	// must be one of the EventValue constants for the given ObjectType.
	Value []int `json:"value,omitempty"`

	// SelectHosts causes all Hosts which contain the object that caused each
	// Event to be attached in the search results.
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`

	// SelectRelatedObject causes the object which caused each Event to be
	// attached in the search results.
	SelectRelatedObject SelectQuery `json:"selectRelatedObject,omitempty"`

	// SelectAlerts causes Alerts generated by each Event to be attached in the
	// search results.
	SelectAlerts SelectQuery `json:"select_alerts,omitempty"`

	// SelectAcknowledgements causes Acknowledgments for each Event to be
	// attached in the search results in reverse chronological order.
	SelectAcknowledgements SelectQuery `json:"select_acknowledges,omitempty"`
}

// ErrEventNotFound describes an empty result set for an Event search.
var ErrEventNotFound = errors.New("No Events were found matching the given search criteria")

// GetEvents queries the Zabbix API for Events matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetEvents(params EventGetParams) (*[]Event, error) {
	req := NewRequest("event.get", params)
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting events: %v", err)
	}

	events := make([]jEvent, 0)
	err = resp.Bind(&events)
	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON data for events %v", err)
	}

	if len(events) == 0 {
		return nil, ErrEventNotFound
	}

	// map JSON Events to Go Events
	out := make([]Event, len(events))
	for i, jevent := range events {
		event, err := jevent.Event()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Event %d in response: %v", err)
		}

		out[i] = *event
	}

	return &out, nil
}
