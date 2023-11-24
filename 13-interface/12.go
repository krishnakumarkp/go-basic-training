package main

import (
	"fmt"
	"math"
)

type AreaCalculator interface {
	Area() float64
}

type VolumeCalculator interface {
	Volume() float64
}

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

	if areaCalc, ok := shape.(AreaCalculator); ok {
		fmt.Printf("Area: %.2f\n", areaCalc.Area())
	}
	if volumeCalc, ok := shape.(VolumeCalculator); ok {
		fmt.Printf("Volume: %.2f\n", volumeCalc.Volume())
	}
}

func main() {
	// Create instances of different shapes.
	circle := Circle{Radius: 5.0}

	//cube := Cube{3}

	// Calculate and print the areas using the CalculateArea function.
	//CalculateAreaAndVolume(circle)
	CalculateAreaAndVolume(circle)
}
