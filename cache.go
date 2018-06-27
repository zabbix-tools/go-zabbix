package zabbix

import (
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

	// Flush removes cached session
	Flush() error
}
