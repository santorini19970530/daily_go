// slices.go

package main

import (
	"fmt"
)

// RunSlicesExample - illustration of slices

func RunSlicesExample(){
	// slice is array without setting fixed size
	a := []int{1, 2, 3}
	fmt.Println(a)

	// everything which can be done in array can also handle slice

	// except some functions
	// e.g. cap() - shows capacity of the slice
	// since not every cells of slice is necessary being filled in
	fmt.Printf("len of a - %v\tcapacity of a - %v\n", len(a), cap(a))

	// different from array, slice is storing values by reference types
	// below example does not show the pointer reference parathensis, but a and b a referencing the the same set of memories
	b := a
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)

	// slice can just reference to specific position of another slice
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]	// slice all elements
	e := c[3:]	// slice starting from the forth element
	f := c[:6]	// slice first six elements
	g := c[3:6]	// slice the forth to the sixth elements

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// make(data type, length, capacity) - create slice using function
	
	h := make([]int, 3)
	fmt.Println(h)
	fmt.Printf("len of h - %v\tcapacity of h - %v\n", len(h), cap(h))

	m := make([]int, 3, 10)
	fmt.Println(m)
	fmt.Printf("len of m - %v\tcapacity of m - %v\n", len(m), cap(m))

	// slice does not need to be filled in in the beginning
	// when there will be elements to be added, use append() to add elements on the end

	m = append(m, 1)
	fmt.Println(m)
	fmt.Printf("len of m - %v\tcapacity of m - %v\n", len(m), cap(m))

	// when there is not enough capacity
	// go will copy all of the elements of the original slice to another memory that is having doubled space as the original allotment
	h = append(h, 2)
	fmt.Println(h)
	fmt.Printf("len of h - %v\tcapacity of h - %v\n", len(h), cap(h))

	// there is no function to remove slice element, but just need to use another slice to extract the wanted elements into a new slice
	// e.g. to remove the middle element
	p := append(c[:2], c[3:]...)
	fmt.Println(p)
}
