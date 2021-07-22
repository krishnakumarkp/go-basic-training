package main

import (
	"fmt"
	"sync"
	"time"
)

//we need to print the result as sonn as any calculation is done.Dont wait for all the calculation
func main() {
	go spinner(1 * time.Second)
	const n = 45
	const o = 46
	const p = 42

	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go func(c chan int) {
		fibN := fib(n)
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func(c chan int) {
		fibO := fib(o)
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func(c chan int) {
		fibP := fib(p)
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go printResult()

	//How will fib() return the value to printResult()

	waitgroup.Wait()

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

func printResult() {
	for {
		fmt.Printf("\rFibonacci = %d\n", result)
	}
}

// We are not able to identify the result is for which number
// All the results are not getting printed
