package main

import (
	"fmt"
)

func main() {
	//Arrays
	integers := [3]int{0, 4, 6}

	//Slices
	b := integers[0:2]

	fmt.Println(integers)
	fmt.Println(b)
}
