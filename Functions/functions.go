package main

import "fmt"

func main() {
	fmt.Println(add1(5, 6))
	fmt.Println(add2(5, 6))

	a, b := addAndSub(5, 6)
	fmt.Println(a, b)

	//To ignore one value out of many return values
	c, _ := addAndSub(5, 6)
	fmt.Println(c)
}

func add1(x int, y int) int {
	return x + y
}

// Alternate
func add2(x, y int) int {
	return x + y
}

// Functions in go can return multiple values
func addAndSub(x, y int) (int, int) {
	return x + y, x - y
}
