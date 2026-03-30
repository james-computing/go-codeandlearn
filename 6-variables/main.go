package main

import "fmt"

var (
	packageLevelInt    int    = 5
	packageLevelString string = "hello"
)

func main() {
	fmt.Println(packageLevelInt)
	fmt.Println(packageLevelString)

	functionLevelInt, functionLevelString := 6, "world"
	fmt.Println(functionLevelInt)
	fmt.Println(functionLevelString)

	var x, y = 7, "again"
	fmt.Println(x, y)
}
