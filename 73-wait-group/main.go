package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	example1()
	example2()
}

func sayHello(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Hello")
}

func example1() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go sayHello(&waitGroup)
	go sayHello(&waitGroup)

	waitGroup.Wait()

	fmt.Println("Both goroutines waited")
	fmt.Println()
}

func printNumber(waitGroup *sync.WaitGroup, number int) {
	defer waitGroup.Done()
	fmt.Println("Printing number:", number)
	time.Sleep(time.Second)
}

func example2() {
	var waitGroup sync.WaitGroup

	for i := 1; i <= 3; i++ {
		waitGroup.Add(1)
		go printNumber(&waitGroup, i)
	}
	waitGroup.Wait()

	fmt.Println("All numbers printed")
}
