package funct

import "testing"

func TestMax(t *testing.T) {
	got := Max(1, 2)
	if got != 2 {
		t.Errorf("max 1,2 failed, got %d", got)
	}
}
