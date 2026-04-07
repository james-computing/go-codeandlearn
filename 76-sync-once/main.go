package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(3)

	for range 3 {
		go doSomething(waitGroup)
	}
	waitGroup.Wait()
}

var once sync.Once
var initialized bool = false

func initSomething() {
	fmt.Println("Initializing...")
	initialized = true
}

func doSomething(waitGroup *sync.WaitGroup) {
	// Executed once
	once.Do(initSomething)
	// Not executed
	once.Do(func() {
		fmt.Println("Do something else")
	})

	fmt.Println("Doing something...")
	waitGroup.Done()
}
