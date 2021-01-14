package main

import (
	"fmt"
	"log"
)

func main() {
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("err", err)
		}
	}()
	panic("something bad happened")
	log.Println("after panic")
}

// Defer
// Used to delay execution of a statement until dunction exits
// Useful to group "open" and "close" functions together
// Run in LIFO
// Arguments evaluated at time defer is executed, not at time of called function execution

// Panic
// Occurs when program cannot continue at all
// cannot obtain TCP port for web server
// if nothing handles panic, program will exit

// Recover
// Used to recover from panics
// Only useful in deferred functions
// Current function will not attempt to
// continue , but higher functions in call stack
// will

// import "net/http"
//
// func main(){
// http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request))
// w.write([]byte("Hello Go!"))
// }
// err := http.ListenAndServe(":8000",nill)
// if err != nill {
// panic(err.Error())
// }
//
