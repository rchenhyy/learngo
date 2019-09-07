package main

import (
	"fmt"
	"sync"
	"time"
)

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()
}

func createWorker(i int, wg *sync.WaitGroup) *worker {
	w := worker{
		in: make(chan int),
		//wg: wg,
		done: func() {
			wg.Done()
		},
	}
	go doWork(i, &w)
	return &w
}

func doWork0(i int, c chan int) {
	func() {
		for v := range c {
			fmt.Printf("Worker %d received %c\n", i, v) // read
		}
	}()
	fmt.Printf("Worker done")
}

func doWork(i int, w *worker) {
	func() {
		for v := range w.in {
			fmt.Printf("Worker %d received %c\n", i, v) // read
			//go func() { w.done <- true }()
			w.done()
		}
	}()
}

func channelDemo() {
	var workers [10]*worker
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg) // go first!
	}
	for i, w := range workers {
		w.in <- 'a' + i // send
	}
	for i, w := range workers {
		w.in <- 'A' + i // send
	}

	/*for _, w := range workers {
		<-w.done
		<-w.done
	}*/
	wg.Wait()
	fmt.Println("All done!")
}

func bufferedChannel() {
	c := make(chan int, 3)
	go doWork0(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go doWork0(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// channel
	fmt.Println("Channel as first-class citizen")
	channelDemo()

	// buffered channel
	fmt.Println("Buffered channel")
	bufferedChannel()

	// close & range
	channelClose()
}
