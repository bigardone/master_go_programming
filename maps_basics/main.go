package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("Number of pairs: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m map[[5]int]string
	_ = m

	// employees["Dan"] = "Programmer" // Runtime error

	people := map[string]float64{}
	people["John"] = 21.4
	people["Mary"] = 10
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 3343.34,
		"EUR": 45345.44,
		"CHF": 3434.34,
	}

	fmt.Println(balances)

	m2 := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m2

	balances["USD"] = 500.3
	balances["GBP"] = 200
	fmt.Println(balances)

	fmt.Println(balances["RON"]) // missing or zero value? check below

	balances["RON"] = 0

	v, ok := balances["RON"] // returns value and if exists or not
	if ok {
		fmt.Println("The RON key balance is:", v)
	} else {
		fmt.Println("The RON key does not exist!")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)
}
