package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go add()

	time.Sleep(5 * time.Second)
}

var cond *sync.Cond = sync.NewCond(&sync.Mutex{})
var integers []int = make([]int, 0, 10)

func remove(delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Length before removing:", len(integers))
	// remove first
	integers = integers[1:]
	fmt.Println("Length after remove:", len(integers))
	//cond.Signal()
	cond.Broadcast()
}

func add() {
	for i := 0; i < 5; i++ {
		cond.L.Lock()
		for len(integers) == 2 {
			cond.Wait()
			fmt.Println("Length before adding:", len(integers))
		}
		integers = append(integers, i)
		fmt.Println("Length after adding:", len(integers))
		go remove(time.Second)
		cond.L.Unlock()
	}
}
