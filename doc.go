/*
Package zabbix provides bindings to interoperate between programs written in Go
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
		"github.com/cavaliercoder/go-zabbix"
	)

	func main() {
		// Default approach - without session caching
		session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Connected to Zabbix API v%s", session.GetVersion())

		// Use session builder with caching.
		// You can use own cache by implementing SessionAbstractCache interface
		// Optionally an http.Client can be passed to the builder, allowing to skip TLS verification,
		// pass proxy settings, etc.

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true
				}
			}
		}
		cache := zabbix.NewSessionFileCache().SetFilePath("./zabbix_session")
		session, err := zabbix.CreateClient("http://zabbix/api_jsonrpc.php").
			WithCache(cache).
			WithHTTPClient(client).
			WithCredentials("Admin", "zabbix").
			Connect()


		fmt.Printf("Connected to Zabbix API v%s", session.GetVersion())
	}

For more information see: https://github.com/cavaliercoder/go-zabbix

*/
package zabbix
