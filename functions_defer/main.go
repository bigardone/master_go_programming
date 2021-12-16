package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func fooBar() {
	fmt.Println("This is fooBar()")
}

func main() {
	defer foo() // executes it after exiting the current function
	bar()

	fmt.Println("Just a string after defering foo() and calling bar()")

	defer fooBar()

	file, err := os.Open("main.go")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// code that works (read/write) with the file
}
