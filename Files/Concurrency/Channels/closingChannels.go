package main

import "fmt"

func main() {
	ch := make(chan int)

	//do some stuff with the channel

	close(ch)

	//Checking if a channel is closed
	c, ok := <-ch

	fmt.Println(c, ok)

	//ok is false if the channel is empty and closed
}
