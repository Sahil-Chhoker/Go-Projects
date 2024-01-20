package main

import (
	"fmt"
)

func main() {
	//Create an instance of the struct
	myCar := car{}
	myCar.FrontWheel.Radius = 5

	fmt.Println(myCar.FrontWheel.Radius)

	r := rect{
		width:  5,
		height: 10,
	}

	fmt.Println(r.area())
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

//Embeded structs

type _car struct {
	make  string
	model string
}

type truck struct {
	car
	bedsize int
}

//Struct Methods

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}
