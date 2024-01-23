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
