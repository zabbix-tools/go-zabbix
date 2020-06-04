package zabbix

import (
	"fmt"
)

// History represents a Zabbix History returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/history/object
type History struct {
	// Clock is the time when that value was received.
	Clock int

	// ItemID is the ID of the related item.
	ItemID int

	// Ns is the nanoseconds when the value was received.
	Ns int

	// Value is the received value.
	// Possible types: 0 - float; 1 - character; 2 - log; 3 - int; 4 - text;
	Value string

	// LogEventID is the Windows event log entry ID.
	LogEventID int

	// Severity is the Windows event log entry level.
	Severity int

	// Source is the Windows event log entry source.
	Source string

	// Timestamp is the Windows event log entry time.
	Timestamp string
}

type HistoryGetParams struct {
	GetParameters

	// History object types to return
	// Possible values: 0 - numeric float, 1 - character, 2 - log,
	// 3 - numeric signed, 4, text
	// Default: 3
	History int `json:"history"`

	// HistoryIDs filters search results to histories with the given History ID's.
	HistoryIDs []string `json:"historyids,omitempty"`

	// ItemIDs filters search results to histories belong to the hosts
	// of the given Item ID's.
	ItemIDs []string `json:"itemids,omitempty"`

	// Return only values that have been received after or at the given time.
	TimeFrom float64 `json:"time_from,omitempty"`

	// Return only values that have been received before or at the given time.
	TimeTill float64 `json:"time_till,omitempty"`
}

// GetHistories queries the Zabbix API for Histories matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetHistories(params HistoryGetParams) ([]History, error) {
	histories := make([]jHistory, 0)
	err := c.Get("history.get", params, &histories)
	if err != nil {
		return nil, err
	}
	if len(histories) == 0 {
		return nil, ErrNotFound
	}
	// map JSON Events to Go Events
	out := make([]History, len(histories))
	for i, jhistory := range histories {
		history, err := jhistory.History()
		if err != nil {
			return nil, fmt.Errorf("Error mapping History %d in response: %v", i, err)
		}
		out[i] = *history
	}

	return out, nil
}
