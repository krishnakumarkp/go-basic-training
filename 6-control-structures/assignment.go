package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Simple Calculator!")

	for {
		// Prompt the user to enter two numbers
		var num1, num2 float64
		fmt.Print("\nEnter the first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)

		// Display menu options
		fmt.Println("\nMenu:")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Print("\nEnter your choice (1-4): ")

		// Get user choice
		var choice int
		fmt.Scanln(&choice)

		// Perform the corresponding operation based on the user's choice
		var result float64
		var operator string
		switch choice {
		case 1:
			result = num1 + num2
			operator = "+"
		case 2:
			result = num1 - num2
			operator = "-"
		case 3:
			result = num1 * num2
			operator = "*"
		case 4:
			if num2 == 0 {
				fmt.Println("Error: Division by zero!")
				continue
			}
			result = num1 / num2
			operator = "/"
		default:
			fmt.Println("Invalid choice!")
			continue
		}

		// Display the result
		fmt.Printf("\nResult: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)

		// Ask the user if they want to perform another calculation
		var repeat string
		fmt.Print("\nDo you want to perform another calculation? (yes/no): ")
		fmt.Scanln(&repeat)
		if repeat != "yes" {
			break
		}
	}

	fmt.Println("\nThank you for using Simple Calculator!")
	os.Exit(0)
}

// Objective: To reinforce understanding of basic programming concepts such as variables, user input, loops, conditional statements, and switch case.
// Do not use functions for this assignment

// Implement a simple calculator program in Go that performs basic arithmetic operations (addition, subtraction, multiplication, and division) on two numbers provided by the user.
// The program should display welcome and goodbye messages as shown in the sample output.
// The program should follow the menu-based approach for selecting the operation and display the result in the specified format.

// Program sample output:

// Welcome to Simple Calculator!

// Enter the first number: 2
// Enter the second number: 3

// Menu:
// 1. Addition
// 2. Subtraction
// 3. Multiplication
// 4. Division

// Enter your choice (1-4): 1

// Result: 2 + 3 = 5   //format Result: <first_number> <operator> <second_number> = <result>

// Do you want to perform another calculation? (yes/no): no

// Thank you for using Simple Calculator!

// If the user types "yes" for "Do you want to perform another calculation?", the program should restart and prompt the user to enter the first number, the second number, and the operation choice. It should continue repeating this process until the user types "no".

// use this code to get user inputs

// 		// Prompt the user to enter two numbers
// 		var num1, num2 float64
// 		fmt.Print("\nEnter the first number: ")
// 		fmt.Scanln(&num1)
// 		fmt.Print("Enter the second number: ")
// 		fmt.Scanln(&num2)

// 		// Display menu options
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Addition")
// 		fmt.Println("2. Subtraction")
// 		fmt.Println("3. Multiplication")
// 		fmt.Println("4. Division")
// 		fmt.Print("\nEnter your choice (1-4): ")

// 		// Get user choice
// 		var choice int
// 		fmt.Scanln(&choice)
