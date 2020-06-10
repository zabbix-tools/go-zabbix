package zabbix

import "testing"

func TestProxies(t *testing.T) {
	session := GetTestSession(t)

	params := ProxyGetParams{}

	proxies, err := session.GetProxies(params)

	if err != nil {
		t.Fatalf("Error getting proxies: %v", err)
	}

	if len(proxies) == 0 {
		t.Fatalf("No proxies found")
	}

	for i, proxy := range proxies {
		if proxy.Hostname == "" {
			t.Fatalf("Proxy %d returned in response body has no Hostname", i)
		}
		switch proxy.Status {
		case 5, 6:
		default:
			t.Fatalf("Proxy %d returned in response body has invalid Proxy Status: %d", i, proxy.Status)
		}
	}

	t.Logf("Validated %d proxies", len(proxies))
}
