package main

import (
	"fmt"
	"sync"
)

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	// bidirectional channel
	var intChannel chan int
	intChannel = make(chan int)

	// unidirecitonal channels
	//readChannel := make(<-chan int)
	//writeChannel := make(chan<- int)

	go func() {
		intChannel <- 10
	}()

	readValue, ok := <-intChannel
	fmt.Println(readValue)
	fmt.Println(ok)

	close(intChannel)

	readValue, ok = <-intChannel
	fmt.Println(readValue)
	fmt.Println(ok)
}

func example2() {
	intChannel := make(chan int)

	go func() {
		for i := range 5 {
			intChannel <- i
		}
		close(intChannel)
	}()

	// reads a value until the channel is closed
	for j := range intChannel {
		fmt.Println(j)
	}
	fmt.Println("end")
}

func example3() {
	signal := make(chan any)
	waitGroup := &sync.WaitGroup{}

	for i := range 5 {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()

			// blocks here
			<-signal

			fmt.Println("Started...")
		}(i)
	}

	// unblocks all created goroutines
	close(signal)

	waitGroup.Wait()
}

func example4() {
	// buffered channel, can write to it n times before blocking
	const n int = 2
	stringChannel := make(chan string, n)
	stringChannel <- "hello"
	stringChannel <- "world"
	fmt.Println(<-stringChannel)
	fmt.Println(<-stringChannel)
	close(stringChannel)
}
