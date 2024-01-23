package main

import (
	"fmt"
)

//They are dictionaries in python, a key is paired with a value

func main() {
	//To make a map
	ages := make(map[string]int)

	//Adding stuff to a made up map
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // Overwrites 24

	fmt.Println(ages)

	//Making map with predefined values
	age := map[string]int{
		"John": 37,
		"Mary": 21,
	}

	fmt.Println(age)
}
