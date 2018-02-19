package zabbix

import (
	"sync/atomic"
)

// A Request represents a JSON-RPC request to be sent by a client.
//
// This struct maps to the JSON request body described in the Zabbix API
// documentation:
// https://www.zabbix.com/documentation/2.2/manual/api#authentication.
type Request struct {
	// JSONRPCVersion is the version string of the Zabbix API. This should
	// always be set to "2.0".
	JSONRPCVersion string `json:"jsonrpc"`

	// Method is the name of the Zabbix API method to be called.
	Method string `json:"method"`

	// Params is the request's body.
	Params interface{} `json:"params"`

	// RequestID is an abitrary identifier for the Request which is returned in
	// the corresponding API Response to assist with multi-threaded
	// applications. This value is automatically incremented for each new
	// Request by NewRequest.
	RequestID uint64 `json:"id"`

	// AuthToken is the Request's authentication token. When used in a Session,
	// this value is overwritten by the Session.
	AuthToken string `json:"auth,omitempty"`
}

// requestId is a global counter used to assign each APi request a unique ID.
var requestID uint64

// NewRequest returns a new Request given an API method name, and optional
// request body parameters.
func NewRequest(method string, params interface{}) *Request {
	if params == nil {
		params = map[string]string{}
	}
	return &Request{
		JSONRPCVersion: "2.0",
		Method:         method,
		Params:         params,
		RequestID:      atomic.AddUint64(&requestID, 1),
		AuthToken:      "", // set by session
	}
}
