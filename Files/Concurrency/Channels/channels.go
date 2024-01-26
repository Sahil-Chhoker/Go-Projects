package main

import "fmt"

//Channels are used to store values mainly from seperate goroutines,
//Since they can't have a return value

func main() {
	//Making a channel
	ch := make(chan int)

	//Sending data to a channel
	ch <- 69

	//Receiving value from a channel
	v := <-ch

	//NOTE :
	//The sending and receiving values from channels are blocking functions

	fmt.Print(v)
}

//BLOCKING AND DEADLOCKS
//A deadlock is when a group of goroutines are all blocking so none of them
//can continue. This is a common bug that you need to watch out for in
//concurrent programming.
