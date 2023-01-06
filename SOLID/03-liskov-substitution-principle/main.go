package main

import "fmt"

type Shape struct {
	name string
}

func (s Shape) getName() string {
	return s.name
}

type Circle struct {
	Shape
	radius float64
}

type Square struct {
	Shape
	length float64
}

type Object interface {
	getName() string
}

type Printer struct {
}

func (Printer) info(o Object) {
	fmt.Println("Name: ", o.getName())
}

func main() {
	sh := Shape{name: "Shape"}
	sq := Square{
		Shape:  Shape{name: "Circle"},
		length: 7,
	}
	c := Circle{
		Shape:  Shape{name: "Square"},
		radius: 5,
	}

	p := Printer{}
	p.info(sh)
	p.info(sq)
	p.info(c)
}
