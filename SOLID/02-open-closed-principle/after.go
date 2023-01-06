package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// returns the area of the circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Square struct {
	length float64
}

// returns the area of the square
func (s Square) area() float64 {
	return s.length * s.length
}

type Shape interface {
	area() float64
}

type Calculator struct {
}

// returns the area sum of the given shapes
func (a Calculator) areaSum(shapes ...Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func main() {
	c := Circle{radius: 5}
	s := Square{length: 7}
	calc := Calculator{}

	fmt.Println("area sum:", calc.areaSum(c, s))
}
