package main

import (
	"fmt"
)

func main() {
	fmt.Println(aggregate(2, 3, 4, add))

	fmt.Println(aggregate(2, 3, 4, mul))
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func aggregate(a, b, c int, arthematic func(int, int) int) int {
	return arthematic(arthematic(a, b), c)
}
