// pointers.go

package main

import (
	"fmt"
)

func main(){
	RunPointerExample01()

	RunPointerExample02()

	RunPointerExample03()

	RunPointerExample04()
}

// example 01
func RunPointerExample01(){
	var a int = 42
	var b int = a
	fmt.Println(a, b)

	a = 27	// when a changes, b will not change at the same time
	fmt.Println(a, b)

	var c *int = &a
	// this means c is an integer pointer, pointing the the address which is storing a
	fmt.Printf("The address of a : %v, %v.\n", &a, c)
	fmt.Printf("The value of a : %v, %v\n", a, *c)

	// when a changes the value, c will change the same time
	a = 41
	fmt.Println(a, b, c)
}

// example02
func RunPointerExample02(){
	a := [3]int{1, 2, 3}

	b := &a[0]
	c := &a[1]

	fmt.Printf("%v , %p , %p\n", a, b, c)

	// go cannot use address to minus the int
	// fmt.Println(&a[1] - 4)
	// but this line is showing that int array elements are having difference of 4
}

// example03
func RunPointerExample03(){
	var stru *fooStruct
	stru = &fooStruct{foo: 46}
	fmt.Println(stru)

	// de-referencing the pointer
	(*stru).foo = 42
	fmt.Println((*stru).foo)
	// below is the same
	// stru,foo = 42
	// fmt.Println(stru.foo)
}

type fooStruct struct{
	foo int 
}

// example04
func RunPointerExample04(){
	// array is actually keeping the space and value of array
	// the assignment is in fact duplicating the array
	Example04Array()
	// slice is manipulating the array
	// just like using linked list, when there would be new element the assignment will go next
	// so having an assignment means there is reference to the another element
	Example04Slice()
	// map is also having this issue
	Example04Map()
	// other data types does not have such issue
}

func Example04Array(){
	a := [3]int{1, 2, 3}

	b := a

	fmt.Println(a, b)

	a[1] = 42
	fmt.Println(a, b)
}

func Example04Slice(){
	a := []int{1, 2, 3}

	b := a

	fmt.Println(a, b)

	a[1] = 42
	fmt.Println(a, b)
}

func Example04Map(){
	a := map[string]string {"foo" : "bar", "cat" : "jazz"}

	b := a

	fmt.Println(a, b)

	a["foo"] = "qux"

	fmt.Println(a, b)
}
