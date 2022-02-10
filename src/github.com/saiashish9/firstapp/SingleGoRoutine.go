package main

import (
	"fmt"
	"sync"
)

// creating goroutines
// synchronization waitgroups mutexes
// parallelism
var wg = sync.WaitGroup{}
func main() {
	// go sayHello()
	// to spin up a new thread and run a new thread in that function
	// time.Sleep(100 * time.Millisecond)
	var msg = "hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Good Bye"
	wg.Wait()
	// closure's anonymous fns do have access to outer scope
	// time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("hello")
}
