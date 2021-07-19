package main

import "fmt"

func main() {
	primes := []int{2, 3, 5}
	x := accessElement(primes, 3)
	//program recovers and rerurns default value of the return type
	fmt.Println(x)
}

// Let us make accessElement return a fallback value, last element in the slice?

// If the deferred function is a function literal and the surrounding function
// has named result parameters that are in scope within the literal,
// the deferred function may access and modify the result parameters
// before they are returned

// func accessElement(a []int, index int) int {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Program recovered but nothing to do!")
// 		}
// 	}()
// 	return a[index]
// }

func accessElement(a []int, index int) (v int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recovered by returning the last element!")
			v = a[len(a)-1] //set v to the last element
		}
	}()
	v = a[index]
	return
}
