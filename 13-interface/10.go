// In Go, an empty interface is an interface that has zero methods.
// It serves as a type that can represent values of any type because any value in Go implements at least zero methods.
// The empty interface is denoted by the keyword interface{}.

package main

import "fmt"

// PrintValue takes an empty interface as a parameter.
func PrintValue(value interface{}) {
	fmt.Println("Value:", value)
}

func main() {
	// Using an empty interface to pass values of different types.
	PrintValue(42)
	PrintValue("Hello, Gopher!")
	PrintValue(3.14)
}

// The empty interface should be avoided as much as possible. If you ever use it,
// you are saying that variable can be anything, even types you don’t expect.
// Using it means you lose all the advantages of using a strongly typed language because the compiler can’t tell you that you made a mistake.
// You won’t find the mistake until runtime.
