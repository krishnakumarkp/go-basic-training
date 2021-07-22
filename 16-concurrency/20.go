package main

import (
	"fmt"
	"sync"
	"time"
)

type result struct {
	number   int
	fibValue int
}

//we need to print the result as sonn as any calculation is done.Dont wait for all the calculation
func main() {
	go spinner(1 * time.Second)
	const n = 45
	const o = 46
	const p = 42
	var waitgroup sync.WaitGroup
	resultChannel := make(chan result)
	defer close(resultChannel)

	waitgroup.Add(1)
	go func(c chan result) {
		fibN := fib(n)
		r := result{
			number:   n,
			fibValue: fibN,
		}
		c <- r
		waitgroup.Done()
	}(resultChannel)

	waitgroup.Add(1)
	go func(c chan result) {
		fibO := fib(o)
		r := result{
			number:   o,
			fibValue: fibO,
		}
		c <- r
		waitgroup.Done()
	}(resultChannel)

	waitgroup.Add(1)
	go func(c chan result) {
		fibP := fib(p)
		r := result{
			number:   p,
			fibValue: fibP,
		}
		c <- r
		waitgroup.Done()
	}(resultChannel)

	waitgroup.Add(1)
	go printResult1(resultChannel, &waitgroup)
	waitgroup.Add(1)
	go printResult1(resultChannel, &waitgroup)
	waitgroup.Add(1)
	go printResult1(resultChannel, &waitgroup)

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

func printResult1(c chan result, w *sync.WaitGroup) {
	fibResult := <-c
	fmt.Printf("\rFibonacci(%d) = %d\n", fibResult.number, fibResult.fibValue)
	w.Done()
}
