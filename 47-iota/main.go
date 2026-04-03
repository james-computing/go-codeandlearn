package main

import (
	"fmt"
)

// Similar to an enum, but different
const (
	Small int = iota
	Medium
	Large
)

const (
	a int = 3 * iota
	b
	c
)

func main() {
	fmt.Println(Small)
	fmt.Println(Medium)
	fmt.Println(Large)
	fmt.Println()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
