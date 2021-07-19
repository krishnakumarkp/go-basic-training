package main

import "fmt"

func main() {
	// passing array pointer as function parameter is not idiomatic in go. You should use slice
	var numbers = [3]int{1, 2, 3}
	changeArrayValue(numbers[:])
	fmt.Printf("numbers = %v \n", numbers)
}

func changeArrayValue(s []int) {
	s[0] = 10
	s[1] = 11
	s[2] = 12
}
