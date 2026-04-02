package main

import "fmt"

func main() {
	type student struct {
		firstName string
		lastName  string
		age       int
		subjects  []string
	}

	var student1 student
	fmt.Println(student1)

	student1.firstName = "John"
	student1.lastName = "Smith"
	student1.age = 21
	student1.subjects = []string{"a", "bc"}
	fmt.Println(student1)

	var student2 student = student{
		firstName: "Paul",
		lastName:  "Jackson",
		age:       30,
		subjects:  []string{"Physics", "Mathematics"},
	}
	fmt.Println(student2)

	var student3 student = student{
		"a",
		"b",
		3,
		[]string{"c"},
	}
	fmt.Printf("%+v\n", student3)

	// Anonymous
	c := struct {
		firstName  string
		secondName string
	}{"a", "b"}
	fmt.Println(c)
	fmt.Printf("%+v\n", c)
}
