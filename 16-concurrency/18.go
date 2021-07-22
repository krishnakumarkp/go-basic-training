package main

import (
	"fmt"
	"sync"
	"time"
)

//we need to print the result as soon as any calculation is done.Dont wait for all the calculation
func main() {
	go spinner(1 * time.Second)
	const n = 45
	const o = 46
	const p = 42

	var waitgroup sync.WaitGroup

	resultChannel := make(chan int)
	defer close(resultChannel)

	waitgroup.Add(1)
	go func(c chan int) {
		fibN := fib(n)
		c <- fibN
		waitgroup.Done()
	}(resultChannel)

	waitgroup.Add(1)
	go func(c chan int) {
		fibO := fib(o)
		c <- fibO
		waitgroup.Done()
	}(resultChannel)

	waitgroup.Add(1)
	go func(c chan int) {
		fibP := fib(p)
		c <- fibP
		waitgroup.Done()
	}(resultChannel)

	//waitgroup.Add(1)
	go printResult(resultChannel)

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

func printResult(c chan int) {
	for {
		result := <-c
		fmt.Printf("\rFibonacci = %d\n", result)
	}
}

// We are not able to identify the result is for which number
// All the results are not getting printed
