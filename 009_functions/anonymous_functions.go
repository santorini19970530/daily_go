// anonymous_functions.go

package main

import(
	"fmt"
)

func RunAnonymousFunctions(){
	// example 11
	// the first pair of brankets insert parameters
	// the second pair is to invoke the function
	func () {
		fmt.Println("Hello World")
	}()

	// example 12
	// running in loop
	for i := 0; i < 5; i++{
		// passing the variable into parameter, not just referencing the variable
		// this can avoid async functions giving out undesired result
		func(i int){
			fmt.Println(i)
		}(i)
	}

	// example 13
	// function can be acting as variable
	f := func(){
		fmt.Println("Hello World")
	}
	/* this also works
	var f func() = func(){
		fmt.Println()
	}
	*/
	f()

	// example 14
	// declare the function first, fill in the commands later
	var divide func(float64, float64) (float64, error)

	divide = func(a, b float64) (float64, error){
		if b == 0.0{
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}

	d, err := divide(5.0, 0.0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}

