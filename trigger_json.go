package zabbix

import "fmt"

// jTrigger is a private map for the Zabbix API Trigger object.
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/trigger/object
type jTrigger struct {
	TriggerID   string       `json:"triggerid"`
	AlarmState  int          `json:"value,string"`
	Description string       `json:"description"`
	Enabled     string       `json:"status"`
	Expression  string       `json:"expression"`
	Groups      jHostgroups  `json:"groups"`
	Hosts       jHosts       `json:"hosts"`
	LastChange  int          `json:"lastchange,string"`
	Severity    int          `json:"priority,string"`
	State       int          `json:"state,string"`
	Tags        jTriggerTags `json:"tags"`
	LastEvent   *jEvent      `json:"lastEvent"`
	URL         string       `json:"url"`
}

type jTriggerTag struct {
	Name  string `json:"tag"`
	Value string `json:"value"`
}

func (c *jTriggerTag) Tag() (*TriggerTag, error) {
	tag := &TriggerTag{}
	tag.Name = c.Name
	tag.Value = c.Value
	return tag, nil
}

type jTriggerTags []jTriggerTag

func (c jTriggerTags) Tags() ([]TriggerTag, error) {
	if c != nil {
		tags := make([]TriggerTag, len(c))
		for i, jTriggerTag := range c {
			tag, err := jTriggerTag.Tag()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Trigger Tag %d in JSON data: %v", i, err)
			}

			tags[i] = *tag
		}

		return tags, nil
	}

	return nil, nil
}

// Trigger returns a native Go Trigger struct mapped from the given JSON Trigger data.
func (c *jTrigger) Trigger() (*Trigger, error) {
	var err error
	trigger := &Trigger{}
	trigger.TriggerID = c.TriggerID
	trigger.AlarmState = c.AlarmState
	trigger.Enabled = (c.Enabled == "1")
	trigger.Description = c.Description
	trigger.Expression = c.Expression
	trigger.URL = c.URL

	if c.LastEvent != nil {
		trigger.LastEvent, err = c.LastEvent.Event()
		if err != nil {
			return nil, err
		}
	}

	// map groups
	trigger.Groups, err = c.Groups.Hostgroups()
	if err != nil {
		return nil, err
	}

	// map hosts
	trigger.Hosts, err = c.Hosts.Hosts()
	if err != nil {
		return nil, err
	}

	trigger.LastChange = c.LastChange
	trigger.Severity = c.Severity
	trigger.State = c.State

	// map tags
	trigger.Tags, err = c.Tags.Tags()
	if err != nil {
		return nil, err
	}

	return trigger, nil
}
