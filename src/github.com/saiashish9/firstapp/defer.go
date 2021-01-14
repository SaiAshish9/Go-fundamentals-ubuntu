package main

import "fmt"

func main() {

	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")
	// defer fmt.Println("middle1")
	// defer takes the value at the time defer is called

	a := "start"
	defer fmt.Println(a)
	a = "end"

}

// LIFO
// nil
// log net/http fmt io/ioutil defer res.Bod y.Close()
