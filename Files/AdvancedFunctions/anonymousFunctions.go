package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	//Here we define an anonymous function that doubles an int
	// and pass it to doMath

	altNumsDoubled := doMath(func(x int) int {
		return x + x
	}, nums)

	fmt.Println(altNumsDoubled)
	// prints:
	// [2 4 6 8 10]
}

// do at accepts a unction t at converts one Int Into anot er
// and a slice of ints. It returns a slice of ints that have be.
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}

	return results
}
