package main

import "fmt"

func main() {
	name := "John Smith"

	firstLetter := string(name[0])
	fmt.Println(firstLetter)

	var x rune = 'x'
	fmt.Println(string(x))

	var y byte = 'y'
	fmt.Println(string(y))

	unicodeString := "hello 😀 world"
	fmt.Println(unicodeString)
	fmt.Printf("unicode string length = %v\n", len(unicodeString))
	bytes := []byte(unicodeString)
	fmt.Println(bytes)
	fmt.Printf("bytes length = %v\n", len(bytes))
	runes := []rune(unicodeString)
	fmt.Println(runes)
	fmt.Printf("runes length = %v\n", len(runes))

	// Slicing the string get the bytes, which can go wrong
	firstHalf := unicodeString[:len(unicodeString)/2]
	fmt.Println(firstHalf)
	secondHalf := unicodeString[len(unicodeString)/2:]
	fmt.Println(secondHalf)
	fmt.Println("***********************************")

	// Use rune to get emoji
	var emoji string = string(runes[6])
	fmt.Println(emoji)

	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
	}
}
