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

	//Append Functions for Slices:

	// slice.append(slice, oneThing) or
	// slice.append(slice, firstThing, secondTheme) or
	// slice.append(slice, anotherSlice...)

	//Slice of Slices:
	a := createMatrix(10, 10)
	fmt.Println(a)

	//Note : DON'T DO THIS::
	// oneSlice = append(otherSlice, element)

	//Easy Way loop over Slices:
	// for INDEX, ELEMENT := range SLICE {

	// }

	fruits := []string{"Apple", "Banana", "Grapes"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

//SLice of Slices

func createMatrix(rows, columns int) [][]int {
	matrix := make([][]int, 0)

	for i := 0; i < (rows); i++ {
		rows := make([]int, 0)
		for j := 0; j < (columns); j++ {
			rows = append(rows, i*j)
		}

		matrix = append(matrix, rows)
	}

	return matrix
}
