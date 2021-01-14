package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var a = sync.RWMutex{}

func main() {

	// runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	a.RLock()
	fmt.Println(counter)
	a.RUnlock()
	wg.Done()
}

func increment() {
	a.Lock()
	counter++
	a.Unlock()
	wg.Done()
}
