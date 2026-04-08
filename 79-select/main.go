package main

import (
	"fmt"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	ch1 := make(chan any)
	ch2 := make(chan any)

	go func() {
		for i := range 3 {
			ch1 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := range 3 {
			ch2 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case value := <-ch1:
			fmt.Println("Received from ch1:", value)
		case value := <-ch2:
			fmt.Println("Received from ch2:", value)
		case <-time.After(time.Second):
			fmt.Println("Returning after timeout")
			return
			/*
				default:
					fmt.Println("default")
					time.Sleep(time.Second)
			*/
		}
	}
}

func example2() {
	ch1 := make(chan any)
	ch2 := make(chan any)

	go func() {
		for i := range 3 {
			ch1 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := range 3 {
			ch2 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for range 6 {
		select {
		case value := <-ch1:
			fmt.Println("Received from ch1:", value)
		case value := <-ch2:
			fmt.Println("Received from ch2:", value)
		}
	}
}
