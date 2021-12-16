package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d before sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c)

	fmt.Println("main goroutine sleeps fot 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c {
		fmt.Println("main goroutine received data:", v)
	}

	fmt.Println(<-c)
}
