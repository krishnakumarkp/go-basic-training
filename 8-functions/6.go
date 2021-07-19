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

	//new counter function returns the anonymous function
	//which refters to a vriable count and is assigned to a variable counter
	counter := newCounter()

	//invoke the anonymous function
	for i := 0; i <= 9; i++ {
		fmt.Println(counter())
	}

}

func newCounter() func() int {
	count := 0
	return func() int {
		//refers to variable count
		count += 1
		return count
	}
}
