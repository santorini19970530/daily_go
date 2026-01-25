// maps.go

package main

import (
	"fmt"
)

// RunMapExample - illustration of map

func RunMapExample(){
	// format of declaring map
	// map [key type] value type {}
	// the key and value types needs to be consistent within the map
	statePopulations := map[string]int{
		"Califonia" : 39250017,
		"Texas" : 27862596,
		"Florida" : 20612439,
		"New York" : 19745289,
		"Pennsylvania" : 12802503,
		"Illinois" : 12801539,
		"Ohio" : 11614373,
	}

	fmt.Println(statePopulations)

	// make can also be used for declaring
	// statePopulations := make(map[string]int)

	// calling out the value based on key
	fmt.Println(statePopulations["Ohio"])

	// add a key-value pair
	statePopulations["Georgia"] = 10310371
	// the return ordering is not guarenteed
	fmt.Println(statePopulations)

	// remove a key-value pair
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	// when the pair is removed, or the key does not exist, the value shall be zero
	// so to validate wheher these is the key-value pair
	// this will get false for ok variable
	pop, ok := statePopulations["Oho"]
	fmt.Println(pop, ok)
	// or just getting the right variable
	// _, ok := statePopulations["Oho"]

	// len() size of map
	fmt.Println(len(statePopulations))

	// duplicating map with another variable
	// the same memory space will be referenced
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
