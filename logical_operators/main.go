package main

import "fmt"

func main() {
	// && logical and
	a, b := 5, 10
	fmt.Println(a > 1 && b == 10)
	fmt.Println(a < 1 && b == 10)

	// || logical or
	fmt.Println(a == 5 || b == 100)
	fmt.Println(a != 5 || b == 100)

	// ! logical negation
	fmt.Println(!(a > 0))
	fmt.Println(!(a == 1) || (b == 100))
}
