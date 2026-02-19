// interfaces.go

package main

import(
	"fmt"
)

func main(){
	RunInterfacesExample01()

	fmt.Println()

	RunInterfacesExample02()

	fmt.Println()

	// example 3 - 5 in example03 go file
	RunInterfacesExample03()

	fmt.Println()

	RunInterfacesExample04()

	fmt.Println()

	RunInterfacesExample05()

	fmt.Println()

	RunInterfacesExample06()
}

/* implementing with values vs pointers

* method set of value is all methods with value receivers

* method set of pointer is all methods, regardless of receiver type

*/

/* practices of using interfaces

* use many, small interfaces

* do not export interfaces for types that will be consumed

* do export interfaces for types that will be used by package

* design functions and nethods to receive interfaces whenever possible

*/
