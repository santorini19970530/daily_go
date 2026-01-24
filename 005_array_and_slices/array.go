// array.go

package main

import (
	"fmt"
)

// RunArrayExample - illustration of array

func RunArrayExample(){
	// creating an array starts with size of array,
	// then the data type of array elements
	// then the array elements
	grades := [3]int{97, 85, 93}

	fmt.Printf("Grades : %v\n", grades)

	// array size with three dots means just create an array with the elements being decided
	gradesA := [...]int{96, 99, 83, 21}

	fmt.Printf("Grades A : %v\n", gradesA)

	// creating an empty array, then assign the values
	var students [3]string

	fmt.Printf("Students before assign : %v\n", students)

	students[0] = "Eunha"
	students[1] = "Sowon"
	students[2] = "SinB"

	fmt.Printf("Students after assign : %v\n", students)
	fmt.Printf("Student One : %v - %v\n", students[0], grades[0])

	// len() - get the size of array
	fmt.Printf("Number of Students : %v\n", len(students))

	// array of arrays
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1}}

	fmt.Println(identityMatrix)

	// array in other languages are pointing to the memory storage
	// in go, array is the storage of the values
	// array b is being stored into different memory allocation as a

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5

	fmt.Printf("a - %v\n", a)
	fmt.Printf("b - %v\n", b)
	fmt.Println()

	// in go, pointer will be pointing to the original allocation

	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 5

	fmt.Printf("c - %v\n", c)
	fmt.Printf("d - %v\n", d)
}
