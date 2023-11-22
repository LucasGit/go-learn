package fn_anonymous
import "testing"


func TestSquares(t *testing.T) {
	f := Squares()
	want := 1
	if got := f(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

