package main

import (
	"fmt"
	"openclosed/math/calculator"
	"openclosed/math/operations"
)

//Open/Closed Principle
//Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.

func main() {
	var a, b int
	a = 100
	b = 50

	calculator := calculator.Calculator{}

	calculator.Operaton = operations.Add{}
	sum := calculator.Calculate(a, b)

	calculator.Operaton = operations.Subtract{}
	difference := calculator.Calculate(a, b)

	fmt.Printf("Sum = %d \n", sum)
	fmt.Printf("Difference = %d", difference)

	//If we want to add a new Behavior (the Multiply for example) we dont have to change the calculator struct.
	// we just need to create a new Multiply struct and embed it in calculator
	//that makes the calculator struct open for extension, but closed for modification

	//The Open / Closed Principle encourages you to compose simple types into more complex ones using embedding
}
