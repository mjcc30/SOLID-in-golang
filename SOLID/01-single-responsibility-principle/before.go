package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// returns the area of the circle with the given radius.
func (c Circle) area() {
	fmt.Printf("circle area: %f\n", math.Pi * c.radius * c.radius)
}

type Square struct {
	length float64
}

// returns the area of the square with the given length.
func (s Square) area() {
	fmt.Printf("square area: %f\n", s.length * s.length)
}

func main() {
	c := Circle{radius: 5}
	c.area()

	s := Square{length: 7}
	s.area()
}
