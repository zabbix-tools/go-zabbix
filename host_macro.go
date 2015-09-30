package zabbix

// HostMacro represents a Zabbix Host Macro returned from the Zabbix API.
type HostMacro struct {
	// HostMacroID is the unique ID of the Host Macro.
	HostMacroID string `json:"hostmacroid"`

	// HostID is the ID of the Host which owns this Macro.
	HostID string `json:"hostid"`

	// Macro is the name of the Macro (e.g. '{HOST.MACRO}').
	Macro string `json:"macro"`

	// Value is the value of the Macro.
	Value string `json:"value"`
}
