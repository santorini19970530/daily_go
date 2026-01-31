// control_flow.go

package main

import (
	"fmt"
)

func main(){
	fmt.Println("Conditional Statments Example")
	RunConditionalExample()

	fmt.Println()

	fmt.Println("Looping Example")
	RunLoopingExample()

	fmt.Println()

	fmt.Println("Defer, Panic and Recover Example")
	RunDeferPanicRecoverExample()
}
