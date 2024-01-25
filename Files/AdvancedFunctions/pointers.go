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

	fmt.Println(*myStringPtr) // Syntax to get the value stored at the address of the pointer
	*myStringPtr = "World"    // Reassigning the value of the pointer

	fmt.Println(*myStringPtr)
}
