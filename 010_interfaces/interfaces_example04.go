// interfaces_example04.go
// type switch

package main

import(
	"fmt"
)

func RunInterfacesExample06(){
	// inputting different types will show different results of Println
	var i interface{} = 0

	switch i.(type){
		case int:
			fmt.Println("i is an integer")
		case string:
			fmt.Println("i is a string")
		default:
			fmt.Println("I don't know what i is")
	}
}
