package main

import "fmt"

type Animal struct{
    Species string
}

func (a Animal) Sound() string{
    return "Some sound"
}

type Dog struct{
    Animal
    Breed string
}

func (d Dog) Sound() string{
    return "Woof"
}

func main(){
    dog := Dog{Animal{"Canine"}, "Labrador"}

    fmt.Println("Species : ", dog.Species)
    fmt.Println("Breed : ", dog.Breed)
    fmt.Println("Sound : ", dog.Sound())

    a := Animal{Species : "Canine"}
    sound := a.Sound()
    fmt.Println("Sound : ", sound)
}
