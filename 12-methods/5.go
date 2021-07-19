package main

import (
	"fmt"
	"math"
)

//Methods with the same name can be defined on different types

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	myRectangle := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", myRectangle.Area())

	myCircle := Circle{
		radius: 12,
	}

	fmt.Printf("Area of circle %f", myCircle.Area())

}

// does this look like object oriented programming ?
