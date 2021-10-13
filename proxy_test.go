package zabbix

import (
	"testing"
)

func TestProxy(t *testing.T) {
	session := GetTestSession(t)
	create_ap := ProxyCreateParams{
		ProxyName:   "Test Active Proxy",
		ProxyStatus: 5,
	}
	create_pp := ProxyCreateParams{
		ProxyName:   "Test Passive Proxy",
		ProxyStatus: 6,
		Interface: ProxyInterface{
			UseIP: 1,
			IP:    "127.0.0.1",
			DNS:   "",
			Port:  "10052",
		},
	}
	resp_create_ap, err := session.CreateProxy(create_ap)
	if err != nil {
		t.Fatalf("Error create active proxy: %v", err)
	}
	t.Logf("Created active proxy with id %q", resp_create_ap.ProxyIDs[0])
	resp_create_pp, err := session.CreateProxy(create_pp)
	if err != nil {
		t.Fatalf("Error create passive proxy: %v", err)
	}
	t.Logf("Created passive proxy with id %q", resp_create_pp.ProxyIDs[0])
	delete_ap := ProxyDeleteParams{
		ProxyID: resp_create_ap.ProxyIDs,
	}
	delete_pp := ProxyDeleteParams{
		ProxyID: resp_create_pp.ProxyIDs,
	}
	resp_delete_ap, err := session.DeleteProxy(delete_ap)
	if err != nil {
		t.Fatalf("Error delete active proxy: %v", err)
	}
	t.Logf("Deleted active proxy with id %q", resp_delete_ap.ProxyIDs[0])
	resp_delete_pp, err := session.DeleteProxy(delete_pp)
	if err != nil {
		t.Fatalf("Error delete passive proxy: %v", err)
	}
	t.Logf("Deleted passive proxy with id %q", resp_delete_pp.ProxyIDs[0])

}
