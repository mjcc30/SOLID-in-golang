package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// returns the name of the circle
func (c Circle) name() string {
	return "circle"
}

// returns the area of the circle with the given radius.
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Square struct {
	length float64
}

// returns the name of the square
func (s Square) name() string {
	return "square"
}

// returns the area of the square with the given length.
func (s Square) area() float64 {
	return s.length * s.length
}

type Shape interface {
	area() float64
	name() string
}

type printer struct {
}

func (printer) info(s Shape) {
	fmt.Printf("%s area: %f\n", s.name(), s.area())
}

func main() {
	c := Circle{radius: 5}
	s := Square{length: 7}

	p := printer{}
	p.info(c)
	p.info(s)
}
