package main

import "fmt"

func main() {
	var myString string = "Hello\nWorld"
	fmt.Println(myString)
	myString =
		`Multi
line
string
`
	fmt.Println(myString)

	var firstName, lastName string
	firstName = "John"
	lastName = "Smith"
	fmt.Printf("%s %s\n", firstName, lastName)

	var fullName string = firstName + " " + lastName
	fmt.Println(fullName)

	firstName = "Albert"
	lastName = "Einstein"
	fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)
}
