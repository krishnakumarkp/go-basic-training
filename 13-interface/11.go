package main

import (
	"fmt"
	"math"
)

// Circle is a defined type for circles.
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
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

// CalculateArea calculates and prints the area of a given shape.
func CalculateAreaAndVolume(shape interface{}) {
	switch v := shape.(type) {
	case Circle:
		fmt.Printf("Area : %.2f\n", v.Area())
	case Cube:
		fmt.Printf("Area: %.2f\n", v.Area())
		fmt.Printf("Volume: %.2f\n", v.Volume())

	default:
		fmt.Println("Unknown shape type!")
	}
}

func main() {
	// Create instances of different shapes.
	//circle := Circle{Radius: 5.0}
	cube := Cube{3}

	// Calculate and print the areas using the CalculateArea function.
	CalculateAreaAndVolume(cube)

	//what if we add more shapes like rectangle trigange etc?
	// we will have to add more case conditions to the function
}
