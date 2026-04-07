package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := newMyStructPool(5)
	for i := range 10 {
		item := pool.Get().(*MyStruct)
		item.data = fmt.Sprintf("data %d", i)
		fmt.Println(item.data)
		pool.Put(item)
	}
}

type MyStruct struct {
	data string
}

/*
var pool sync.Pool = sync.Pool{
	New: func() any {
		return &MyStruct{}
	},
}
*/

func newMyStructPool(initialCap int) *sync.Pool {
	p := &sync.Pool{
		New: func() any {
			return &MyStruct{}
		},
	}

	for range initialCap {
		p.Put(&MyStruct{})
	}

	return p
}
