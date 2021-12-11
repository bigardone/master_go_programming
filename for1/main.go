package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// or
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// or
	i := 10
	for i >= 0 {
		fmt.Println(i)
		i--
	}

	// infinite loop
	// sum := 0
	// for {
	// 	sum++
	// }
	// fmt.Println(sum)
}
