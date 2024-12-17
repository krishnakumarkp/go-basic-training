package main

import "fmt"

// global variable
var globalVar = "I am global"

func main() {
	// Function-level scope
	var localVar = "I am local to main function"
	fmt.Println("Inside main function:")
	fmt.Println("Local variable:", localVar)
	fmt.Println("Global variable:", globalVar)

	// Accessing global variable inside function
	accessGlobal()

	// Block-level scope
	if true {
		var blockVar = "I am local to a block"
		fmt.Println("\nInside block:")
		fmt.Println("Block variable:", blockVar)
		// Accessing local variable from enclosing block
		fmt.Println("Local variable from main function:", localVar)
		// Accessing global variable from enclosing block
		fmt.Println("Global variable:", globalVar)
	}

	//fmt.Println("Block variable outside block:", blockVar) // Error: undefined variable

	// shadowing
	if true {
		localVar := "I am shadowing the outer localVar"
		fmt.Println("\nInside if scope after shadowing:")
		fmt.Println("Local variable:", localVar)
		fmt.Println("Global variable:", globalVar)
	}
	fmt.Println("\nInside main function after shadowing:")
	fmt.Println("Local variable:", localVar)
	fmt.Println("Global variable:", globalVar)

	// shadowing in function scope
	shadowFunction()
}

func accessGlobal() {
	// Function-level scope
	// Accessing global variable inside a function
	fmt.Println("\nInside accessGlobal function:")
	fmt.Println("Global variable:", globalVar)
}

func shadowFunction() {
	localVar := "I am shadowing the outer localVar in the function scope"
	fmt.Println("\nInside shadowFunction:")
	fmt.Println("Local variable:", localVar)
	fmt.Println("Global variable:", globalVar)
}

// Shadowing occurs in Go when a variable declared in a nested scope has the same name as
// a variable declared in an outer scope. In such cases, the variable in the inner scope shadows
// or hides the variable in the outer scope.
