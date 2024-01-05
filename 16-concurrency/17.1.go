package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string
	c = make(chan string)

	//c := make(chan string)

	go sayHi(c)

	result := <-c

	fmt.Println(result)
}

func sayHi(ch chan string) {
	time.Sleep(8 * time.Second)
	ch <- "Hi"
}
