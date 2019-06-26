package zabbix

import (
	"fmt"
	"strconv"
)

// jHistory is a private map for the Zabbix API History object.
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/history/get
type jHistory struct {
	ItemID     string `json:"itemid"`
	Clock      string `json:"clock"`
	Ns         string `json:"ns"`
	Value      string `json:"value"`
	LogEventID string `json:"logeventid,omitempty"`
	Severity   string `json:"severity,omitempty"`
	Source     string `json:"source,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
}

// History returns a native Go History struct mapped from the given JSON History data.
func (c *jHistory) History() (*History, error) {
	var err error
	history := &History{}

	history.Clock, err = strconv.Atoi(c.Clock)
	if err != nil {
		return nil, fmt.Errorf("Error parsing History Clock: %v", err)
	}

	history.ItemID, err = strconv.Atoi(c.ItemID)
	if err != nil {
		return nil, fmt.Errorf("Error parsing History ItemID: %v", err)
	}

	history.Ns, err = strconv.Atoi(c.Ns)
	if err != nil {
		return nil, fmt.Errorf("Error parsing History Ns: %v", err)
	}

	history.Value = c.Value

	if c.LogEventID != "" {
		history.LogEventID, err = strconv.Atoi(c.LogEventID)
		if err != nil {
			return nil, fmt.Errorf("Error parsing History LogEventID: %v", err)
		}
	}

	if c.Severity != "" {
		history.LogEventID, err = strconv.Atoi(c.Severity)
		if err != nil {
			return nil, fmt.Errorf("Error parsing History Severity: %v", err)
		}
	}

	history.Source = c.Source

	history.Timestamp = c.Timestamp

	return history, err
}

// jHistories is a slice of jHistory structs.
type jHistories []jHistory

// Histories returns a native Go slice of Histories mapped from the given JSON HISTORIES
// data.
func (c jHistories) Histories() ([]History, error) {
	if c != nil {
		histories := make([]History, len(c))
		for i, jhistory := range c {
			history, err := jhistory.History()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling History %d in JSON data: %v", i, err)
			}
			histories[i] = *history
		}

		return histories, nil
	}

	return nil, nil
}
