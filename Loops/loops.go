package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	//Example of while loop
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("Still growing!! Current Height : ", plantHeight)
		plantHeight++
	}

	fmt.Println("Plant has grown to ", plantHeight, " inches")
}

//We can ommit for loops in GO
// for Initial ; ; After{

// }
//The above code is totally valid

//Therfore, the while loop in GO is:
// for CONDITION{

// }
