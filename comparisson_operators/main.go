package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(a > 5, a >= 5)
	fmt.Println(b < a, 10 <= 5)
}
