package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(1 * time.Second)
	const n = 45
	const o = 46
	const p = 42
	var fibN, fibO, fibP int

	go func() {
		fibN = fib(n)
	}()

	go func() {
		fibO = fib(o)
	}()

	go func() {
		fibP = fib(p)
	}()

	//Make main wait for go routines finish
	time.Sleep(60 * time.Second)
	//What is the problem here?
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("Fibonacci(%d) = %d\n", o, fibO)
	fmt.Printf("Fibonacci(%d) = %d\n", p, fibP)
}

func spinner(delay time.Duration) {
	for {
		for _, s := range `-\|/` {
			fmt.Printf("\r%c", s)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
