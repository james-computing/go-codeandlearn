package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	twoHours := 2 * time.Hour
	futureTime := now.Add(twoHours)
	fmt.Println("Time 2 hours from now:", futureTime)
	fmt.Println()

	// Measure execution time.
	start := time.Now()
	// This loop should be optimized to be executed at compile time,
	// but the non optimized code will be used to demonstrate the time functionality.
	a := 1
	for i := 0; i < 10000000; i++ {
		a++
	}
	elapsed := time.Since(start)
	fmt.Println("Loop execution time:", elapsed)
}
