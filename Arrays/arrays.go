package main

import (
	"fmt"
)

func main() {
	//Arrays
	integers := [3]int{0, 4, 6}

	//Slices
	b := integers[0:2]

	//Make to create new slices
	mySlice := make([]int, 5, 10)

	fmt.Println(integers)
	fmt.Println(b)
	fmt.Println(len(mySlice))

	//Length func for slices
	fmt.Println(len(mySlice))
	//Capacity func for slices
	fmt.Println(cap(mySlice))
}
