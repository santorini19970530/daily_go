// numerical_integer.go

package main

import "fmt"

// contains all the integers
func RunIntegerDemo() {
	fmt.Println("Numerical primitive demo - Integer\n")	

	// int

	/* types of integers
	int8 , int16 , int32 , int64
	*/
	n := 42 
	fmt.Printf("%v, %T\n", n, n)

	// unsigned integer cannot have negative numbers
	// ./numerical.go:19:18: cannot use -42 (untyped int constant) as uint16 value in variable declaration (overflows)
	/* types of unsigned integers
	uint8 , uint16 , uint32
	*/
	var un uint16 = 42
	fmt.Printf("%v, %T\n", un, un)

	// simple operations

	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// different types cannot add together
	// ./numerical.go:41:25: invalid operation: aaa + bbb (mismatched types int and int8)
	var aaa int = 10
	var bbb int8 = 3

	//fmt.Printf("%v, %T\n", aaa + bbb, aaa + bbb)
	fmt.Printf("%v, %T\n", aaa + int(bbb), aaa + int(bbb))

	// bit operations

	fmt.Println(a & b)	// and operator
	fmt.Println(a | b)	// or operator
	fmt.Println(a ^ b)	// exclusive or operator
	fmt.Println(a &^ b)	// and not operator

	// bit shifting
	aa := 8

	fmt.Println(aa << 3)
	fmt.Println(aa >> 3)
}

