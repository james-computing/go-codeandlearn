package main

import "fmt"

func main() {
	var car Car = Car{
		Model:  "car model",
		Engine: Engine{Horsepower: 400},
		GPS:    GPS{Model: "gps model", Position: 15},
	}

	// Model is ambiguous, have to specify if it is from Car or GPS
	fmt.Println(car.Model)
	fmt.Println(car.GPS.Model)
	// Position isn't ambiguous, can access directly
	fmt.Println(car.Position)
}

type Engine struct {
	Horsepower int
}

type GPS struct {
	Model    string
	Position int
}

// Car with embedded Engine and GPS
type Car struct {
	Model string
	Engine
	GPS
}
