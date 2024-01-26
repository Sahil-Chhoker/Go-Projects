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

	fmt.Print(v)
}
