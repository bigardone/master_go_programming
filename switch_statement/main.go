package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "python":
		fmt.Println("You are learning Python...")

	case "Go", "golang":
		fmt.Println("Good, you are learning Go!")

	default:
		fmt.Println("You are learning any other language")
	}

	n := 5
	switch {
	case n%2 == 0:
		fmt.Println("Even number")
	case n%2 != 0:
		fmt.Println("Odd number")
	default:
		fmt.Println("Never here")
	}

	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
