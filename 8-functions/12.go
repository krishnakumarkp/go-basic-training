package main

import "fmt"

func main() {
	// Print a message to indicate the start of the main function.
	fmt.Println("main() started")

	// Call the foo function, which may panic.
	foo()

	// This line will not be executed if a panic occurs in the foo function.
	fmt.Println("main() done")
}

func foo() {
	// Print a message to indicate the start of the foo function.
	fmt.Println("foo() started")

	// Use the panic function to trigger a panic with a custom message.
	panic("panic from foo()")

	// This line will not be executed because a panic has occurred.
	fmt.Println("foo() done")
}
