package main

import "fmt"

const untypedConstant = 10
const typedConstant int = 10
const hello = "Hello world"
const l = len(hello)
const t = 1 < 2

func main() {
	var x int = untypedConstant
	fmt.Println(x)
	x = typedConstant
	fmt.Println(x)

	var y int8
	y = untypedConstant
	fmt.Println(y)
	// Typed constant need a cast if using different type
	y = int8(typedConstant)
	fmt.Println(y)

	fmt.Println(hello)
}
