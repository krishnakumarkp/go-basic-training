package main

import "fmt"

func main() {
	//closures
	//Closures are anonymous functions that access
	//the variables defined outside the body of the function.
	num := 5
	func() {
		fmt.Println("num=", num)
	}()

	//example demonstrating a function with state using closure
	//new counter function returns the anonymous function
	//which refters to a vriable count and is assigned to a variable counter
	counter := newCounter()

	//invoke the anonymous function
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	//refer to dependency_closure_2.go to show use of closure
	// for dependency injection.

}

func Counter() int {
	count := 0
	count += 1
	return count
}

func newCounter() func() int {
	count := 0
	return func() int {
		//refers to variable count
		count += 1
		return count
	}
}
