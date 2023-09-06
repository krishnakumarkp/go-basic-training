package main

//Context With Deadline

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context, timeToRun time.Duration) {
	select {
	case <-time.After(timeToRun):
		fmt.Println("completed before context deadline passed")
	case <-ctx.Done():
		fmt.Println("bailed because context deadline passed")
	}
}

const duration = 5 * time.Second

func main() {
	ctx := context.Background()

	// this will bail because the function runs longer than the context's deadline allows
	ctx1, _ := context.WithDeadline(ctx, time.Now().Add(duration))
	go longRunningTask(ctx1, 10*time.Second)

	// this will complete because the function completes before the context's deadline arrives
	//ctx2, _ := context.WithDeadline(ctx, time.Now().Add(duration))
	go longRunningTask(ctx1, 3*time.Second)

	time.Sleep(7 * time.Second)
	fmt.Println("Main done")
}
