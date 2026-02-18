// functions.go

package main

import (
	"fmt"
)

// main is already a function
// to initiate function, need to declare this as a func
// then have the name of the function
// and the parameter in the brackets
// finally input commands into the curly brackets
// (the opening curly bracket needs to be the same line as function declaration line)
func main(){
	// example 01
	fmt.Println("Hello World")

	// example 02
	sayMessage("Hello World")

	// example 03
	for i := 0; i < 5; i++{
		sayMessageA("Hello World", i)
	}

	// example 04
	sayGreeting("Hello", "Eunha")

	// example 05
	// as function is passing by value, the parameter and the function will be duplicated and processed
	// passing by pointer can ensure the variable pointing to the same capacity
	// this technique does not apply for maps and slices
	greeting := "Good morning"
	name := "Sowon"
	// this is before changed by pointer
	sayGreeting(greeting, name)

	sayGreetingPointer(&greeting, &name)

	// this is after changed by pointer
	sayGreeting(greeting, name)

	// example 06
	// variable parameters
	sum(1, 2, 3, 4, 5)
	// if there is another parameter in another type, only one is allowed and this needs to be in front of the variable parameters
	// sum ("The sum is", 1, 2, 3, 4, 5)

	// example 07
	s := sum_return(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is", s)

	// example 08
	// using pointer will come out the same result
	s1 := sum_return_pointer(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("The sum is", *s1)

	// example 09
	// named return value, declare the return variable and datatype in prior
	s2 := sum_named_return(1, 2, 3, 4)
	fmt.Println("The sum is", s2)

	// example 10
	// multiple return variables
	// this can avoid raising up DivideZeroError
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The result is", d)
	d, err = divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		// the whole program should have been terminated
		// just to go on the other examples, return has been hide
		// return
	}
	fmt.Println("The result is", d)

	RunAnonymousFunctions()

	RunMethods()
}

func sayMessage(msg string){
	// this msg is local variable
	fmt.Println(msg)
}

func sayMessageA(msg string, indx int){
	fmt.Println(msg)
	fmt.Printf("The value of index is %v.\n", indx)
}

// if the type of the parameters are the same, the previous datatype can be skipped
func sayGreeting(msg, name string){
	fmt.Println(msg, name)
}

func sayGreetingPointer(msg, name *string){
	fmt.Println(*msg, *name)
	*msg = "Good afternoon"
	*name = "Umji"
	fmt.Println(*msg, *name)
}

// inputing several inputs of same datatype together
func sum(values ... int){
	fmt.Printf("Values being input %v.\n", values)

	result := 0
	for _, v := range values{
		result += v
	}
	fmt.Println("The sum is", result)
}

// need to mention what the type of return values will be in the first line
func sum_return(values ... int) int{
	result := 0

	for _, v := range values{
		result += v
	}
	
	return result
}

func sum_return_pointer(values ... int) *int{
	result := 0

	for _, v := range values{
		result += v
	}

	// this will assign the memory that result is pointing to the main func
	return &result
}

func sum_named_return(values ... int) (result int){
	result = 0

	for _, v := range values{
		result += v
	}

	return
}

func divide(a, b float64) (float64, error){
	/* adding an panic will stop the function, which is not recommended
	if b == 0.0 {
		panic("Cannot provide zero as second value")
	}
	*/
	// but adding an error statement as second parameter, and adding another nonsence value to complete the return, can continue the program
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot provide zero as second value")
	}

	// adding nil when there is no error occurred
	return a / b, nil
}
