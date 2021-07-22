package main

import (
	"fmt"
	"time"
)

// Default case
// The default case in a select statement is executed when none of the other cases is ready.
// This is generally used to prevent the select statement from blocking.

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		default:
			fmt.Println("no value received")
		}
	}
}
