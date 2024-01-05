package main

import (
	"fmt"
	"sync"
	"time"
)

// better way to write this will be to use a chanel
//Don’t communicate by sharing memory; share memory by communicating

func main() {
	// Declare a number of goroutines
	total := 10
	totalCount := 0
	var waitgroup sync.WaitGroup

	// Create a channel for communication
	counterCh := make(chan int, total)

	for i := 0; i < total; i++ {
		waitgroup.Add(1)
		// Launch a goroutine with an anonymous function
		go func() {
			// Simulate some processing time
			time.Sleep(1 * time.Second)

			// Send the updated counter value through the channel
			counterCh <- 1

			waitgroup.Done()
		}()
	}

	// Close the channel once all goroutines finish
	go func() {
		waitgroup.Wait()
		close(counterCh)
	}()

	// Wait for values from the channel and update the counter
	for count := range counterCh {
		totalCount += count
	}

	fmt.Println("Counter:", totalCount)
}
