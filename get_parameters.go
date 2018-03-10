package zabbix

const (
	// SortOrderAscending is a valid value for GetParameters.SortOrder and
	// causes an API query to return all results sorted in ascending order by
	// the fields specified in GetParmeters.SortField.
	SortOrderAscending = "ASC"

	// SortOrderDescending is a valid value for GetParameters.SortOrder and
	// causes an API query to return all results sorted in descending order by
	// the fields specified in GetParmeters.SortField.
	SortOrderDescending = "DESC"
)

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

	// Filter causes an API query to return only results that exactly match the
	// given filter where map keys are the API fields to query and map values
	// are an exact value (or array of values) that must match each key field.
	//
	// Not valid for text fields.
	Filter map[string]interface{} `json:"filter,omitempty"`

	// ResultLimit limits the number of records returned by an API query.
	ResultLimit int `json:"limit,omitempty"`

	// NodeIDs causes an API query to return only the result that belong to the
	// given Zabbix nodes.
	NodeIDs []string `json:"nodeids,omitempty"`

	// OutputFields causes an API query to return only the given fields for each
	// result.
	//
	// Default: SelectExtendedOutput
	OutputFields SelectQuery `json:"output,omitempty"`

	// PreserveKeys causes an API query to return all results using the IDs of
	// each result as a key in the JSON response.
	// This parameter is deliberately not implemented as the the current
	// implementation of JSON decoding expects an array of objects, not an
	// associative array.
	// PreserveKeys bool `json:"preservekeys,omitempty"`

	// TextSearch causes an API query to return only results that match the
	// given wilcard search where the map keys are the desired field names and
	// the map values are the search expression.
	//
	// Only string and text fields are supported.
	TextSearch map[string]string `json:"search,omitempty"`

	// TextSearchByStart causes an API query to return only results that match
	// the search parameters given in TextSearch where each given field starts
	// with the given search expressions.
	//
	// By default, TextSearch will return results that match a search expression
	// anywhere in a field's value; not just the start.
	TextSearchByStart bool `json:"startSearch,omitempty"`

	// SearchByAny currently has an unknown affect (TODO). According to the
	// Zabbix API documentation: If set to true return results that match any of
	// the criteria given in the filter or search parameter instead of all of
	// them.
	SearchByAny bool `json:"searchByAny,omitempty"`

	// EnableTextSearchWildcards enables the use of "*" as a wildcard character
	// in the given TextSearch search criteria.
	EnableTextSearchWildcards bool `json:"searchWildcardsEnabled,omitempty"`

	// Return only hosts that have inventory data matching the given wildcard search.
	// This parameter is affected by the same additional parameters as search.
	SearchInventory map[string]string `json:"searchInventory,omitempty"`

	// SortField causes an API query to return all results sorted by the given
	// field names.
	SortField []string `json:"sortfield,omitempty"`

	// SortOrder causes an API query to return all results sorted in the given
	// order if SortField is defined.
	//
	// Must be one of SortOrderAscending or SortOrderDescending.
	SortOrder string `json:"sortorder,omitempty"`
}
