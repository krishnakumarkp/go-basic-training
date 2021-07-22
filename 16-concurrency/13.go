package main

import (
	"fmt"
	"sync"
	"time"
)

//Waiting for Goroutines to Finish Execution
func main() {
	go spinner(1 * time.Second)
	const n = 45
	const o = 46
	const p = 42
	var fibN, fibO, fibP int

	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go func() {
		fibN = fib(n)
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func() {
		fibO = fib(o)
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func() {
		fibP = fib(p)
		waitgroup.Done()
	}()

	waitgroup.Wait()

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
