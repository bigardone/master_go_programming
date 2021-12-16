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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Permiter: %v\n", s.perimeter())
}

func main() {
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	// s.volume() // compile error, can't access volume() from the interface value
	// type assertion
	ball, ok := s.(circle)

	if ok == true {
		fmt.Printf("Ball volume: %v\n", ball.volume())
	}

	s = rectangle{width: 3.4, height: 2.2}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has a circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has a rectangle type\n", value)
	}
}
