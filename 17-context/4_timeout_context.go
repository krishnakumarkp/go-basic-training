package main

//Context With Timeout

//https://betterprogramming.pub/how-and-when-to-use-context-in-go-b365ddf42ae2
//https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context, timeToRun time.Duration) {
	select {
	case <-time.After(timeToRun):
		fmt.Println("completed before context timed out")
	case <-ctx.Done():
		fmt.Println("bailed because context timed out")
	}
}

const timeout = 5 * time.Second

func main() {
	ctx := context.Background()

	// this will bail because the function takes longer than the context allows
	ctx1, _ := context.WithTimeout(ctx, timeout)
	//Under the hood, they are calling: WithDeadline(parent, time.Now().Add(timeout))
	go longRunningTask(ctx1, 10*time.Second)

	// this will complete because the function completes before the context times out
	//ctx2, _ := context.WithTimeout(ctx, timeout)
	go longRunningTask(ctx1, 1*time.Second)

	time.Sleep(7 * time.Second)
	fmt.Println("Main done")
}
