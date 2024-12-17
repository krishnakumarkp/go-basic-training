package main

import "fmt"

func main() {
	// Print a message to indicate the start of the main function.
	fmt.Println("main() started")

	// Call the panickingFunction.
	panickingFunction()

	// This line will not be executed if a panic occurs in panickingFunction.
	fmt.Println("main() done")
}

func panickingFunction() {
	// Defer the execution of the defferingFunction to the end of the function.
	defer defferingFunction()

	// Print a message to indicate the start of the panickingFunction.
	fmt.Println("panickingFunction() started")

	// Trigger a panic with a custom message.
	panic("panic from panickingFunction()")

	// This line will not be executed because a panic has occurred.
	fmt.Println("panickingFunction() done")
}

func defferingFunction() {
	// Print a message to indicate the execution of the defferingFunction.
	fmt.Println("defferingFunction() executed")
}
