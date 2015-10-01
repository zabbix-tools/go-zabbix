package zabbix

import (
	"testing"
)

func TestAlerts(t *testing.T) {
	session := GetTestSession(t)

	params := AlertGetParams{
		SelectHosts: SelectExtendedOutput,
	}

	alerts, err := session.GetAlerts(params)
	if err != nil {
		t.Fatalf("Error getting alerts: %v", err)
	}

	if len(alerts) == 0 {
		t.Fatal("No alerts found")
	}

	for i, alert := range alerts {
		if alert.AlertID == "" {
			t.Fatalf("Alert %d has no Alert ID", i)
		}
	}

	t.Logf("Validated %d Alerts", len(alerts))
}
