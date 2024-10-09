package main

import "fmt"

// slice referencing to the slice of an array
func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	numSlice := numbers[2:4]

	fmt.Printf("numSlice=%v\n", numSlice)
	fmt.Printf("length=%d\n", len(numSlice))
	fmt.Printf("capacity=%d\n", cap(numSlice))

	numSlice[1] = 10

	fmt.Printf("numSlice=%v\n", numSlice)
	fmt.Printf("numbers=%v\n", numbers)

}
