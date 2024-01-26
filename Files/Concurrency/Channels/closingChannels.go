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

	//NOTE :
	//Don't send on a closed channel

	//SELECT statement:
	select {
	case i, ok := <-ch:
		fmt.Println(i, ok)
	case s, ok := <-ch:
		fmt.Println(s, ok)
	default:
		//receiving from ch would block
		//so do somethong else
	}
}
