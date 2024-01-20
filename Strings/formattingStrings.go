package main

import "fmt"

func main() {
	//Default
	fmt.Printf("This file %v formatting strings\n", "is related to")

	//String
	fmt.Printf("This file %s formatting strings\n", "is related to")

	//Integer
	fmt.Printf("This is %dnd file\n", 2)

	//Float(Used frequently)
	fmt.Printf("I am %.0f years old", 17.98)
}
