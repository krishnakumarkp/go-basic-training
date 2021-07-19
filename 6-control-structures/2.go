package main

import "fmt"

func main() {
	//var sum int = 0

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
