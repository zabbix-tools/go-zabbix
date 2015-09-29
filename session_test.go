package zabbix

import (
	"os"
	"testing"
)

var session *Session = nil

func GetTestSession(t *testing.T) *Session {
	var err error
	if session == nil {
		url := os.Getenv("ZBX_URL")
		if url == "" {
			url = "http://localhost:8080/api_jsonrpc.php"
		}

		username := os.Getenv("ZBX_USERNAME")
		if username == "" {
			username = "Admin"
		}

		password := os.Getenv("ZBX_PASSWORD")
		if password == "" {
			password = "zabbix"
		}

		session, err = NewSession(url, username, password)
		if err != nil {
			t.Fatalf("Error creating a session: %v", err)
		}
	}

	return session
}

func TestSession(t *testing.T) {
	s := GetTestSession(t)

	if s.Version() == "" {
		t.Errorf("No API version found for session")
	}
}
