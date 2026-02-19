// interfaces_example02.go

package main

import(
	"fmt"
)

func RunInterfacesExample02(){
	a := IntCounter(0)

	var inc Incrementer = &a

	for i := 0; i < 10; i++{
		fmt.Println(inc.Increment())
	}
}

// there is no need to use struct as interface
type Incrementer interface{
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int{
	*ic++
	
	return int(*ic)
}
