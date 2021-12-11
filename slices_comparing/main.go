package main

import "fmt"

func main() {
	// not initialised
	var n []int
	fmt.Println(n, n == nil)

	// initialised
	m := []int{}
	fmt.Println(m, m == nil)

	// slices can only be compared to nil
	a, b := []int{1, 2}, []int{1, 2, 3}
	// fmt.Println(a == b) // error

	// comparing slices
	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	// always compare lengths
	if len(a) != len(b) {
		eq = false
	}

	fmt.Println("a equal to b?", eq)
}
