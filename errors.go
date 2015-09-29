package zabbix

import (
	"errors"
	"fmt"
)

func newError(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a...))
}
