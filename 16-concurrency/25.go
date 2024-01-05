package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with a capacity of 5
	bufferedChannel := make(chan int, 5)

	// Start a goroutine to send values to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			// Send values to the buffered channel
			fmt.Printf("Sent: %d\n", i)
			bufferedChannel <- i

		}

		// Close the channel after sending all values
		close(bufferedChannel)
	}()

	// Allow some time for the goroutine to send values
	time.Sleep(time.Second)

	// Receive values from the buffered channel
	for value := range bufferedChannel {
		fmt.Printf("Received: %d\n", value)
		time.Sleep(time.Millisecond * 500) // Simulate some processing time
	}
}
