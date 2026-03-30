package main

import "fmt"

func main() {
	var myString string
	// prints ""
	fmt.Println(myString)
	myString = "hello"
	fmt.Println(myString)

	var myInt int
	// prints 0
	fmt.Println(myInt)
	myInt = 10
	fmt.Println(myInt)

	var myBool bool
	// prints false
	fmt.Println(myBool)
	myBool = true
	fmt.Println(myBool)
}
