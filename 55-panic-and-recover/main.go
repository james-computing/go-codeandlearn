package main

import "fmt"

func main() {
	/* panic
	defer func() {
		fmt.Println("main deffered function called")
	}()
	func1()
	*/

	// Recover
	fmt.Println("Start of program")
	panicExample()
	fmt.Println("End of program")
}

// Panic

func func1() {
	defer func() {
		fmt.Println("func1 deffered function called")
	}()
	func2()
}

func func2() {
	defer func() {
		fmt.Println("func2 deffered function called")
	}()
	panic("func2 panicked")
}

// Recover

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something went wrong")
}
