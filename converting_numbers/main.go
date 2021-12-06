package main

import "fmt"

func main() {
	var x = 3   // int type
	var y = 3.1 // float type

	// x = x * y
	x = x * int(y)
	fmt.Println(x)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	x = int(float64(x) * y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5
	fmt.Printf("%T\n", a)

	var b int64 = 2
	a = int(b)
}
