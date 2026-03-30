package main

import "fmt"

func main() {
	var x float64 = 0.0
	// Doesn't give any errors or warnings, just evaluates to +Inf or -Inf
	fmt.Println(1 / x)
	fmt.Println(-1 / x)

	// But this doesn't compile
	//var y float64 = 1 / 0
}
