package main

import (
	"fmt"
)

func main() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(9))
	fmt.Println(doubleFunc(69))
}

func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

// Currying
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}
