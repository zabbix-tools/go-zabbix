package zabbix

import (
	"testing"
)

func TestHostinterfaces(t *testing.T) {
	session := GetTestSession(t)

	params := HostinterfaceGetParams{}

	hostinterfaces, err := session.GetHostinterfaces(params)

	if err != nil {
		t.Fatalf("Error getting hostinterfaces: %v", err)
	}

	if len(hostinterfaces) == 0 {
		t.Fatalf("No hostinterfaces found")
	}

	for i, hostinterface := range hostinterfaces {
		if hostinterface.HostID == "" {
			t.Fatalf("Hostinterface %d returned in response body has no HostID", i)
		}
		if hostinterface.InterfaceID == "" {
			t.Fatalf("Hostinterface %d returned in response body has no InterfaceID", i)
		}
		switch hostinterface.Main {
		case 0, 1:
		default:
			t.Fatalf("Hostinterface %d returned in response body has invalid Hostinterface Main value: %v", i, hostinterface.Main)
		}
		switch hostinterface.Type {
		case 0, 1, 2, 3, 4:
		default:
			t.Fatalf("Hostinterface %d returned in response body has invalid Hostinterface Type value: %v", i, hostinterface.Type)
		}
		switch hostinterface.Useip {
		case 0, 1:
		default:
			t.Fatalf("Hostinterface %d returned in response body has invalid Hostinterface Useip value: %v", i, hostinterface.Useip)
		}
	}
}
