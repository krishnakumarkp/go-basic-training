package main

import "fmt"

// variadic function,
// A variadic function can accept variable number of parameters.
func main() {
	s1 := addAll(1, 2, 3)
	s2 := addAll(1, 2, 3, 4)
	s3 := addAll(1, 2, 3, 4, 5)
	fmt.Println(s1, s2, s3)

	//passing a slice to variadic function

	numbers := []int{1, 2, 3, 4, 5}
	total := addAll(numbers...)
	fmt.Println("The sum is:", total)
}

// A variadic function accepts an infinite number of arguments
func addAll(nums ...int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}
