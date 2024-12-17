package main

import "fmt"

func main() {
	// Call the fun1 function and print its return value.
	fmt.Println("Function Return Value:", fun1())
}

func fun1() int {
	// Initialize a variable i with a value of 0.
	i := 0

	// Defer three functions, which will be executed at the end of the fun1 function.
	// The deferred functions capture and modify the value of the variable i.
	// These deferred functions are executed in reverse order when the fun1 function returns.

	// Deferred function 3:
	// - Print the value of i (which is 0 at this point) and increment it.
	defer func() {
		fmt.Println("3rd", i)
		i++
	}()

	// Deferred function 2:
	// - Print the value of i (which is 0 at this point) and increment it.
	defer func() {
		fmt.Println("2nd", i)
		i++
	}()

	// Deferred function 1:
	// - Print the value of i (which is 0 at this point) and increment it.
	defer func() {
		fmt.Println("1st", i)
		i++
	}()

	// Return the current value of i (which is 0).
	// The deferred functions will not affect the returned value of fun1.
	return i
}
