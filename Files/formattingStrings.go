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
	fmt.Printf("I am %.0f years old\n", 17.98)

	const name = "Saul Goodman"
	const openRate = 30.5

	//.Sprintf is used to assign the value to some other variable
	msg := fmt.Sprintf("%v has a %.1f open rate", name, openRate)

	fmt.Println(msg)
}
