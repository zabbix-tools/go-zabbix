package zabbix

import (
	"time"
)

// jAlert is a private map for the Zabbix API Alert object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/alert/object
type jAlert struct {
	AlertID     string `json:"alertid"`
	ActionID    string `json:"actionid"`
	AlertType   int    `json:"alerttype,string"`
	Clock       int64  `json:"clock,string"`
	Error       string `json:"error"`
	EscStep     int    `json:"esc_step,string"`
	EventID     string `json:"eventid"`
	MediaTypeID string `json:"mediatypeid"`
	Message     string `json:"message"`
	Retries     int    `json:"retries,string"`
	SendTo      string `json:"sendto"`
	Status      int    `json:"status,string"`
	Subject     string `json:"subject"`
	UserID      string `json:"userid"`
	Hosts       jHosts `json:"hosts"`
}

// Alert returns a native Go Alert struct mapped from the given JSON Alert data.
func (c *jAlert) Alert() (*Alert, error) {
	var err error

	alert := &Alert{}
	alert.AlertID = c.AlertID
	alert.ActionID = c.ActionID
	alert.AlertType = c.AlertType
	alert.Timestamp = time.Unix(c.Clock, 0)
	alert.ErrorText = c.Error
	alert.EscalationStep = c.EscStep
	alert.EventID = c.EventID
	alert.MediaTypeID = c.MediaTypeID
	alert.Message = c.Message
	alert.RetryCount = c.Retries
	alert.Recipient = c.SendTo
	alert.Status = c.Status
	alert.Subject = c.Subject
	alert.UserID = c.UserID

	// map Hosts
	alert.Hosts, err = c.Hosts.Hosts()
	if err != nil {
		return nil, err
	}

	return alert, nil
}
