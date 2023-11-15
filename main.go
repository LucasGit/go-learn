package main

import "fmt"
import "example.com/go-learn/greet"
import "example.com/go-learn/fn"


func main() {
    fmt.Println("hello world!");

	greet.Hello("world");

    fmt.Printf("10 vs 3 max is %d\n", fn.Max_val(10, 3));
	fmt.Println(fn.Hello());

}
