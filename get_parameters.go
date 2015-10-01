package zabbix

// GetParameters represents the common parameters for all API Get methods.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference_commentary#common_get_method_parameters
type GetParameters struct {
	// CountOutput indicates whether an API call should return the number of
	// records in the result instead of the actual data.
	CountOutput bool `json:"countOutput,omitempty"`

	// EditableOnly indicates whether an API call should only return results
	// that the user has write permissions to.
	EditableOnly bool `json:"editable,omitempty"`

	// ExcludeSearch indicates whether an API call should only return result
	// that do not match the given Search parameter.
	ExcludeSearch bool `json:"excludeSearch,omitempty"`

	Filter map[string]interface{} `json:"filter,omitempty"`
}
