package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32

func main() {
	//declare a number of goroutines
	total := 10
	var waitgroup sync.WaitGroup

	for i := 0; i < total; i++ {
		waitgroup.Add(1)
		//launch a goroutine with anonymous function
		go func() {
			atomic.AddInt32(&counter, 1)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	fmt.Println("Counter: ", counter)
}
