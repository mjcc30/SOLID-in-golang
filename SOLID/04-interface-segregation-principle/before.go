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

// square is a flat surface -> does not have volume
func (s Square) volume() float64 {
	return 0
}

type Cube struct {
	length float64
}

// returns the area of the cube
func (c Cube) area() float64 {
	return math.Pow(c.length, 2)
}

// returns the volume of the cube
func (c Cube) volume() float64 {
	return math.Pow(c.length, 3)
}

type Shape interface {
	area() float64
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

// returns the sum of the volume and area of all shapes
func volumeSum(shapes ...Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.volume()
	}
	return sum
}

func main() {
	s1 := Square{length: 3}
	c1 := Cube{length: 4}
	fmt.Println("area sum: ",areaSum(s1, c1))
	fmt.Println("volume sum: ",volumeSum(s1, c1))
}
