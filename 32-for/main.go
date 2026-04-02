package main

import "fmt"

func main() {
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}

	b := 0
	for {
		fmt.Println(b)
		b++
		if b == 5 {
			break
		}
	}

	for i := 4; i >= 0; i-- {
		fmt.Println(i)
	}

	c := []int{1, 2, 3, 4, 5}
	for index, value := range c {
		fmt.Println(index, value)
	}

	s := "hello😀"
	for index, runeValue := range s {
		fmt.Printf("index = %d, char = %c\n", index, runeValue)
	}

	ages := map[string]int{
		"John": 20,
		"Paul": 24,
	}
	for name, age := range ages {
		fmt.Printf("name = %s, age = %d\n", name, age)
	}
}
