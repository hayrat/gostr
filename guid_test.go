package gostr

import "testing"

func TestGuid(t *testing.T) {
	if guid := Guid(); len(guid) != 36 {
		t.Errorf("Geçersi guid: %v", guid)
	}
}
