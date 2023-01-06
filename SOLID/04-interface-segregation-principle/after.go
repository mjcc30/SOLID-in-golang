package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
}

// returns the area of the square
func (s Square) area() float64 {
	return math.Pow(s.length, 2)
}

type Cube struct {
	Square
}

func (c Cube) volume() float64 {
	return math.Pow(c.length, 3)
}

type Shape interface {
	area() float64
}

type Object interface {
	Shape
	volume() float64
}

// returns the sum of all shapes's area
func areaSum(shapes ...Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

// returns the sum of the volume and area of all object
func volumeSum(shapes ...Object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s1 := Square{length: 3}
	c1 := Cube{Square: Square{length: 4}}
	fmt.Println("area sum: ",areaSum(s1, c1))
	s2 := Square{length: 6}
	c2 := Cube{Square: Square{length: 4}}
	fmt.Println("area sum: ",areaSum(s1, s2, c1, c2))
	fmt.Println("volume sum: ",volumeSum(c1, c2))
}
