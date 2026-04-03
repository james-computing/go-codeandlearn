package main

import "fmt"

func main() {
	// empty interface
	var mixedSlice []interface{} = []interface{}{42, "Hello", 3.14, true}

	for _, value := range mixedSlice {
		fmt.Println(value)
	}
	fmt.Println()

	// type assertion
	var emptyInterface any = "Hello World!"
	str, ok := emptyInterface.(string)
	if ok {
		fmt.Println(str)
	}
	fmt.Println()

	// type switch
	for _, value := range mixedSlice {
		checkType(value)
	}
}

func printValue(value any) {
	fmt.Println(value)
}

func checkType(value interface{}) {
	// type switch
	switch value := value.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	case float64:
		fmt.Println("float64", value)
	default:
		fmt.Println("Unknown type")
	}
}
