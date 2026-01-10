// boolean.go

package main

import "fmt"

// RunBooleanDemo contains all the boolean demonstration
func RunBooleanDemo() {
	fmt.Println("Boolean primitive demo\n")

	var n bool = true

	fmt.Printf("%v, %T\n", n, n)

	n = false

	fmt.Printf("%v, %T\n", n, n)

	// for Go, whenever there has been a variable initialised
	// the first assigned value will be zero value

	var m bool

	fmt.Printf("%v, %T\n", m, m)
}

