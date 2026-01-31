// defer_panic_recover.go

package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func RunDeferPanicRecoverExample(){
	RunDPRE01()
	RunDPRE02()
	RunDPRE03()
	RunDPRE04()
	RunDPRE05()
	RunDPRE06()
	RunDPRE07()
	RunDPRE08()
	RunDPRE09()
}

func RunDPRE01(){
	// example01
	// normally go reads and execute progarm line by line
	// defer can make the line of progrme run later
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")
}

func RunDPRE02(){
	// example02
	// defer is running like last-in-first-out mode (stack)
	// the latest defer will be executed first
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
}

func RunDPRE03(){
	// example03
	// usage of defer - reading web document
	// having defer in the line of closing connection can ensure that all the works have done before connection close
	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%s\n", robots)
}

func RunDPRE04(){
	// example04
	// under this case, the "start" will be printed out, since the function stack has already got the value and execution before new assignment
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func RunDPRE05(){
	// example05
	// un go there is no exception, as exception cases in other languages are considered as normal in go
	// instead there is panic which shows really panic when "error cases" comes in
	// below will show "panic: runtime error: integer divide by zero" when error

	/*
	a, b := 1, 0
	ans := a / b
	fmt.Printf("%v / %v = %v \n", a, b, ans)
	*/
}

func RunDPRE06(){
	//example06
	// panic is just like consol.err, that shows error message
	/*
	fmt.Println("Start")
	panic("something wrong happened")
	fmt.Println("End")
	*/
}

func RunDPRE07(){
	// example07
	// simple local server example

	// routing to index, just pring out Hello Go!
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello Go!"))
	})*/

	// connect to server, with port 8000
	// if there will be error (e.g. the port has been occupied), use panic to print out the error 
	/*err := http.ListenAndServe(":8000", nil)
	if err != nil{
		panic(err.Error())
	}*/
}

func RunDPRE08(){
	// example08
	// the sequence of running code : normal statements, defer, panic
	/*
	fmt.Println("Start")
	defer func(){
		if err := recover(); err != nil{
			log.Println("Error: ", err)
		}
	}()
	panic("Something happened")
	fmt.Println("End")
	*/
}

func RunDPRE09(){
	// example09
	// having a panicker can limit the error recovery
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

// this can remove panicker function stack, then run the remaining functions in func stack
func panicker(){
	fmt.Println("something going to happen")
	defer func(){
		if err := recover() ; err != nil{
			log.Println("Error : ", err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("panicing done")
}
