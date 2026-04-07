package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	counter := &Counter{count: 0}

	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for j := 0; j < 100; j++ {
				counter.Inc()
			}
		}()
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Count:", counter.Get())
			time.Sleep(2 * time.Millisecond)
		}()
	}

	waitGroup.Wait()
	fmt.Println("Final count:", counter.count)
}

type Counter struct {
	mu    sync.RWMutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
