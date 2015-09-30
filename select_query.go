package zabbix

// SelectQuery represents the query data type for a Zabbix API call.
// Wherever a SelectQuery is required, one of SelectFields, SelectExtendedOutput
// or SelectCount should be given.
//
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference_commentary#data_types
type SelectQuery interface{}

// SelectFields may be given as a SelectQuery in search parameters where each
// member string is the name of a JSON field which should be returned for each
// search result.
//
// For example, for a Host search query:
//
//     query := SelectFields{ "hostid", "host", "name" }
//
type SelectFields []string

const (
	// SelectExtendedOutput may be given as a SelectQuery in search parameters
	// to return all available feilds for all objects in the search results.
	SelectExtendedOutput = "extend"

	// SelectCount may be given as a SelectQuery for supported search parameters
	// to return only the number of available search results, instead of the
	// search result details.
	SelectCount = "count"
)
