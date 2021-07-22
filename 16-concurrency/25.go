package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		c <- `foo`
	}()

	go func() {
		defer wg.Done()
		c <- `bar`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 4)
		println(`Message 1: ` + <-c)
		println(`Message 2: ` + <-c)
	}()

	wg.Wait()
}
