package zabbix

import (
	"testing"
)

func TestHosts(t *testing.T) {
	session := GetTestSession(t)

	params := HostGetParams{
		IncludeTemplates:      true,
		SelectGroups:          SelectExtendedOutput,
		SelectApplications:    SelectExtendedOutput,
		SelectDiscoveries:     SelectExtendedOutput,
		SelectDiscoveryRule:   SelectExtendedOutput,
		SelectGraphs:          SelectExtendedOutput,
		SelectHostDiscovery:   SelectExtendedOutput,
		SelectWebScenarios:    SelectExtendedOutput,
		SelectInterfaces:      SelectExtendedOutput,
		SelectInventory:       SelectExtendedOutput,
		SelectItems:           SelectExtendedOutput,
		SelectMacros:          SelectExtendedOutput,
		SelectParentTemplates: SelectExtendedOutput,
		SelectScreens:         SelectExtendedOutput,
		SelectTriggers:        SelectExtendedOutput,
	}

	hosts, err := session.GetHosts(params)
	if err != nil {
		t.Fatalf("Error getting Hosts: %v", err)
	}

	if len(hosts) == 0 {
		t.Fatal("No Hosts found")
	}

	for i, host := range hosts {
		if host.HostID == "" {
			t.Fatalf("Host %d returned in response body has no Host ID", i)
		}
	}

	t.Logf("Validated %d Hosts", len(hosts))
}
