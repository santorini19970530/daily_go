// interfaces_example03.go
// composing interfaces

package main

import(
	"fmt"
	"bytes"
	"io"
)

func RunInterfacesExample03(){
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello World, this is a test"))
	wc.Close()
}

// get the address of the closer
func RunInterfacesExample04(){
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello World, this is a test"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)
}

func RunInterfacesExample05(){
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello World, this is a test"))
	wc.Close()

	r, ok := wc.(io.Reader)
	// this will succeed
	// r, ok := wc.(*BufferedWriterCloser)
	if ok{
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

// same as example 01, skip
/*
type Writer interface{
	Write([]byte) (int, error)
}
*/

type Closer interface{
	Close() error
}

type WriterCloser interface{
	Writer
	Closer
}

type BufferedWriterCloser struct{
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error){
	n, err := bwc.buffer.Write(data)

	if err != nil{
		return 0, err
	}

	w := make([]byte, 8)

	for bwc.buffer.Len() > 8{
		_, err := bwc.buffer.Read(w)

		if err != nil{
			return 0, err
		}

		_, err = fmt.Println(string(w))

		if err != nil{
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error{
	for bwc.buffer.Len() > 0{
		data := bwc.buffer.Next(8)

		_, err := fmt.Println(string(data))

		if err != nil{
			return err
		}
	}

	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser{
	return &BufferedWriterCloser{
		buffer : bytes.NewBuffer([]byte{}),
	}
}
