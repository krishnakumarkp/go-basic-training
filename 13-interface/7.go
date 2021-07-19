package main

import "fmt"

// Embedding interfaces

// In Go, an interface cannot implement other interfaces or extend them,
// but we can create a new interface by merging two or more interfaces.

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Material interface {
	Shape
	Object
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
	var s Shape
	var o Object
	var m Material

	s = Cube{3}
	o = Cube{3}
	m = Cube{3}

	fmt.Printf("Area of shape is %f \n", s.Area())
	fmt.Printf("Volume of object is %f", o.Volume())
	fmt.Printf("Area and volume of Material is %f and %f \n", m.Area(), m.Volume())
}
