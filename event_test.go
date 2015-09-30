package zabbix

import (
	"testing"
)

func TestEvents(t *testing.T) {
	session := GetTestSession(t)

	params := EventGetParams{
		SelectAcknowledgements: SelectExtendedOutput,
		SelectAlerts:           SelectExtendedOutput,
		SelectHosts:            SelectExtendedOutput,
		SelectRelatedObject:    SelectExtendedOutput,
	}

	events, err := session.GetEvents(params)
	if err != nil {
		t.Fatalf("Error getting events: %v", err)
	}

	if len(events) == 0 {
		t.Fatal("No events found")
	}

	for i, event := range events {
		if event.EventID == "" {
			t.Fatalf("Event %d has no Event ID", i)
		}

		if event.Timestamp.IsZero() {
			t.Fatalf("Event %d has no timestamp", i)
		}
	}

	t.Logf("Validated %d Events", len(events))
}
