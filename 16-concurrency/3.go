package main

import (
	"fmt"
)

func main() {
	const n = 45
	fibN := fib(n)
	fmt.Printf("\r Fibonacci(%d) = %d\n", n, fibN)
}

//Function to Find the Nth Fibonacci Number using Recursion
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
