package main

import (
	"fmt"
	_ "init-example/init"
)

func init() {
	fmt.Println("init 1 called")
}

func init() {
	fmt.Println("init 2 called")
}

func init() {
	fmt.Println("init 3 called")
}

func main() {
	fmt.Println("main called")
}
