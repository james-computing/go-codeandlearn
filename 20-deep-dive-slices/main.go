package main

import "fmt"

func main() {
	// Nil slice
	var a []int
	fmt.Println(a == nil)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = []int{1, 2, 3}
	fmt.Println(a == nil)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println()

	// Empty slice
	var b []int = []int{}
	fmt.Println(b == nil)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println()

	// Make with length = capacity
	var c []int = make([]int, 5)
	fmt.Println(c == nil)
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	fmt.Println()

	// General make
	var d []int = make([]int, 5, 7)
	fmt.Println(d == nil)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	fmt.Println()

	// Append
	a = append(a, 4)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// Append multiple
	a = append(a, 5, 6, 7)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println()

	// Append slice
	a = []int{1, 2, 3}
	b = []int{4, 5, 6, 7}
	a = append(a, b...)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println()
}
