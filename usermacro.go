package zabbix

// UserMacroGetParams represent the parameters for a `usermacro.get` API call (see zabbix documentation).
type UserMacroGetParams struct {
	GetParameters

	// Return global macros instead of host macros.
	GlobalMacro bool `json:"globalmacro,omitempty"`

	// Return only global macros with the given IDs.
	GlobalMacroIDs []string `json:"globalmacroids,omitempty"`

	// Return only host macros that belong to hosts or templates from the given host groups.
	GroupIDs []string `json:"groupids,omitempty"`

	// Return only macros that belong to the given hosts or templates.
	HostIDs []string `json:"hostids,omitempty"`

	// Return only host macros with the given IDs.
	HostMacroIDs []string `json:"hostmacroids,omitempty"`

	// Return only host macros that belong to the given templates. (Zabbix 2.x only)
	TemplateIDs []string `json:"templateids,omitempty"`

	// TODO: add selectGroups, selectHosts and selectTemplates queries
}

// GetUserMacro queries the Zabbix API for user macros matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetUserMacro(params UserMacroGetParams) ([]HostMacro, error) {
	macros := make([]HostMacro, 0)

	if err := c.Get("usermacro.get", params, &macros); err != nil {
		return nil, err
	}

	if len(macros) == 0 {
		return nil, ErrNotFound
	}

	return macros, nil
}
