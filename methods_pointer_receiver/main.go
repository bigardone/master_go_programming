package main

import "fmt"

type Car struct {
	brand string
	price int
}

func changeCar(c Car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c Car) changeCar1(newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c *Car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

// type distance *int
//
// func (d *distance) m1() {
// 	fmt.Println("just a message")
// }

func main() {
	myCar := Car{brand: "Auid", price: 40000}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar)

	myCar.changeCar2("Seat", 20000)
	fmt.Println(myCar)

	var yourCar *Car
	yourCar = &myCar
	fmt.Println(*yourCar)

	// valid ways
	yourCar.changeCar2("VW", 30000)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("Another VW", 30000)
	fmt.Println(*yourCar)
	fmt.Println(myCar)
}
