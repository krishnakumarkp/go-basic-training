package main

import (
	"fmt"
)

func main() {
	//the defer statement defers the execution of a function
	//untill the surrounding function returns
	// It schedules a function call before the function returns.
	// Always executes regardless of execution path the functions follow
	// “defer” function call can be placed anywhere in the function.
	// We can have multiple Defer function calls
	// Can be used to release function resources
	defer fmt.Println("sky")
	defer fmt.Println("sky1")

	primes := []int{2, 3, 5}
	x := accessElement(primes, 3)
	fmt.Println(x)

	fmt.Println("falcon")
}

func accessElement(a []int, index int) int {
	return a[index]
}
