package main

import "fmt"

//  i:=42 X
var i int = 42

func main() {
	// var i int
	// i = 42
	// var i int = 42
	// i := 42
	// var j float32 = 27
	// go compiler can
	// automatically figure out
	// the datatype itself
	// fmt.Println(i)
	// fmt.Printf("%v ,%T", j, j)

	fmt.Printf("%v ,%T", i, i)
}
