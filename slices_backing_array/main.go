package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600
	fmt.Println(s1)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2)

	/////////////////////////

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	sl1 := []int{10, 20, 30, 40, 50}
	newSlice := sl1[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = sl1[2:5] // {30, 40, 50}
	fmt.Println(len(newSlice), cap(newSlice))

	// Print the backing array header address
	fmt.Printf("%p\n", &sl1)
	fmt.Printf("%p %p \n", &sl1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("sl1:", sl1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("array's size in bytes: %d\n", unsafe.Sizeof(a))
	fmt.Printf("slice's size in bytes: %d\n", unsafe.Sizeof(s))
}
