package fn_recursive
import "testing"

func TestFibonacci(t *testing.T) {
	want := 2
	if got := Fibonacci(3); got != want {
		t.Errorf("go= %d, want %d", got, want)
	}
}
