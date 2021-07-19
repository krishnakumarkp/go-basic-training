package main

import "fmt"

func main() {
	//go defer argument evaluation
	fmt.Println("start")
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End")
}
