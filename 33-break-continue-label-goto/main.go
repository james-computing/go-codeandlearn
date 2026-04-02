package main

import "fmt"

func main() {
	// break
	a := 100
	for {
		fmt.Println(a)
		a--
		if a == 90 {
			fmt.Println("break loop")
			break
		}
	}
	fmt.Println()

	// continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("continue to next iteration")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// continue with label
outLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("continue to outLoop")
				continue outLoop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println()

	// goto
	b := 10

	if b == 10 {
		fmt.Println("jumping Println")
		goto end
	}
	fmt.Println(b)
end:
	fmt.Println("end")
	fmt.Println()

	// goto loop
	i := 0
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
}
