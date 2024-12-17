package main

import "fmt"

func main() {
	// Define a slice of prime numbers.
	primes := []int{2, 3, 5}

	// Attempt to access an element at index 3 in the slice.
	// If the index is out of bounds, the program will panic.
	// However, we have deferred a function to recover from the panic.
	// The program will recover and return the default value of the return type (0 in this case).
	x := accessElement(primes, 3)

	// Print the result.
	fmt.Println(x)
}

// accessElement attempts to access an element at the given index in a slice.
// If the index is out of bounds, the program panics.
// We use a deferred function to recover from the panic.
func accessElement(a []int, index int) int {
	// Defer a function to recover from a panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recovered but nothing to do!")
		}
	}()

	// Attempt to access the element at the given index.
	// If the index is out of bounds, this line will panic.
	return a[index]
}
