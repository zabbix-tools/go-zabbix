# go-zabbix [![GoDoc](https://godoc.org/github.com/cavaliercoder/go-zabbix?status.svg)](https://godoc.org/github.com/cavaliercoder/go-zabbix)

*Go bindings for the Zabbix API*

This project provides bindings to interoperate between programs written in Go
language and the Zabbix monitoring API.

A number of Zabbix API bindings already exist for Go with varying levels of
maturity. This project aims to provide an alternative implementation which is
stable, fast, and allows for loose typing (using types such as`interface{}` or
`map[string]interface{}`) as well as strong types (such as `Host` or `Event`).

The package aims to have comprehensive coverage of Zabbix API methods from v1.8
through to v3.0 without introducing limitations to the native API methods.

## Example

```go
package main

import (
	"fmt"
	"github.com/cavaliercoder/zabbix"
)

func main() {
	// Default approach - without session caching
	session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix")
	if err != nil {
		panic(err)
	}

	// Use session builder with caching.
	// You can use own cache by implementing SessionAbstractCache interface

	cache := zabbix.NewSessionFileCache().SetFilePath("./zabbix_session")
	session, err := zabbix.CreateClient("http://zabbix/api_jsonrpc.php").
		WithCache(cache).
		WithCredentials("Admin", "zabbix").
		Connect()

	fmt.Printf("Connected to Zabbix API v%s", session.Version())
}
```