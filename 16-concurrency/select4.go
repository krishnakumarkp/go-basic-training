package main

import (
	"fmt"
	"time"
)

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
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
		//After waits for the duration to elapse and then sends the current time on the returned channel.
	case <-time.After(1 * time.Second):
		fmt.Println("Program timed out")
	}
}

//If you really want to poll multiple times, make the timer outside of the for loop,
// timeout := time.After(5 * time.Second)
// for {
//     select {
//     case <-timeout:
//         fmt.Println("There's no more time to this. Exiting!")
//         return
//     default:
//         fmt.Println("still waiting")
//     }
//     time.Sleep(pollInt)
// }
