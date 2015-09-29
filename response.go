package zabbix

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Response struct {
	StatusCode     int             `json:"-"`
	JSONRPCVersion string          `json:"jsonrpc"`
	Body           json.RawMessage `json:"result"`
	RequestID      int             `json:"id"`
	Error          struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (c *Response) Err() error {
	if c.Error.Code != 0 {
		return errors.New(fmt.Sprintf("HTTP %d %s (%d)\n%s", c.StatusCode, c.Error.Message, c.Error.Code, c.Error.Data))
	}

	return nil
}

func (c *Response) Bind(v interface{}) error {
	err := json.Unmarshal(c.Body, v)
	if err != nil {
		return newError("Error deocding JSON response body: %v", err)
	}

	return nil
}
