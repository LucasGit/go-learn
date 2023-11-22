package fn
import "math"
// import "strings"

/*
func name(parameter-list) (result-list) {
    body
}
*/

func Max_val(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Hello() string {
	return "hello world"
}


// declaration
func Hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}


func f_sum_args1(i, j, k int, s, t string) {

}

func f_sum_args2(i int, j int, k int,  s string, t string) {

}

func Add1(x int, y int) int   {return x + y}
func Sub1(x, y int) (z int)   { z = x - y; return}
func First1(x int, _ int) int { return x }
func Zero1(int, int) int      { return 0 }

/*
fmt.Printf("%T\n", add)   // "func(int, int) int"
fmt.Printf("%T\n", sub)   // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero)  // "func(int, int) int"

*/

// func add1(r rune) rune { return r + 1 }

// fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
// fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
// fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"

// anonymous function
// strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")


func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

// fmt.Println(sum())           // "0"
// fmt.Println(sum(3))          // "3"
// fmt.Println(sum(1, 2, 3, 4)) // "10"



