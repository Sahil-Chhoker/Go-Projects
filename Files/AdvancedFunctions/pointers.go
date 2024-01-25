package main

import (
	"fmt"
)

func main() {
	var p *int // Null Pointer

	myString := "Hello"
	myStringPtr := &myString

	fmt.Println("Empty Pointer : ", p)
	fmt.Println("Actual Pointer : ", myStringPtr)
}
