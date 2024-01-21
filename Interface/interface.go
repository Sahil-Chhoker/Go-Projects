package main

import (
	"fmt"
	"math"
)

func main() {

	r := rect{
		width:  5,
		height: 10,
	}

	c := circle{
		radius: 5,
	}

	checkType(shape(r))
	checkType(shape(c))

	_checkType(shape(r))
	_checkType(shape(c))

	fmt.Printf("Area of Rectangle : %f\nPerimeter of Rectangle : %f\n", r.area(), r.perimeter())
	fmt.Printf("Area of Circle : %f\nPerimeter of circle : %f\n", c.area(), c.perimeter())
}

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func checkType(s shape) {
	cir, ok := s.(circle)

	if ok {

		fmt.Println("It is a Circle with area = ", cir.area())
	}

	rec, ok := s.(rect)

	if ok {
		fmt.Println("It is a Rectangle with area = ", rec.area())
	}
}

//Alternate to checkType Method

func _checkType(s shape) {
	switch v := s.(type) {
	case circle:
		fmt.Printf("It is a %T\n", v)
	case rect:
		fmt.Printf("It is a %T\n", v)
	}

}
