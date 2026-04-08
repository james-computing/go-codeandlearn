package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//leakExample()
	doneChannelExample()
}

func someWork(waitGroup *sync.WaitGroup) chan int {
	ch := make(chan int)

	go func() {
		defer waitGroup.Done()
		ch <- 1
		fmt.Println("Goroutine finished")
	}()

	return ch
}

func leakExample() {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	someWork(waitGroup)
	// At this point, the goroutine from someWork is blocked and won't finish execution
}

func anotherWork(done <-chan struct{}, waitGroup *sync.WaitGroup) chan int {
	ch := make(chan int)

	go func() {
		defer waitGroup.Done()
		select {
		case ch <- 1:
			fmt.Println("Unblocked")
		case <-done:
			fmt.Println("Done")
		}

		fmt.Println("Goroutine finished")
	}()

	return ch
}

func doneChannelExample() {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	done := make(chan struct{})
	anotherWork(done, waitGroup)
	fmt.Println("Wait 1 second")
	time.Sleep(time.Second)
	close(done)
	waitGroup.Wait()
}
