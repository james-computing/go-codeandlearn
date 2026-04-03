package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(sum(nums...))

	f("John", []int{1, 2}...)
}

func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func f(name string, nums ...int) {
	fmt.Println(name)
	fmt.Println(nums)
}
