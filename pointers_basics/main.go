package main

import "fmt"

func main() {
	name := "Andrei"
	fmt.Println(&name)

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T with value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("Address of x is %p\n", &x)

	var ptr1 *float64 // pointer to a type
	_ = ptr1

	p := new(int)
	x = 100
	p = &x
	fmt.Printf("p is of type %T with a value of %v\n", p, p)
	fmt.Printf("address of x is %p\n", &x)

	*p = 90 // pointers value, equivalent to x = 90
	fmt.Println(x, *p)
	fmt.Println("*p==x:", *p == x)

	*p = 10     // x = 10
	*p = *p / 2 // x / 2
	fmt.Println(x)

	// &value returns pointer
	// *pointer reuturns value
}
