package main

import "fmt"

func main() {
	firstInts, secondlnts := splitAnySlice([]int{0, 1, 2, 3})
	fmt.Println(firstInts, secondlnts)
}

// Without Generics we will have to make separate functions for every
// data type
func splitIntStice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStringStice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

//and so on...

//-----------------------------------------------------------//

// With Generics:
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// Constraints:
type stringer interface {
	String() string
}

func Concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

// Interface type lists:
// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32
	//and so on...
}
