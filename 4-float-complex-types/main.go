package main

import (
	"fmt"
	"math"
)

func main() {
	var pi32 float32 = math.Pi
	fmt.Println(pi32)
	var pi64 float64 = math.Pi
	fmt.Println(pi64)

	var z complex128 = complex(math.Cos(math.Pi/3), math.Sin(math.Pi/3))
	fmt.Printf("z = %v\n", z)
	fmt.Println(real(z), imag(z))

	fmt.Printf("z^2 = %v\n", z*z)

	var w complex64 = complex64(z)
	fmt.Printf("w = %v\n", w)
}
