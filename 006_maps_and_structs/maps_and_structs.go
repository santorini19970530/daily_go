// maps_and_structs.go

package main

import (
	"fmt"
	"reflect"	// get the tag text
)

// go does not support OOP, like inheritance
// variables having capital letter in front can be accessed everywhere
// tag can be used to add comments on the variable, not a rule for the variable
type Animal struct{
	Name string	`required max:"100`
	Origin string
}

type Bird struct{
	Animal
	SpeedKPH float32
	CanFly bool
}

func main(){
	fmt.Println("Map Examples")
	RunMapExample()

	fmt.Println()
	fmt.Println("Struct Examples")
	RunStructExample()

	fmt.Println()
	fmt.Println("Inheritance Example")
	
	// using composition (has-a) to handle inheritance (is-a)
	// this is called embedding
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	// when declaring in first line, need to consider Animal
	/*
	b := Bird{
		Animal: Animal{Name : "Emu", Origin: "Australia"},
		SpeedKPH : 48,
		CanFly : false
	}
	*/

	// getting the tag text
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
