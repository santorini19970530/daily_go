// conditional.go

package main

import(
	"fmt"
	"math"
)

func RunConditionalExample(){
	// example 01
	if true{
		fmt.Println("This is true.")
	}

	// example 02
	if false{
		fmt.Println("This is false.")
	}

	// example 03
	// check if the key-value pair exist in map
	statePopulations := map[string]int{
		"Califonia" : 39250017,
		"Texas" : 27862596,
		"Florida" : 20612439,
		"New York" : 19745289,
		"Pennsylvania" : 12802503,
		"Illinois" : 12801539,
		"Ohio" : 11614373,
	}

	// if initializer; tester {execute}
	// get variable for validation, then check if the variable is ok
	// variable in the initializer is scope limited into the if block
	if pop, ok := statePopulations["Florida"] ; ok{
		fmt.Println(pop)
	}
	
	if pop, ok := statePopulations["Oho"] ; ok{
		fmt.Println(pop)
	}

	// example 04
	// number guessing game
	number := 50
	guess := 30

	if guess < number{
		fmt.Println("Too low")
	}
	if guess > number{
		fmt.Println("Too high")
	}
	if guess == number{
		fmt.Println("You got it!")
	}

	// there are three more operators
	fmt.Println(number <= guess, number >= guess, number != guess, !true)

	// logical operator
	// or
	if guess < 1 || guess > 100{
		fmt.Println("The guess needs to be between 1 and 100.")
	}

	// and
	if guess >=1 && guess <= 100{
		fmt.Println("The guess is between 1 and 100.")
	}

	// example 05
	// adding else

	if guess < 1 || guess > 100{
		fmt.Println("The guess needs to be between 1 and 100.")
	} else {
		if guess < number{
			fmt.Println("Too low")
		}
		if guess > number{
			fmt.Println("Too high")
		}
		if guess == number{
			fmt.Println("You got it!")
		}
	}

	// example 06
	// adding else if

	if guess < 1 {
		fmt.Println("The guess needs to be greater than 1.")
	} else if guess > 100 {
		fmt.Println("The guess needs to be less than 100.")
	} else {
		if guess < number{
			fmt.Println("Too low")
		}
		if guess > number{
			fmt.Println("Too high")
		}
		if guess == number{
			fmt.Println("You got it!")
		}
	}

	// example 06
	// check if two decimal places are true
	// due to floating point storage issues, directly comparing two floating points can sometimes be not realistic
	// instead, if one number is very near that another number (rounding up some several decimal places), these two numbers can be considered as the same
	num := 0.123
	
	if math.Abs(num / math.Pow(math.Sqrt(num), 2) - 1) < 0.001{
		fmt.Println("The same")
	} else {
		fmt.Println("DIfferent")
	}
	// 0.001 is error parameter (decimal places that can accpt the error)

	// example 06
	// switch statement
	switch 2{	// switch i := 2 + 3; i{} << this is also working
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3, 4, 5:
		// several cases can be included into one line
			fmt.Println("Three to Five")
		// cannot have diplicated cases
		// case 1, 6:
		default:
			fmt.Println("Not within one and five")
	}

	// example 07
	// switch with tagging, just like if-else if
	// this can allow overlap cases
	i := 10
	switch{
		case i <= 10:
			fmt.Println("i is less than or equal to 10")
			// as break has already be implemented in switch-case logic
			// break can be omitted (even though break can still be added)
			// if i is 5, to allow running the first and second cases, use fallthrough
			// but fallthrough is keeps running regardness whether the second case is true or not, so better not to use that
		case i <= 20:
			fmt.Println("i is less than or equal to 20")
		default:
			fmt.Println("i is greater than 20")
	}

	// example 08
	// type switch
	var m interface{} = 1

	switch m.(type){
		case int:
			fmt.Println("m is an int")
		case float64:
			fmt.Println("m is a float64")
		case string:
			fmt.Println("m is a string")
		case [3]int:
			fmt.Println("m is [3]int")
		default:
			fmt.Println("m is another type")
	}
}
