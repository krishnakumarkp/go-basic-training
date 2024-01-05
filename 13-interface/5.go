package main

import "fmt"

// Multiple interfaces
// A type can implement multiple interfaces.

type Geometry interface {
	Area() float64
}

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}

	PrintArea(c)

}

//func PrintArea(s Geometry) {

//fmt.Printf("Area of shape is %f \n", s.Area())

func PrintArea(s Object) {

	fmt.Printf("Volume of object is %f \n", s.Volume())
}
