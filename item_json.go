package zabbix

import (
	"fmt"
	"strconv"
)

// jItem is a private map for the Zabbix API Host object.
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/item/get
type jItem struct {
	HostID    string `json:"hostid,omitempty"`
	ItemID    string `json:"itemid"`
	ItemName  string `json:"name"`
	ItemDescr string `json:"description,omitempty"`
	LastClock string `json:"lastclock,omitempty"`
	LastValue string `json:"lastvalue,omitempty"`
	LastValueType string `json:"value_type"`
}

// Item returns a native Go Item struct mapped from the given JSON Item data.
func (c *jItem) Item() (*Item, error) {
	var err error
	item := &Item{}
	item.HostID, err = strconv.Atoi(c.HostID)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Host ID: %v", err)
	}
	item.ItemID, err = strconv.Atoi(c.ItemID)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Item ID: %v", err)
	}
	item.ItemName = c.ItemName
	item.ItemDescr = c.ItemDescr

	item.LastClock, err = strconv.Atoi(c.LastClock)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Item LastClock: %v", err)
	}
	item.LastValue = c.LastValue

	item.LastValueType, err = strconv.Atoi(c.LastValueType)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Item LastValueType: %v", err)
	}
	return item, err
}

// jItems is a slice of jItems structs.
type jItems []jItem

// Items returns a native Go slice of Items mapped from the given JSON ITEMS
// data.
func (c jItems) Items() ([]Item, error) {
	if c != nil {
		items := make([]Item, len(c))
		for i, jitem := range c {
			item, err := jitem.Item()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Item %d in JSON data: %v", i, err)
			}
			items[i] = *item
		}

		return items, nil
	}

	return nil, nil
}
