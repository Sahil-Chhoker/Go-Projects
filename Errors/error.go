package main

import (
	"fmt"
	"strconv"
)

func main() {
	// normalErrorHandling()
	msg, _ := sendSMS("Hello Dummy", "Sahil")
	fmt.Printf("Message : %v\n", msg)
}

func normalErrorHandling() {
	i, err := strconv.Atoi("42b")

	if err != nil {
		fmt.Println("Couldn't Convert : ", err)
		return
	}

	fmt.Println("The result is = ", i)
}

// Creating our own Errors
// The Error Interface
type error interface {
	Error() string
}

type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e)
}

func sendSMS(msg, userName string) (string, error) {
	if !canSendToUser(userName) {
		return "", userError{name: userName}
	}

	return msg, nil
}

func canSendToUser(userName string) bool {
	return true
}
