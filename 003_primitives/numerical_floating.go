// numerical_floating.go

package main

import "fmt"

// RunIntegerDemo contains all the integers
func RunFloatingDemo() {
	fmt.Println("Numerical primitive demo - Floating Points\n")	

	/* types of floating points
	float32 , float64
	*/
	var a float32 = 3.14

	fmt.Printf("%v, %T\n", a, a)

	b := 13.7e72

	fmt.Printf("%v, %T\n", b, b)

	c := 2.1E14

	fmt.Printf("%v, %T\n", c, c)

	// simple operations

	aa := 10.2
	bb := 3.7

	fmt.Println(aa + bb)
	fmt.Println(aa - bb)
	fmt.Println(aa * bb)
	fmt.Println(aa / bb)
}
