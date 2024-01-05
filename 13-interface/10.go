// In Go, an empty interface is an interface that has zero methods.
// It serves as a type that can represent values of any type because any value in Go implements at least zero methods.
// The empty interface is denoted by the keyword interface{}.
package main

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

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// PrintObject takes an empty interface as a parameter.
func PrintObject(object interface{}) {
	fmt.Printf("Object: %+v \n", object)
}

func main() {

	dog := Dog{"Tommy"}
	cat := Cat{"Mitten"}

	circle := Circle{Radius: 1.0}

	// Using an empty interface to pass values of different types.

	PrintObject(dog)
	PrintObject(cat)
	PrintObject(circle)
}

// The empty interface should be avoided as much as possible. If you ever use it,
// you are saying that variable can be anything, even types you don’t expect.
// Using it means you lose all the advantages of using a strongly typed language because the compiler can’t tell you that you made a mistake.
// You won’t find the mistake until runtime.
