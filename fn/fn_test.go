package fn
import "math"
import "testing"


func equal(a, b, eps float64) bool {
    return math.Abs(a-b) < eps
}

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

func TestHypot(t *testing.T) {
	x := 1.0
	y := 2.0
	want := math.Sqrt(x*x + y*y)
	got  := Hypot(x, y)
	if !equal(want, got, 0.0001) {
		t.Errorf("Hypot() = %.2f, want %.2f", got, want)
	}
}

func TestAdd1(t *testing.T) {
	want := 2
	got  :=Add1(1, 1)
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}	
}

func TestSub1(t *testing.T) {
	want := 0
	got  :=Sub1(1, 1)
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}	
}

func TestFirst1(t *testing.T) {
	want := 1
	got  :=First1(1, 2)
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}	
}

func TestZero1(t *testing.T) {
	want := 0
	got  :=Zero1(1, 2)
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}	
}