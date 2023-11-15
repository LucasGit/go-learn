package fn

import "testing"

func TestMax_val(t *testing.T) {

	got := Max_val(1, 2)
	if got != 2 {
		t.Errorf("max 1,2 failed, got %d", got)
	}

	
}

func TestHello(t *testing.T) {
	want := "hello world"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
