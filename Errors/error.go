package main

import (
	"fmt"
	"strconv"
)

func main() {
	normalErrorHandling()
}

// The Error Interface
type error interface {
	Error() string
}

func normalErrorHandling() {
	i, err := strconv.Atoi("42b")

	if err != nil {
		fmt.Println("Couldn't Convert : ", err)
		return
	}

	fmt.Println("The result is = ", i)
}
