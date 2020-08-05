# go-zabbix

Go bindings for the Zabbix API

[![go report card](https://goreportcard.com/badge/github.com/cavaliercoder/go-zabbix "go report card")](https://goreportcard.com/report/github.com/cavaliercoder/go-zabbix)
[![GPL license](https://img.shields.io/badge/license-GPL-brightgreen.svg)](https://opensource.org/licenses/gpl-license)
[![GoDoc](https://godoc.org/github.com/cavaliercoder/go-zabbix?status.svg)](https://godoc.org/github.com/cavaliercoder/go-zabbix)

## Overview

This project provides bindings to interoperate between programs written in Go
language and the Zabbix monitoring API.

A number of Zabbix API bindings already exist for Go with varying levels of
maturity. This project aims to provide an alternative implementation which is
stable, fast, and allows for loose typing (using types such as`interface{}` or
`map[string]interface{}`) as well as strong types (such as `Host` or `Event`).

The package aims to have comprehensive coverage of Zabbix API methods from v1.8
through to v3.0 without introducing limitations to the native API methods.

## Getting started

```go
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/cavaliercoder/go-zabbix"
)

func main() {
	// Default approach - without session caching
	session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix")
	if err != nil {
		panic(err)
	}

	version, err := session.GetVersion()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to Zabbix API v%s", version)

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
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	version, err := session.GetVersion()

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("Connected to Zabbix API v%s", version)
}
```

## License

Released under the [GNU GPL License](https://github.com/cavaliercoder/go-zabbix/blob/master/LICENSE)
