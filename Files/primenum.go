package main

import (
	"fmt"
)

func printPrime(max int) {
	for n := 1; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; (i * i) < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if !isPrime {
			continue
		}

		fmt.Println(n)
	}
}

func main() {
	fmt.Println("The prime numbers from 1 to 100 are :")

	printPrime(100)
}
