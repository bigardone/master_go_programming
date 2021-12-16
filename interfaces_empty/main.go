package main

import "fmt"

// empty interface, it represents any value
type emptyInterface interface{}

type Person struct {
	info interface{} // info can store ANY type
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{1, 2, 3}
	fmt.Println(empty)
	// fmt.Println(len(empty)) // compile error
	fmt.Println(len(empty.([]int)))

	you := Person{}
	you.info = "Your name"
	you.info = 40
	you.info = []float64{5.6, 4., 3.}
	fmt.Println(you.info)
}
