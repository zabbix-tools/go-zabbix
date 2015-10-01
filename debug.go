package zabbix

import (
	"fmt"
	"os"
)

// debug caches the value of environment variable ZBX_DEBUG from program start.
var debug bool = (os.Getenv("ZBX_DEBUG") == "1")

// dprintf prints formatted debug message to STDERR if the ZBX_DEBUG environment
// variable is set to "1".
func dprintf(format string, a ...interface{}) {
	if debug {
		fmt.Fprintf(os.Stderr, format, a...)
	}
}
