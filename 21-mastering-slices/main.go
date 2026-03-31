package main

import "fmt"

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	b := a[2:5]
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	c := a[:5]
	fmt.Println(c)

	d := a[2:]
	fmt.Println(d)

	e := a[:]
	fmt.Println(e)
	fmt.Println()

	// Changing a value from a slice changes the original value in the array
	e[0] = 0
	fmt.Println(a)
	fmt.Println(e)
	fmt.Println()

	// Appending creates a new slice, which doesn't change the original array
	e = append(e, 7)
	fmt.Println(a)
	fmt.Println(e)
	fmt.Println()

	f := []int{1, 2, 3}
	g := f[:]
	// Same underlying array, so changing a value from g changes also in f
	g[0] = 0
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println()

	// Appending changes the underlying array
	g = append(g, 4)
	// Now changes in g don't affect f
	g[0] = 10
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println()

	// Make slice with capacity.
	// Capacity must be at most same as a.
	a[0] = 1
	h := a[2:3:5]
	fmt.Println(h)
	fmt.Println(len(h))
	fmt.Println(cap(h))
	fmt.Println()

	// Copy slice to slice.
	i := make([]int, len(f)-1)
	numValuesCopyed := copy(i, f)
	fmt.Println(i)
	fmt.Printf("Number of values copyed = %v\n", numValuesCopyed)
}
