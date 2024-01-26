package main

func main() {

	go doSometing()

	//In the example above, doSomething() will be executed concurrently
	//with the rest of the code in the function. The go keyword is used to spawn
	//a new gonutLne.

	//NOTE : We can't have a return value from this function
}

func doSometing() {
	return
}
