package main

import "fmt"

func main() {
	//A sent to a nil channel blocks forever:

	var c chan string        //c is nil
	c <- "let's get started" // Blocks

	//--------------------------------------------------------//

	//A receive from a nil channel blocks forever:

	var ch chan string //ch is nill
	fmt.Println(<-ch)  //Blocks

	//--------------------------------------------------------//

	//A sent to a closed channel panics:
	var cha = make(chan int, 100)
	close(cha)

	cha <- 1 //Panics

	//--------------------------------------------------------//

	//A receive from a closed channel returns the zero value immediately

	var channel = make(chan int, 100)
	close(channel)

	fmt.Println(<-channel) // 0
}
