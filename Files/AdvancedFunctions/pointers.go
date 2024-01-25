package main

import (
	"fmt"
)

type car struct {
	color string
}

func (c *car) setCotor(color string) {
	c.color = color
}

func main() {
	var p *int // Null Pointer

	myString := "Hello"
	myStringPtr := &myString

	fmt.Println("Empty Pointer : ", p)
	fmt.Println("Actual Pointer : ", myStringPtr)

	fmt.Println(*myStringPtr) // Syntax to get the value stored at the address of the pointer
	*myStringPtr = "World"    // Reassigning the value of the pointer

	fmt.Println(*myStringPtr)

	//NOTE : NIL POINTERS
	//Always check if the pointer is nil or not before dereferencing it

	//POINTER RECEIVERS
	c := car{
		color: "white",
	}

	c.setCotor("Blue")
	fmt.Println(c.color)
	// prints "blue"
}
