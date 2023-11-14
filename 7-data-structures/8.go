package main

import "fmt"

func main() {
	//what happens when a slice exceeds capacity
	//show differnece between length and capacity.

	//make a slice of strings of length 0 capacity 5
	var a []int = make([]int, 0, 5)
	fmt.Println("emptyslice:", a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(a), len(a), a)
	}
}
