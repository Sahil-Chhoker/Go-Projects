package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	// normalErrorHandling()
	msg, _ := sendSMS("Hello Dummy", "Sahil")
	_msg, err := _sendSMS("Hello Dummy", "Sahil")
	fmt.Printf("Message : %v\n", msg)
	fmt.Printf("Message : %v\nError : ", _msg, err)
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

//Alternate Method for custom errors(Using "errors" Package)

func _sendSMS(msg, userName string) (string, error) {
	var err error = errors.New("Error is here")

	if !canSendToUser(userName) {
		return "", err
	}

	return msg, nil
}

func canSendToUser(userName string) bool {
	return false
}
