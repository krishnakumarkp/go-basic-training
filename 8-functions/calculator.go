package main

import "fmt"

// Calculator is a function that returns a function to perform arithmetic operations.
func Calculator(operation string) func(float64, float64) float64 {
	var opFunc func(float64, float64) float64
	switch operation {
	case "add":
		opFunc = func(a, b float64) float64 {
			return a + b
		}
	case "subtract":
		opFunc = func(a, b float64) float64 {
			return a - b
		}
	case "multiply":
		opFunc = func(a, b float64) float64 {
			return a * b
		}
	case "divide":
		opFunc = func(a, b float64) float64 {
			if b != 0 {
				return a / b
			} else {
				return 0
			}
		}
	default:
		opFunc = func(a, b float64) float64 {
			return 0
		}
	}
	return opFunc
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
