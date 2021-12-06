package main

import "fmt"

func main() {
	a, b := 4, 2
	r := (a + b) / (a - b) * 2
	fmt.Println(r)

	r = 9 % a
	fmt.Println(r)
}
