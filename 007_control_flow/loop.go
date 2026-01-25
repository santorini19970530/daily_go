// loop.go

package main

import(
	"fmt"
)

func RunLoopingExample(){
	// for loop

	// example 01
	for i :=  0 ; i < 5 ; i++{
		fmt.Printf("%v\t", i)
	}
	fmt.Println()

	// example 02
	// declaring two looping counters at the same time
	// this is not working, since go does not have comma operator
	// and two sentences cannot be stayed into one line
	// for i := 0, j := 0 ; i < 5 ; i++, j++
	for i, j := 0, 0 ; i < 5 ; i, j = i + 1, j + 1 {
		fmt.Printf("%v %v\t", i, j)
	}
	fmt.Println()

	// example 03
	for i := 0 ; i < 5 ; i++{
		fmt.Printf("%v\t", i)

		if i%2 == 0{
			i /= 2
		} else {
			i = 2 * i + 1
		}
	}
	fmt.Println()

	// example 04
	// go does not have while loop
	// leaving the first and third input empty in for loop can do the same thing
	k := 0
	for ; k < 5 ; {
	// or just leave the middle input alone
	// for k < 5
		fmt.Printf("%v\t", k)
		k++
	}
	fmt.Println()

	// example 05
	// do...while loop
	m := 0
	for {
		fmt.Printf("%v\t", m)
		m++
		if m >= 5{
			break
		}
	}
	fmt.Println()

	// example 06
	// continue
	for i := 0 ; i < 10 ; i++{
		if i % 2 == 0{
			// using continue to skip the even numbers
			continue
		}
		fmt.Printf("%v\t", i)
	}
	fmt.Println()
	fmt.Println()

	// example 07
	// nested loop
	for i := 0; i <= 3; i++{
		for j := 0 ; j <=3; j++{
			fmt.Printf("%v * %v = %v\n", i, j, i * j)
		}
	}
	fmt.Println()

	// example 08
	// having a break in the inner for loop cannot break the outer for loop
	// need a label to let the break knows which loop to break
	Loop:
	for i := 0; i <= 3; i++{
		for j := 0 ; j <=3; j++{
			fmt.Printf("%v + %v = %v\n", i, j, i + j)

			if i + j >= 4{
				break Loop
			}
		}
	}
	fmt.Println()

	// example 09
	// looping through collections (slice and array)
	s := []int {1, 2, 3}
	for key, value := range s {
		fmt.Println(key, value)
	}
	fmt.Println()

	// example 10
	// looping through maps
	statePopulations := map[string]int{
		"Califonia" : 39250017,
		"Texas" : 27862596,
		"Florida" : 20612439,
		"New York" : 19745289,
		"Pennsylvania" : 12802503,
		"Illinois" : 12801539,
		"Ohio" : 11614373,
	}
	for key, value := range statePopulations{
		fmt.Println(key, value)
	}
	fmt.Println()

	// example 11
	// looping through elements of string
	greeting := "Hello World!"
	for key, value := range greeting{
		fmt.Println(key, string(value))
	}
}
