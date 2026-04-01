package main

import "fmt"

func main() {
	var nameAge map[string]int
	fmt.Println(nameAge["john"])
	fmt.Println(nameAge)
	fmt.Println(nameAge == nil)
	nameAge = make(map[string]int)
	fmt.Println(nameAge == nil)
	fmt.Println(nameAge)
	fmt.Printf("length = %v\n", len(nameAge))
	nameAge["john"] = 24
	fmt.Printf("length = %v\n", len(nameAge))
	fmt.Println(nameAge)

	var anotherMap map[string]int = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(anotherMap)
	fmt.Println(len(anotherMap))

	var age int
	var ok bool
	age, ok = anotherMap["d"]
	if !ok {
		fmt.Println("No value")
	} else {
		fmt.Printf("age = %v\n", age)
	}

	a := map[string][]int{
		"a": {1, 2, 3},
	}
	fmt.Println(a)
}
