package zabbix

import (
	"fmt"
)

// APIError represents a Zabbix API error.
type APIError struct {
	// Code is the Zabbix API error code.
	Code int `json:"code"`

	// Message is a short error summary.
	Message string `json:"message"`

	// Data is a detailed error message.
	Data string `json:"data"`
}

func (e *APIError) String() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.Code)
}

// Error returns the string representation of an APIError
func (e *APIError) Error() string {
	return e.String()
}
