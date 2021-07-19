package shapes

import "math"

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

func NewRectangle(length int, width int) Rectangle {
	r := Rectangle{
		length: length,
		width:  width,
	}
	return r
}

func NewCircle(radius float64) Circle {
	c := Circle{
		radius: radius,
	}
	return c
}
