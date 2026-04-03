package main

import (
	"fmt"
)

func main() {
	message := "Hello, "

	greetingFunc := func(name string) {
		fmt.Println(message + name)
	}

	defer greetingFunc("Alice")
	defer greetingFunc("Bob")

	message = "Hi, "

	// If use os.Exit, the function doesn't return, so the defer isn't executed
	//os.Exit(0)
}
