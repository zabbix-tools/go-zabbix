package zabbix

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type UnixTimestamp struct {
	*time.Time
}

func (t UnixTimestamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Unix())
	return []byte(stamp), nil
}

func (t *UnixTimestamp) UnmarshalJSON(data []byte) (err error) {
	var unixString string
	err = json.Unmarshal(data, &unixString)
	if err != nil {
		return
	}

	unix, err := strconv.ParseInt(unixString, 10, 64)

	// Fractional seconds are handled implicitly by Parse.
	tt := time.Unix(unix, 0)
	*t = UnixTimestamp{&tt}

	return
}
