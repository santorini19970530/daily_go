// numerical_complex_number.go

package main

import "fmt"

// RunIntegerDemo contains all the integers
func RunComplexDemo() {
	fmt.Println("Numerical primitive demo - Complex Numbers\n")	

	/* types of complex numbers
	complex64 , complex128
	the difference are 32 real + 32 imaginary or 64 real + 64 imaginary
	*/
	var a complex64 = 1 + 2i

	fmt.Printf("%v, %T\n", a, a)

	b := 2i

	fmt.Printf("%v, %T\n", b, b)

	// simple operations

	aa := 1 + 2i
	bb := 2 + 5.2i

	fmt.Println(aa + bb)
	fmt.Println(aa - bb)
	fmt.Println(aa * bb)
	fmt.Println(aa / bb)

	// getting the real or imaginary part

	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))

	fmt.Printf("%v, %T\n", real(b), real(b))
	fmt.Printf("%v, %T\n", imag(b), imag(b))

	// generate complex number

	var c complex128 = complex(5, 12)

	fmt.Printf("%v, %T\n", c, c)
}
