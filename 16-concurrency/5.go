package main

import (
	"fmt"
	"time"
)

func main() {
	// Need to give an indication to user that program is running
	// To run a function as a goroutine, call that function prefixed with the go statement.
	go spinner(1 * time.Second)
	const n = 45
	const o = 46
	const p = 42
	fibN := fib(n)
	fibO := fib(o)
	fibP := fib(p)

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("Fibonacci(%d) = %d\n", o, fibO)
	fmt.Printf("Fibonacci(%d) = %d\n", p, fibP)

}

//Function to Find the Nth Fibonacci Number using Recursion
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, s := range `-\|/` {
			fmt.Printf("\r%c", s)
			time.Sleep(delay)
		}
	}
}
