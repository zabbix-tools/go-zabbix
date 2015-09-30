package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Session is an authenticated Zabbix JSON-RPC API client. It must be
// initialized and connected with NewSession.
type Session struct {
	// url is the URL of the Zabbix JSON-RPC API (ending in `/api_jsonrpc.php`).
	url string

	// authToken is the cached authentication token returned by `user.login` and
	// used to authenticate all API calls in this Session.
	authToken string

	// apiVersion is the software version string of the connected Zabbix API.
	apiVersion string
}

// NewSession returns a new Session given an API connection URL and an API
// username and password.
//
// An error is returned if there was an HTTP protocol error, the API credentials
// are incorrect or if the API version is indeterminable.
//
// The authentication token returned by the Zabbix API server is cached to
// authenticate all subsequent requests in this Session.
func NewSession(url string, username string, password string) (session *Session, err error) {
	// create session
	session = &Session{url: url}

	// get Zabbix API version
	res, err := session.Do(NewRequest("apiinfo.version", nil))
	if err != nil {
		return nil, fmt.Errorf("Error getting Zabbix API version: %v", err)
	}

	err = res.Bind(&session.apiVersion)
	if err != nil {
		return
	}

	// login to API
	params := map[string]string{
		"user":     username,
		"password": password,
	}

	res, err = session.Do(NewRequest("user.login", params))
	if err != nil {
		return nil, fmt.Errorf("Error logging in to Zabbix API: %v", err)
	}

	err = res.Bind(&session.authToken)
	if err != nil {
		return
	}

	return
}

// Version returns the software version string of the connected Zabbix API.
func (c *Session) Version() string {
	return c.apiVersion
}

// AuthToken returns the authentication token used by this session to
// authentication all API calls.
func (c *Session) AuthToken() string {
	return c.authToken
}

// Do sends a JSON-RPC request and returns an API Response, using connection
// configuration defined in the parent Session.
//
// An error is returned if there was an HTTP protocol error, a non-200 response
// is received, or if an error code is set is the JSON response body.
//
// When err is nil, resp always contains a non-nil resp.Body.
//
// Generally Get, Post, or PostForm will be used instead of Do.
func (c *Session) Do(req *Request) (resp *Response, err error) {
	// configure request
	req.AuthToken = c.authToken

	// encode request as json
	b, err := json.Marshal(req)
	if err != nil {
		return
	}

	// create HTTP request
	r, err := http.NewRequest("POST", c.url, bytes.NewReader(b))
	if err != nil {
		return
	}
	r.ContentLength = int64(len(b))
	r.Header.Add("Content-Type", "application/json-rpc")

	// send request
	client := http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return
	}

	defer res.Body.Close()

	// read response body
	b, err = ioutil.ReadAll(res.Body)

	// map response
	resp = &Response{
		StatusCode: res.StatusCode,
	}

	// unmarshal response body
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}

	// check for API errors
	if err = resp.Err(); err != nil {
		return
	}

	return
}
