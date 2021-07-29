package main

import (
	"fmt"
	"openclosed/math/calculator"
)

//Open/Closed Principle
//Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.

func main() {
	var a, b int
	a = 100
	b = 50
	calculator := calculator.Calculator{}

	result := calculator.Calculate(a, b)

	fmt.Printf("Result = %d", result)
	//If we want to add a new Behavior (the Subtract for example) we will have to change the calculator struct.
	//that makes the calculator struct not closed for modification
}
