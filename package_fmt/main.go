package main

import "fmt"

func main() {
	name := "Ricardo"
	fmt.Println("Hello, I'm", name)

	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean value:", (a+b)/2)

	fmt.Printf("Your age is %d\n", 21)

	fmt.Printf("x is %d, y is %f\n", 5, 6.8)

	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of %s with a radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	// %q for quoted strings
	fmt.Printf("This is a %q\n", figure)

	// %v replaced by any type
	fmt.Printf("The diameter of %v with a radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	// %T to print types
	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	// %t format a boolan value
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// %b convert to base 2

	fmt.Printf("%b\n", 55)
	fmt.Printf("%08b\n", 55)

	fmt.Printf("%x\n", 100)

	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %f\n", x*y)
	fmt.Printf("x * y = %.3f\n", x*y)
}
