package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 40000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Streen 20, London",
			phone:   00321234567,
		},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@company.com"
	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "bigardone@domain.com", address: "My address", phone: 123456}
	fmt.Println(myContact)
}
