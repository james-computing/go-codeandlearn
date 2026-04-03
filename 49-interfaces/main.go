package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// The interface accepts a value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// The interface also accepts a pointer receiver
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	var r Shape = Rectangle{
		Width:  2,
		Height: 3,
	}
	fmt.Println(r.Area())

	var c *Circle = &Circle{
		Radius: 1,
	}
	fmt.Println(c.Area())
}
