package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string
	c = make(chan string)

	//c := make(chan string)

	go func(ch chan string) {
		time.Sleep(8 * time.Second)
		ch <- "Hi"
	}(c)
	fmt.Println(<-c)
}
