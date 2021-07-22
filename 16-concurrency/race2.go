package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

func main() {
	//declare a number of goroutines
	total := 10
	var waitgroup sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < total; i++ {
		waitgroup.Add(1)
		//launch a goroutine with anonymous function
		go func() {
			lock.Lock()
			value := counter 
			time.Sleep(1 * time.Second)
			value++
			counter = value
			lock.Unlock()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	fmt.Println("Counter: ", counter)
}
