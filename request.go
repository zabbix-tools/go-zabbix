package zabbix

import "sync/atomic"

type Request struct {
	JSONRPCVersion string      `json:"jsonrpc"`
	Method         string      `json:"method"`
	Params         interface{} `json:"params"`
	RequestID      uint64      `json:"id"`
	AuthToken      string      `json:"auth,omitempty"`
}

// requestId is a global counter used to assign each APi request a unique ID.
var requestId uint64 = 0

func NewRequest(method string, params interface{}) *Request {
	return &Request{
		JSONRPCVersion: "2.0",
		Method:         method,
		Params:         params,
		RequestID:      atomic.AddUint64(&requestId, 1),
		AuthToken:      "", // set by session
	}
}
