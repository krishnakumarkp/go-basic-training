package main

import (
	"fmt"
	"math"
)

// Type assertions
// We can find out the underlying dynamic value of an interface using the syntax i.(Type)
// where i is a variable of type interface and Type is a type that implements the interface.

// The type assertion x.(T) asserts that the concrete value stored in x is of type T,
// and that x is not nil.

// If T is not an interface, it asserts that the dynamic type of x is identical to T.
// If T is an interface, it asserts that the dynamic type of x implements T.

type Shape interface {
	Area() float64
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

type Circle struct {
	radious float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radious * c.radious
}

func main() {
	var s Shape = Cube{3}
	var c Shape = Circle{5}

	measure(s)
	measure(c)

}

func measure(i Shape) {

	c, ok := i.(Circle)
	if ok {
		fmt.Println("area of c of type Circle is", c.Area())
	}

	cb, ok := i.(Cube)

	if ok {
		fmt.Println("area of cb of type Cube is", cb.Area())
		fmt.Println("volume of cb of type Cube is", cb.Volume())
	}

}
