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
func areaCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

type Square struct {
	length float64
}

// returns the name of the square
func (s Square) name() string {
	return "square"
}

// returns the area of the square with the given length.
func areaSquare(length float64) float64 {
	return length * length
}

type Calculator struct {
}

type Shape interface {
	name() string
}

// returns the area sum of the given shapes
func (a Calculator) areaSum(shapes ...Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		switch shape.name() {
		case "circle":
			r := shape.(Circle).radius
			sum += areaCircle(r)
		case "square":
			l := shape.(Square).length
			sum += areaSquare(l)
		}
	}
	return sum
}

func main() {
	c := Circle{radius: 5}
	s := Square{length: 7}
	calc := Calculator{}

	fmt.Println("area sum:", calc.areaSum(c, s))
}
