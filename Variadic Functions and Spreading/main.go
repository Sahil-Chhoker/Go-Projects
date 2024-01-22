package main

import (
	"fmt"
)

func main() {
	//Calling Variadic function
	fmt.Println(sum(1, 2, 3, 4))

	//Calling by Spread Operators:
	names := []string{"bob", "sue", "damn"}
	printStrings(names...)

	//Append Functions for Slices:

	// slice.append(slice, oneThing) or
	// slice.append(slice, firstThing, secondTheme) or
	// slice.append(slice, anotherSlice...)

	//Slice of Slices:
	a := createMatrix(10, 10)
	fmt.Println(a)
}

// Variadic Function:
func sum(nums ...int) int { // Where ...int is a slice
	var num int = 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}

//Spread Operators:

// Variadic func
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
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
