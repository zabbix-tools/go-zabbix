package zabbix

import (
	"fmt"
	"time"
)

// JMaintenance is a private map for the Zabbix API Maintenance object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/maintenance/object
type JMaintenance struct {
	MaintenanceID   string `json:"maintenanceid"`
	Name            string `json:"name"`
	ActiveSince     int64  `json:"active_since,string"`
	ActiveTill      int64  `json:"active_till,string"`
	Description     string `json:"description"`
	MaintenanceType int    `json:"maintenance_type,string"`
	TagsEvaltype    int    `json:"tags_evaltype,string"`
}

// Timeperiods is a private map for the Zabbix API Maintenance object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/maintenance/object
type Timeperiods struct {
	TimeperiodType int `json:"timeperiod_type,int"`
	Every          int `json:"every,string"`
	Dayofweek      int `json:"dayofweek,string"`
	StartTime      int `json:"start_time,string"`
	Period         int `json:"period,string"`
}

// Maintenance returns a native Go Maintenance struct mapped from the given JSON Maintenance
// data.
func (c *JMaintenance) Maintenance() (result *Maintenance, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	maintenance := &Maintenance{
		ActionEvalTypeAndOr: TagsEvaltype(c.MaintenanceType),
		Type:                MaintenanceType(c.TagsEvaltype),
		ActiveSince:         time.Unix(c.ActiveSince, 0),
		ActiveTill:          time.Unix(c.ActiveTill, 0),
		Description:         c.Description,
		MaintenanceID:       c.MaintenanceID,
		Name:                c.Name,
	}

	return maintenance, nil
}
