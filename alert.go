package zabbix

import (
	"fmt"
	"time"
)

const (
	// AlertTypeMessage indicates that an Alert is a notification message.
	AlertTypeMessage = iota

	// AlertTypeRemoteCommand indicates that an Alert is a remote command call.
	AlertTypeRemoteCommand
)

const (
	// AlertMessageStatusNotSent indicates that an Alert of type
	// AlertTypeMessage has not been sent yet.
	AlertMessageStatusNotSent = iota

	// AlertMessageStatusSent indicates that an Alert of type AlertTypeMessage
	// has been sent successfully.
	AlertMessageStatusSent

	// AlertMessageStatusFailed indicates that an Alert of type AlertTypeMessage
	// failed to send.
	AlertMessageStatusFailed
)

const (
	// AlertCommandStatusRun indicates that an Alert of type
	// AlertTypeRemoteCommand has been run.
	AlertCommandStatusRun = 1 + iota

	// AlertCommandStatusAgentUnavailable indicates that an Alert of type
	// AlertTypeRemoteCommand failed to run as the Zabbix Agent was unavailable.
	AlertCommandStatusAgentUnavailable
)

// Alert represents a Zabbix Alert returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/2.2/manual/config/notifications
type Alert struct {
	// AlertID is the unique ID of the Alert.
	AlertID string

	// ActionID is the unique ID of the Action that generated this Alert.
	ActionID string

	// AlertType is the type of the Alert.
	// AlertType must be one of the AlertType constants.
	AlertType int

	// Timestamp is the UTC timestamp at which the Alert was generated.
	Timestamp time.Time

	// ErrorText is the error message if there was a problem sending a message
	// or running a remote command.
	ErrorText string

	// EscalationStep is the escalation step during which the Alert was
	// generated.
	EscalationStep int

	// EventID is the unique ID of the Event that triggered this Action that
	// generated this Alert.
	EventID string

	// MediaTypeID is the unique ID of the Media Type that was used to send this
	// Alert if the AlertType is AlertTypeMessage.
	MediaTypeID string

	// Message is the Alert message body if AlertType is AlertTypeMessage.
	Message string

	// RetryCount is the number of times Zabbix tried to send a message.
	RetryCount int

	// Recipient is the end point address of a message if AlertType is
	// AlertTypeMessage.
	Recipient string

	// Status indicates the outcome of executing the Alert.
	//
	// If AlertType is AlertTypeMessage, Status must be one of the
	// AlertMessageStatus constants.
	//
	// If AlertType is AlertTypeRemoteCommand, Status must be one of the
	// AlertCommandStatus constants.
	Status int

	// Subject is the Alert message subject if AlertType is AlertTypeMessage.
	Subject string

	// UserID is the unique ID of the User the Alert message was sent to.
	UserID string

	// Hosts is an array of Hosts that triggered this Alert.
	//
	// Hosts is only populated if AlertGetParams.SelectHosts is given in the
	// query parameters that returned this Alert.
	Hosts []Host
}

// AlertGetParams is query params for alert.get call
type AlertGetParams struct {
	GetParameters

	// SelectHosts causes all Hosts which triggered the Alert to be attached in
	// the search results.
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`

	// SelectMediaTypes causes the Media Types used for the Alert to be attached
	// in the search results.
	SelectMediaTypes SelectQuery `json:"selectMediatypes,omitempty"`

	// SelectUsers causes all Users to which the Alert was addressed to be
	// attached in the search results.
	SelectUsers SelectQuery `json:"selectUsers,omitempty"`
}

// GetAlerts queries the Zabbix API for Alerts matching the given search
// parameters.
//
// ErrNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetAlerts(params AlertGetParams) ([]Alert, error) {
	alerts := make([]jAlert, 0)
	err := c.Get("alert.get", params, &alerts)
	if err != nil {
		return nil, err
	}

	if len(alerts) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Alerts to Go Alerts
	out := make([]Alert, len(alerts))
	for i, jalert := range alerts {
		alert, err := jalert.Alert()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Alert %d in response: %v", i, err)
		}

		out[i] = *alert
	}

	return out, nil
}
