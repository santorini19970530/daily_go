// text.go

package main

import "fmt"

// contains all the text operation
func RunTextDemo() {
	fmt.Println("Text demo\n")	

	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("This is the third character of string - %v, %T\n", s[2], s[2])

	// but we cannot directly change the assignment of the character
	// s[2] = "u"
	// however, string concatenation is working
	s2 := "this is also a string"

	fmt.Printf("%v, %T\n", s+" "+s2, s+s2)

	// convert string into bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	/* string is representing utf8 characters
		rune is represenging utf32 characters 
	*/

	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	// that will show int32 as rune type
}
