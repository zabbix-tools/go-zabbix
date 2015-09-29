package zabbix

type HostMacro struct {
	HostMacroID string `json:"hostmacroid"`
	HostID      string `json:"hostid"`
	Macro       string `json:"macro"`
	Value       string `json:"value"`
}
