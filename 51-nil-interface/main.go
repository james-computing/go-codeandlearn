package main

import "fmt"

func main() {
	var a *int
	var i interface{}

	// both are nil
	fmt.Println(a)
	fmt.Println(a == nil)
	fmt.Println(i)
	fmt.Println(i == nil)

	i = a
	// Printing still writes <nil>
	fmt.Println(i)
	// But comparing to nil gives false,
	// because now i has type *int.
	fmt.Println(i == nil)
}
