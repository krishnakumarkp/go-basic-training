package main

import "fmt"

func main() {
	// Print a message to indicate the start of the main function.
	fmt.Println("main() started")

	// Defer the execution of defferingFromMain to the end of main.
	defer defferingFromMain()

	// Call the panickingFunction.
	panickingFunction()

	// This line will not be executed if a panic occurs in panickingFunction.
	fmt.Println("main() done")
}

func panickingFunction() {
	// Print a message to indicate the start of the panickingFunction.
	fmt.Println("panickingFunction() started")

	// Defer the execution of defferingFunction1 to the end of panickingFunction.
	defer defferingFunction1()

	// Trigger a panic with a custom message.
	panic("panic from panickingFunction()")

	// This line will not be executed because a panic has occurred.
	fmt.Println("panickingFunction() done")
}

func defferingFunction1() {
	// Print a message to indicate the execution of defferingFunction1.
	fmt.Println("defferingFunction1() executed")
}

func defferingFunction2() {
	// Print a message to indicate the execution of defferingFunction2.
	fmt.Println("defferingFunction2() executed")
}

func defferingFromMain() {
	// Print a message to indicate the execution of defferingFromMain.
	fmt.Println("defferingFromMain() executed")
}
