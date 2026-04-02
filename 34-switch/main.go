package main

import "fmt"

func main() {
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("Something else")
	}

	switch i := 0; i {
	case 0:
		fmt.Println("zero")
	default:
		fmt.Println("non zero")
	}

	switch i := 0; {
	case i == 0:
		fmt.Println("i = 0")
	}

	i := 0
	switch i {
	case 0:
		fmt.Println("0")
		fallthrough
	case 1:
		fmt.Println("0 or 1")
	}

	j := 0
	switch j {
	case 0, 1:
		fmt.Println(j)
	}
}
