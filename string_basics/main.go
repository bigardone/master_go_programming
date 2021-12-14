package main

import "fmt"

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello!\"")

	// Raw strings
	fmt.Println(`He says: "Hello!"`)

	s2 := `I like \n Go!` // No new line
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")

	// Multi-line
	fmt.Println(`
Price: 100
Brand: Nike
  `)

	fmt.Println(`C:\Users\Ricardo`)
	fmt.Println("C:\\Users\\Ricardo")

	// Concatenation
	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	// Indexing
	fmt.Println("Element at index 0", s3[0])

	// Strings are inmutable
	// s3[5] = "x" // error

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)
}
