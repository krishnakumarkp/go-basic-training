package main

//organizing code with packages
//create folder shapes
//create file shapes.go in shapes folder

import (
	"fmt"
	"go-training/methods/shapes"
)

func main() {
	r := shapes.NewRectangle(10, 15)
	fmt.Printf("Area of rectangle %d\n", r.Area())

	c := shapes.NewCircle(12)
	fmt.Printf("Area of circle %f", c.Area())
}
