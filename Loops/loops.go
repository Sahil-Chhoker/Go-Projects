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

	//Infinte Loop
	j := 1
	for {
		fmt.Print(j)
		j++
	}

	//The fizzbuzz game
	fizzbuzz()
}

//We can ommit for loops in GO
// for Initial ; ; After{

// }
//The above code is totally valid

//Therfore, the while loop in GO is:
// for CONDITION{

// }

func fizzbuzz() {
	for a := 1; a <= 100; a++ {
		if (a%3) == 0 && (a%5) == 0 {
			fmt.Println("fizzbuzz")
		} else if a%3 == 0 {
			fmt.Println("fizz")
		} else if a%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(a)
		}
	}
}
