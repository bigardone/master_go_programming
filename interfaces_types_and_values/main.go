package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Permiter: %v\n", s.perimeter())
}

func main() {
	var s shape
	fmt.Printf("%T\n", s)

	ball := circle{radius: 2.5}
	s = ball
	print(s)
	fmt.Printf("Type of s: %T\n", s)

	room := rectangle{width: 2., height: 3.}
	s = room
	fmt.Printf("Type of s: %T\n", s)
}
