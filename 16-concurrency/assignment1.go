package main

import (
	"fmt"
	"sync"
)

// write function sum that will take in a slice of int, a channel and a reference to wait group
// and calculate the sum of all numbers in the slice and write the result to channel
func sum(numbers []int, resultChan chan int, wg *sync.WaitGroup) {

}

func main() {
	// we have a slice of numbers, we need to find the sum of all the numbers in the slice using concurrency
	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	}

	resultChan := make(chan int)
	var wg sync.WaitGroup

	// write you code here so that the numbers in the numbers slice above is grouped to groups of 10 numbers
	// passed in to the function sum as a goroutine. so if there are 100 numbers in slice you will call function
	// sum as goroutine 10 times in a loop.

	// Wait for all goroutines to finish

	total := 0
	//Collect results from resultChan and calculate total sum, if you have created 10 go routines,
	//there will be 10 results in the channel, adding all the 10 results will give the total sum of the slice numbers

	fmt.Printf("Sum: %d\n", total)
	// this should give the output Sum: 5050
}
