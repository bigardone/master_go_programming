package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255 // max value
	x++               // overflow, x set to 0
	fmt.Println(x)

	// a := int8(255 + 1)
	var b int8 = 127 // max value
	fmt.Printf("%d\n", b+1)

	b = -128
	b--
	fmt.Println(b)

	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f = f * 1.2 // floats overflow to infinite!!
	fmt.Println(f)

	// const i int8 = 300
}
