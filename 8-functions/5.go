package main

import "fmt"

func main() {
	//Higher-order functions
	//s a function which does at least one of the following
	// 1) takes one or more functions as arguments
	// 2) returns a function as its result

	//passing function as arguments to other functions
	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	//returning functions from other function
	add := getAdder()
	fmt.Println("add :", add(60, 7))
}

//passing functions as arguments to other functions
// can use USER defined type addition
func simple(a func(int, int) int) {
	fmt.Println("simple: ", a(60, 7))
}

//returning function form other function
func getAdder() func(int, int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}
