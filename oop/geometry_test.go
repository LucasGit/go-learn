package geometry
import "math"
import "testing"


func equal(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func TestDistance(t *testing.T) {

	p := Point{1, 2}
	q := Point{4, 6}
	want := 5.0
	got := p.Distance(q)

	if  !equal(want, got, 0.0001) {
		t.Errorf("want %.3f, got %.3f", want, got)
	}

}

