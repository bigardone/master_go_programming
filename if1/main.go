package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too expensive")
	}

	// _ = inStock

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	// if price {
	//     fmt.Println("something")
	// }

	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's expensive!")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years", 18-age)
	} else if age == 18 {
		fmt.Println("This is your first vote")
	} else if age > 18 && age <= 100 {
		fmt.Println("Please vote")
	} else {
		fmt.Printf("Invalid age!")
	}
}
