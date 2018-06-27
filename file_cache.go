package zabbix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

/*
cachedSessionData represents a model of cached session.

Example:
{
	"createdAt": 1530056885,
	"session": {
		"url": "...",
		"token": "...",
		"apiVersion": "..."
	}
}
*/
type cachedSessionContainer struct {
	CreatedAt int64 `json:"createdAt"`
	Session   `json:"session"`
}

// SessionFileCache is Zabbix session filesystem cache.
type SessionFileCache struct {
	filePath        string
	sessionLifeTime time.Duration
	filePermissions uint32
}

// SetFilePath sets Zabbix session cache file path. Default value is "./zabbix_session"
func (c *SessionFileCache) SetFilePath(filePath string) *SessionFileCache {
	c.filePath = filePath
	return c
}

// SetFilePermissions sets permissions for a session file. Default value is 0655.
func (c *SessionFileCache) SetFilePermissions(permissions uint32) *SessionFileCache {
	c.filePermissions = permissions
	return c
}

// SetSessionLifetime sets lifetime in seconds of cached Zabbix session. Default value is 4 hours.
func (c *SessionFileCache) SetSessionLifetime(d time.Duration) {
	c.sessionLifeTime = d
}

// SaveSession saves session to a cache
func (c *SessionFileCache) SaveSession(session *Session) error {
	sessionContainer := cachedSessionContainer{
		CreatedAt: time.Now().Unix(),
		Session:   *session,
	}

	serialized, err := json.Marshal(sessionContainer)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.filePath, []byte(serialized), os.FileMode(c.filePermissions))
}

// GetSession returns cached Zabbix session
func (c *SessionFileCache) GetSession() (*Session, error) {
	contents, err := ioutil.ReadFile(c.filePath)

	if err != nil {
		return nil, err
	}

	var sessionContainer cachedSessionContainer

	if err := json.Unmarshal(contents, &sessionContainer); err != nil {
		return nil, err
	}

	// Check if session is expired
	if !c.checkSessionLifeTime(&sessionContainer) {
		// Delete the session file and throw an error if TTL is expired
		os.Remove(c.filePath)
		return nil, fmt.Errorf("cached session lifetime expired")
	}

	return &sessionContainer.Session, err
}

// checkSessionLifeTime checks if session is still actual
func (c *SessionFileCache) checkSessionLifeTime(sessionContainer *cachedSessionContainer) bool {
	now := time.Now().Unix()
	createdAt := sessionContainer.CreatedAt
	timeDiff := now - createdAt

	// Check session TTL by time diff
	isExpired := timeDiff > int64(c.sessionLifeTime)

	return !isExpired
}

// HasSession checks if any valid Zabbix session has been cached and available
func (c *SessionFileCache) HasSession() bool {
	_, err := os.Stat(c.filePath)

	return err == nil
}

// Flush removes a cached session
func (c *SessionFileCache) Flush() error {
	return os.Remove(c.filePath)
}

// NewSessionFileCache creates a new instance of session file system cache
func NewSessionFileCache() *SessionFileCache {
	return &SessionFileCache{
		filePath:        "./zabbix_session",
		sessionLifeTime: 14400, // Default TTL is 4 hours
		filePermissions: 0600,
	}
}
