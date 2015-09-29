package zabbix

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Session struct {
	url        string
	authToken  string
	apiVersion string
}

func NewSession(url string, username string, password string) (session *Session, err error) {
	// create session
	session = &Session{url: url}

	// get Zabbix API version
	res, err := session.Do(NewRequest("apiinfo.version", nil))
	if err != nil {
		return nil, newError("Error getting Zabbix API version: %v", err)
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
		return nil, newError("Error logging in to Zabbix API: %v", err)
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
