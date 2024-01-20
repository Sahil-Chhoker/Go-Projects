package main

import "fmt"

func main() {
	fmt.Println(add1(5, 6))
	fmt.Println(add2(5, 6))
}

func add1(x int, y int) int {
	return x + y
}

// Alternate
func add2(x, y int) int {
	return x + y
}
