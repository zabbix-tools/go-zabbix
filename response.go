package zabbix

import (
	"encoding/json"
	"fmt"
)

// Response represents the response from a JSON-RPC API request.
//
// This struct maps to the JSON response body described in the Zabbix API
// documentation:
// https://www.zabbix.com/documentation/2.2/manual/api#authentication.
type Response struct {
	// HTTP status code of the API response e.g. 200
	StatusCode int `json:"-"`

	// JSONRPCVersion is the version string of the Zabbix API. This should
	// always be set to "2.0".
	JSONRPCVersion string `json:"jsonrpc"`

	// Body represents the response body as an array of bytes.
	//
	// The Body may be decoded later into a struct with Bind or json.Unmarshal.
	Body json.RawMessage `json:"result"`

	// RequestID is an abitrary identifier which matches the RequestID set in
	// the corresponding API Request.
	RequestID int `json:"id"`

	// Error is populated with error information if the JSON-RPC request
	// succeeded but there was an API error.
	//
	// This struct maps to the JSON response body described in the Zabbix API
	// documentation:
	// https://www.zabbix.com/documentation/2.2/manual/api#error_handling.
	Error APIError `json:"error"`
}

// Err returns an error if the Response includes any error information returned
// from the Zabbix API.
func (c *Response) Err() error {
	if c.Error.Code != 0 {
		return fmt.Errorf("HTTP %d %s (%d)\n%s", c.StatusCode, c.Error.Message, c.Error.Code, c.Error.Data)
	}

	return nil
}

// Bind unmarshals the JSON body of the Response into the given interface.
func (c *Response) Bind(v interface{}) error {
	err := json.Unmarshal(c.Body, v)
	if err != nil {
		return fmt.Errorf("Error decoding JSON response body: %v", err)
	}

	return nil
}
