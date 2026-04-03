package main

import "fmt"

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println()

	numbers := []int{1, 2, 3, 4}
	result := sliceOperation(numbers, func(x int) int {
		return x * x
	})
	fmt.Println(result)
}

func counter() func() int {
	counter := 0
	// Same function will be called multiple times without resseting counter,
	// this increments the counter.
	// The closure is this function to returned, which closes over the variable counter.
	return func() int {
		counter++
		return counter
	}
}

func sliceOperation(s []int, op func(int) int) []int {
	// Can't use s, otherwise will change the underlying array.
	// Must use a new slice.
	result := make([]int, len(s))
	for i, x := range s {
		result[i] = op(x)
	}
	return result
}
