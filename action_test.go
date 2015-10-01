package zabbix

import (
	"testing"
)

func TestActions(t *testing.T) {
	session := GetTestSession(t)

	params := ActionGetParams{}

	actions, err := session.GetActions(params)
	if err != nil {
		t.Fatalf("Error getting actions: %v", err)
	}

	if len(actions) == 0 {
		t.Fatal("No actions found")
	}

	for i, action := range actions {
		if action.ActionID == "" {
			t.Fatalf("Action %d has no Action ID", i)
		}

		if action.Name == "" {
			t.Fatalf("Action %d has no name", i)
		}

		if action.EventType == EventSourceTrigger && action.ProblemMessageSubject == "" {
			t.Fatalf("Action %d has no problem message subject", i)
		}
	}

	t.Logf("Validated %d Actions", len(actions))
}
