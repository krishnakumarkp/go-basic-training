package main

import "fmt"

func main() {
	primes := []int{2, 3, 5}
	x := accessElement(primes, 3)
	//program recovers and rerurns default value of the return type
	fmt.Println(x)
}

func accessElement(a []int, index int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recovered but nothing to do!")
		}
	}()
	return a[index]
}
