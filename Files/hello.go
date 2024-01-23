package main

import (
	"fmt"
)

func main() {
	var a int = 32
	var float_ float64 = 30.23
	var str string = "Hi"
	var trueorfalse bool = true

	const helloThere = "Hello, There"
	const basicThere = "jk"

	congrats := "Congratulations"

	averageOpenRate, displayTime := 89.00, "23 seconds"

	fmt.Printf(helloThere)

	fmt.Printf("Type %T\n", averageOpenRate)

	fmt.Println(a, float_, str, trueorfalse, congrats, averageOpenRate, displayTime)
}
