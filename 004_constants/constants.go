// constants.go

package main

import (
	"fmt"
)

const c_const int16 = 27

const (
	d_const = iota
	// enum constant can share the same declare formula
	// so e_const and f_const are not necessary to declare
	e_const = iota
	f_const = iota
)

const (
	g_const = iota
)

const (
	error = iota
	cat
	dog
	snake
)

// convert file size
const (
    _  = iota // ignore first value by assigning to blank identifier
    KB = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

// access right control

const(
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main(){
	// having upper letters for const variables will become public variables
	// so need to keep the similar camal casing with the other variables
	const a_const int = 42

	fmt.Printf("%v, %T\n", a_const, a_const)

	// compiler will reject any re-assignment of const
	// a_const = 27

	const pi float64 = 3.14

	fmt.Printf("%v, %T\n", pi, pi)

	// const also cannot be assigned by runtime functions
	// const b_const float64 = math.Sin(1.57)

	const c_const int = 42

	// the inner declaration works, and the type can be changed
	fmt.Printf("%v, %T\n", c_const, c_const)

	// adding const and variable is working
	var c int = 11

	fmt.Printf("%v, %T\n", c_const + c, c_const + c)

	// enum constant

	fmt.Printf("%v, %T\n", d_const, d_const)
	fmt.Printf("%v, %T\n", e_const, e_const)
	fmt.Printf("%v, %T\n", f_const, f_const)

	// this will reset to zero
	// having two iota is very useful for ID comparison for two datasets

	fmt.Printf("%v, %T\n", g_const, g_const)

	// better to have error for zero , since every variable are preset to zero
	// this can prevent error

	var animal int = 1
	fmt.Printf("%v\n", animal == cat)

	// calculate file size

	fileSize := 4000000000.
	fmt.Printf("%.2f GB\n", fileSize / GB)

	// chack the admin roles
	
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope

	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters & roles == isHeadquarters)
}
