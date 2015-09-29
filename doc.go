/*
This project provides bindings to interoperate between programs written in Go
language and the Zabbix monitoring API.

A number of Zabbix API bindings already exist for Go with varying levels of
maturity. This project aims to provide an alternative implementation which is
stable, fast, and allows for loose typing (using types such as interface{} or
map[string]interface{}) as well as strong types (such as zabbix.Host or
zabbix.Event).

The package aims to have comprehensive coverage of Zabbix API methods from v1.8
through to v3.0 without introducing limitations to the native API methods.

	package main

	import (
		"fmt"
		"github.com/cavaliercoder/zabbix"
	)

	func main() {
		session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Connected to Zabbix API v%s", session.Version())
	}

For more information see: https://github.com/cavaliercoder/go-zabbix

*/
package zabbix
