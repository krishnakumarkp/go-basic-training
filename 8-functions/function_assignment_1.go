package main

import "fmt"

// write function Calculator that returns a function to perform arithmetic operations
// as per the value of the parameter "operation"
func Calculator(operation string) func(float64, float64) float64 {

	//your code goes here

}
func main() {
	add := Calculator("add")
	subtract := Calculator("subtract")
	multiply := Calculator("multiply")
	divide := Calculator("divide")
	a, b := 10.0, 5.0
	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, add(a, b))
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, subtract(a, b))
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, multiply(a, b))
	fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, divide(a, b))
}
