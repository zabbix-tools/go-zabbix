package zabbix

import (
	"fmt"
)

type ZBXBoolean bool

func (bit *ZBXBoolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return fmt.Errorf("boolean unmarshal error: invalid input %s", asString)
	}
	return nil
}
