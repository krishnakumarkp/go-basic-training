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

//we need to print the result as soon as any calculation is done.Dont wait for all the calculation
func main() {
	numbers := []int{45, 46, 42}
	var waitgroup sync.WaitGroup
	resultChannel := make(chan result)
	defer close(resultChannel)

	go spinner(1 * time.Second)
	for _, num := range numbers {
		waitgroup.Add(1)
		go calculateFib(num, resultChannel, &waitgroup)
		waitgroup.Add(1)
		go printResult(resultChannel, &waitgroup)
	}
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

func calculateFib(x int, c chan result, w *sync.WaitGroup) {
	fibN := fib(x)
	r := result{
		number:   x,
		fibValue: fibN,
	}
	c <- r
	w.Done()
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func printResult(c chan result, w *sync.WaitGroup) {
	fibResult := <-c
	fmt.Printf("\rFibonacci(%d) = %d\n", fibResult.number, fibResult.fibValue)
	w.Done()
}
