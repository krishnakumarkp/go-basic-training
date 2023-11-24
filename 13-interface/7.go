package main

//a slice of structs cannot be directly passed to a function that accepts a slice of an interface
//https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces/12754757#12754757

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func CalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func CalculateArea(shape Shape) float64 {
	totalArea := shape.Area()

	return totalArea
}

func main() {
	circles := []Circle{
		{Radius: 1.0},
		{Radius: 2.0},
	}

	fmt.Println(CalculateArea(circles[0]))

	// In Go, a slice of structs cannot be directly passed to a function that accepts a slice of an interface the struct implements because
	//Go is statically typed, and types play a crucial role in function signatures.

	//CalculateTotalArea(circles)

	var shapes []Shape
	for _, circle := range circles {
		shapes = append(shapes, circle)
	}

	fmt.Println(CalculateTotalArea(shapes))
}
