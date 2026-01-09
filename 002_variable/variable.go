// variable.go

package main

import("fmt")

// the innerst declaration will be covering this value (shadowing)
var i int = 72

func main(){
	// three ways to define variables in GO

	// method 1 - just declare the variable and use
	
	var i int

	i = 42

	fmt.Println(i)

	// method 2 - combine the declare and assignment into one line

	var j int = 21

	fmt.Println(j)

	// method 3 - let GO decide what the data type is

	k := 7

	fmt.Println(k)

	// to print a line with formatting

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// method 3 will be having less control on variable type

	x := 7.
	var y float32 = 7

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

	// cannot re-declare the same variable twice

	// declare variable by brunch

	var (
		actorName string = "Eunha"
		actressName string = "Sowon"
		doctorNumber int = 6
		season int = 11
	)

	var (
		counter int = 0
	)

	fmt.Printf("%s, %s, %v, %v, %v\n", actorName, actressName, doctorNumber, season, counter)

	// variables not being used will cause error, so need to use up

	// variables with lower case first letter will be used for package scope
	// variable names with first capital letter will be for export
	// there is not private scope, but can include the variable within a block

	// converting variable type

	var z int = 1
	var a float32
	a = float32(z)

	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("%v, %T\n", a, a)

	// use strconv package for strings
}
