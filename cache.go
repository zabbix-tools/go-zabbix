package zabbix

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// SessionAbstractCache represents abstract Zabbix session cache backend
type SessionAbstractCache interface {
	// SetSessionLifetime sets lifetime of cached Zabbix session
	SetSessionLifetime(d time.Duration)

	// SaveSession saves session to a cache
	SaveSession(session *Session) error

	// HasSession checks if any valid Zabbix session has been cached and available
	HasSession() bool

	// GetSession returns cached Zabbix session
	GetSession() (*Session, error)
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
	serialized, err := json.Marshal(session)

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

	var session Session

	if err := json.Unmarshal(contents, &session); err != nil {
		return nil, err
	}

	return &session, err
}

// HasSession checks if any valid Zabbix session has been cached and available
func (c *SessionFileCache) HasSession() bool {
	stats, err := os.Stat(c.filePath)

	// Try to check if session file exists and modify date
	if os.IsNotExist(err) {
		return false
	}

	// If file exists, check mod time as used in original portmon-sync
	now := time.Now().Unix()
	lastMod := stats.ModTime().Unix()
	timeDiff := now - lastMod

	// Check session TTL by time diff
	if timeDiff > int64(c.sessionLifeTime) {

		// Delete outdated session file
		os.Remove(c.filePath)

		return false
	}

	return true
}

// NewSessionFileCache creates a new instance of session file system cache
func NewSessionFileCache() *SessionFileCache {
	return &SessionFileCache{
		filePath:        "./zabbix_session",
		sessionLifeTime: 14400, // Default TTL is 4 hours
		filePermissions: 0655,
	}
}
