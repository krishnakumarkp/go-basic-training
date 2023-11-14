package main

import "fmt"

//Higher-order functions
//s a function which does at least one of the following
// 1) takes one or more functions as arguments
// 2) returns a function as its result

func main() {
	// Sample list of numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//passing function as arguments to other functions
	// Call AddNumbers with the isEven callback function.
	result := AddNumbers(numbers, isEven)
	// Print the result.
	fmt.Printf("Sum of even numbers: %d\n", result)

	//returning functions from other function
	// Create a function that multiplies by 2.
	double := createMultiplier(2)
	// Create a function that multiplies by 3.
	triple := createMultiplier(3)
	// Use the generated functions.
	fmt.Println("Double of 5:", double(5))   // Output: Double of 5: 10
	fmt.Println("Triple of 5:", triple(5))   // Output: Triple of 5: 15
	fmt.Println("Double of 10:", double(10)) // Output: Double of 10: 20
	fmt.Println("Triple of 10:", triple(10)) // Output: Triple of 10: 30

}

// Function to sum the even numbers in a list based on a callback function.
func AddNumbers(numbers []int, filter func(int) bool) int {
	sum := 0
	for _, num := range numbers {
		if filter(num) {
			sum += num
		}
	}
	return sum
}

// Example callback function to check if a number is even.
func isEven(num int) bool {
	return num%2 == 0
}

// callback function to check if a number is odd
func isOdd(number int) bool {
	return number%2 != 0
}

// createMultiplier returns a function that multiplies its argument by the given factor.
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
