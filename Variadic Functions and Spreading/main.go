package main

import (
	"fmt"
)

func main() {
	//Calling Variadic function
	fmt.Println(sum(1, 2, 3, 4))
}

// Variadic Function:
func sum(nums ...int) int {
	var num int = 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}
