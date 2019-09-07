package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() { // lock in the scale of code block
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func main() {
	var a atomicInt

	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.value)
}
