package main

import "fmt"

func main() {
	//Creating a channel with a buffer
	ch := make(chan int, 100)

	//Sending on a buffered channel only blocks when the buffer is full.
	//Receiving blocks only when the buffer is empty.

	fmt.Println(ch)
}
