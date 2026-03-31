package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	var b [5]int = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// Sparse initialization.
	// Can provide values to corresponding indices in any order.
	// If no index is provided, the index is the successor of the previous index used,
	// or 0 if no index was used.
	var c [10]int = [10]int{0, 2: 2, 3, 5: 5, 6, 1: 1, 9: 9, 7: 7, 8}
	fmt.Println(c)

	// Implicit length
	d := [...]int{1, 2, 3}
	fmt.Println(d)
	fmt.Println(len(d))

	// Multidimensional array
	var e [2][2]int = [...][2]int{{1, 2}, {3, 4}}
	fmt.Println(e)

	// Accessing array elements
	fmt.Println(a[0])
	fmt.Println(a[1])

	// Change element
	a[1] = 1
	fmt.Println(a[1])

	// Can try to access an element outside of bounds using a for loop.
	// This makes Go panic.
	fmt.Println("Will try to access index out of bounds, making Go panic.")
	const n int = len(d) + 1
	for i := 0; i < n; i++ {
		fmt.Println(d[i])
	}
}
