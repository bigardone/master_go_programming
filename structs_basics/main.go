package main

import "fmt"

func main() {
	title1, author1, year1 := "The Divine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"The Divine Comedy", "Dante Aligheri", 1320} // not recommended
	fmt.Println(myBook)

	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945} // recommended way
	_ = bestBook

	aBook := book{title: "Random book"}
	fmt.Printf("%#v\n", aBook)

	// anonymous struct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's age: %d\n", diana.age)

	// Anonymous fields
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by gGeorge Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)
	fmt.Println(b1.string)

	// Mix both named and anonymous fields
	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\v", e)
	e.bool = true
}
