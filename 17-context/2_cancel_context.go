package main

// context with cancel
import (
	"context"
	"fmt"
	"time"
)

func longRunningTask1(ctx context.Context) {
	fmt.Println("background long running task 1 launched")
	select {
	case <-ctx.Done():
		fmt.Println("long running task1 bailed because context cancelled")
	}
}

func longRunningTask2(ctx context.Context) {
	fmt.Println("background long running task 2 launched")
	select {
	case <-ctx.Done():
		fmt.Println("long running task2 bailed because context cancelled")
	}
}

func main() {
	// this will bail when cancelFunc is called
	ctx, cancelFunc := context.WithCancel(context.Background())

	go longRunningTask1(ctx)
	go longRunningTask2(ctx)

	time.Sleep(5 * time.Second)

	fmt.Println("background long running tasks still going")
	time.Sleep(5 * time.Second)

	fmt.Println("going to cancel background tasks")
	cancelFunc()

	time.Sleep(5 * time.Second)
	fmt.Println("some time has elapsed after cancelling")
}
