package main

import "fmt"

func main() {
	height := 5
	if height > 4 {
		fmt.Println("You are tall enough")
	} else {
		fmt.Println("You are not tall enough")
	}

	//Alternate
	if length := 5; length > 4 {
		fmt.Printf("You may pass")
	} else {
		fmt.Println("Stop")
	}
}
