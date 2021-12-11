package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			// start next iteration in the loop
			continue
		}
		fmt.Println(i)
	}
}
