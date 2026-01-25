// structs.go

package main

import (
	"fmt"
)

// RunStructExample - illustration of structs

type Doctor struct{
	number int
	actorName string
	companions []string
}

func RunStructExample(){
	aDoctor := Doctor{
		number : 3,
		actorName : "Jon Pertwee",
		companions : []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// removing the keys can also declare value
	// but better avoid using this for maintenance issue
	bDoctor := Doctor{
		1, "Joe Doe", []string{"Doe Joe"},
	}

	fmt.Println(bDoctor)

	// copying struct will duplicating the values into another memory space
	cDoctor := bDoctor
	cDoctor.actorName = "Tom Baker"

	fmt.Println(bDoctor)
	fmt.Println(cDoctor)

	// to reference to the same memory space, use pointer
	dDoctor := &bDoctor
	dDoctor.number = 2

	fmt.Println(bDoctor)
	fmt.Println(dDoctor)
}
