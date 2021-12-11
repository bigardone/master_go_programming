package main

import "fmt"

// import "fmt" // error
import f "fmt" // permitted in go

const done = false // package scoped

var b int = 10

func main() {
	var task = "running" // local (block) scoped
	fmt.Println(task, done)

	const done = true // local scoped
	f.Printf("done in main() is %v\n", done)
	f1()

	f.Println("Bye bye!")
}

func f1() {
	const done = true
	f.Printf("done in f1(): %v\n", done) // this is done from package scope
}
