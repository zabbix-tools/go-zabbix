package zabbix

import (
	"fmt"
)

const (
	// ActionEvalTypeAndOr indicated that an Action will evaluate its conditions
	// using AND/OR bitwise logic.
	ActionEvalTypeAndOr = iota

	// ActionEvalTypeAnd indicated that an Action will evaluate its conditions
	// using AND bitwise logic.
	ActionEvalTypeAnd

	// ActionEvalTypeOr indicated that an Action will evaluate its conditions
	// using OR bitwise logic.
	ActionEvalTypeOr
)

// Action represents a Zabbix Action returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/2.2/manual/config/notifications/action
type Action struct {
	// ActionID is the unique ID of the Action.
	ActionID string

	// StepDuration is the interval in seconds between each operation step.
	StepDuration int

	// EvaluationType determines the bitwise logic used to evaluate the Actions
	// conditions.
	//
	// EvaluationType must be one of the ActionEvalType constants.
	//
	// EvaluationType is only supported up to Zabbix v2.2.
	EvaluationType int

	// EventType is the type of Events that this Action will handle.
	//
	// Source must be one of the EventSource constants.
	EventType int

	// Name is the name of the Action.
	Name string

	// ProblemMessageBody is the message body text to be submitted for this
	// Action.
	ProblemMessageBody string

	// ProblemMessageSubject is the short summary text to be submitted for this
	// Action.
	ProblemMessageSubject string

	// RecoveryMessageBody is the message body text to be submitted for this
	// Action.
	RecoveryMessageBody string

	// RecoveryMessageSubject is the short summary text to be submitted for this
	// Action.
	RecoveryMessageSubject string

	// RecoveryMessageEnabled determines whether recovery messages will be
	// submitted for the Action when the source problem is resolved.
	RecoveryMessageEnabled bool

	// Enabled determines whether the Action is enabled or disabled.
	Enabled bool

	// Conditions are the conditions which must be met for this Action to
	// execute.
	Conditions []ActionCondition

	// Operations are the operations which will be exectuted for this Action.
	Operations []ActionOperation
}

// ActionCondition is action condition
type ActionCondition struct{}

// ActionOperation is action operation
type ActionOperation struct{}

// ActionGetParams is query params for action.get call
type ActionGetParams struct {
	GetParameters
}

// GetActions queries the Zabbix API for Actions matching the given search
// parameters.
//
// ErrNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetActions(params ActionGetParams) ([]Action, error) {
	actions := make([]jAction, 0)
	err := c.Get("action.get", params, &actions)
	if err != nil {
		return nil, err
	}

	if len(actions) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Actions to Go Actions
	out := make([]Action, len(actions))
	for i, jaction := range actions {
		action, err := jaction.Action()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Action %d in response: %v", i, err)
		}

		out[i] = *action
	}

	return out, nil
}
