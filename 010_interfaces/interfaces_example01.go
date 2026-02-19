// interfaces_example01.go

package main

import(
	"fmt"
)

func RunInterfacesExample01(){
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Eunha"))
}

// interface describes only behaviors, but not data
type Writer interface{
	// just write to something, but do not know what to be written
	// int output is the number of bytes being output
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error){
	n, err := fmt.Println(string(data))
	return n, err
}
