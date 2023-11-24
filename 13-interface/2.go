package main

import (
	"fmt"
	"math"
)

// interface definition
type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radious float64
}

// rectange implements geometry
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// circle implements geometry
func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radious
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radious: 5}

	//possible since rectangle and circle imlements geometry
	//We have just achieved polymorphism.
	measure(r)
	measure(c)
}

func measure(g geometry) {
	fmt.Printf("geometry = %+v \n", g)
	fmt.Printf("area = %+v \n", g.area())
	fmt.Printf("perimeter = %+v \n", g.perimeter())
}
