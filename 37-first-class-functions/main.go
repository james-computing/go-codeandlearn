package main

import "fmt"

func main() {
	result := apply(double, 5)
	fmt.Println(result)

	result = apply(multiplyBy(5), 5)
	fmt.Println(result)
}

func double(x int) int {
	return 2 * x
}

func apply(f func(int) int, x int) int {
	return f(x)
}

type multiplier func(int) int

func multiplyBy(m int) multiplier {
	return func(x int) int {
		return m * x
	}
}
