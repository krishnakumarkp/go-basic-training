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
	var s Shape = c
	var o Object = c
	var g Geometry = c
	fmt.Println("are of s of interface type Shape is", s.Area())
	fmt.Println("volume of o of interface type Object is", o.Volume())
	fmt.Println("area of g of interface type Geometry is", g.Area())
}
