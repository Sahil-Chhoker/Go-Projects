package main

import (
	"fmt"
)

func main() {
	//Create an instance of the struct
	myCar := car{}
	myCar.FrontWheel.Radius = 5

	fmt.Println(myCar.FrontWheel.Radius)
}

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

//OR

//Anonymous Structs

type another_car struct {
	Make   string
	Model  string
	Height int
	Width  int

	Whell struct {
		Radius   int
		Material string
	}
}
