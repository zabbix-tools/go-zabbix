package zabbix

import (
	"testing"
)

func TestUserMacros(t *testing.T) {
	session := GetTestSession(t)

	params := UserMacroGetParams{}

	macros, err := session.GetUserMacro(params)

	if err != nil {
		t.Fatalf("Error getting user macros: %v", err)
	}

	if len(macros) == 0 {
		t.Fatal("No usermacro found")
	}

	for i, macro := range macros {
		if macro.HostID == "" {
			t.Fatalf("User macro %d returned in response body has no Host ID", i)
		}
	}

	t.Logf("Validated %d user macros", len(macros))
}
