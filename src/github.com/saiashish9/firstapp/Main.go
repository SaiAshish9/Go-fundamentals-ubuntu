package main

// import "fmt"

import (
	"fmt"
	"strconv"
)

//  i:=42 X
var i int = 42

// local package access

// I
// global package access

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

	var a float32 = 42.5
	var j int
	j = int(a)
	// we have to do explicit conversion ourself
	println(j)

	var p int = 32
	fmt.Println(p)

	var q string
	q = string(p)
	// it will return a byte

	fmt.Printf("%v, %T", q, q)

	q = strconv.Itoa(p)

	fmt.Printf("%v, %T", q, q)

	var (
		x string = "A"
		y string = "B"
	)

	// variable blocks

	var (
		i = 32
	)

	fmt.Println(x, y)
	fmt.Printf("%v ,%T", i, i)
}

// we can declare variables twice in same scope
