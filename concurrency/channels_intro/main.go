package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int) // stores an address
	fmt.Println(ch)

	c := make(chan int)

	// <- channel operator

	// SEND
	c <- 10

	// RECEIVE
	num := <-c
	_ = num
	fmt.Println(<-c)

	close(c)
}
