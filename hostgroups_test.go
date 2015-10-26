package zabbix

import (
	"testing"
)

func TestHostgroups(t *testing.T) {
	session := GetTestSession(t)

	params := HostgroupGetParams{}

	hostgroups, err := session.GetHostgroups(params)
	if err != nil {
		t.Fatalf("Error getting Hostgroups: %v", err)
	}

	if len(hostgroups) == 0 {
		t.Fatal("No Hostgroups found")
	}

	for i, hostgroup := range hostgroups {
		if hostgroup.GroupID == "" {
			t.Fatalf("Hostgroup %d returned in response body has no Group ID", i)
		}
	}

	t.Logf("Validated %d Hostgroups", len(hostgroups))
}
