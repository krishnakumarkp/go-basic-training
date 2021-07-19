package main

import "fmt"

func main() {
	// make a slice of int of length 3 capacity 5
	var a []int

	a = make([]int, 3, 5)

	//or
	var b = make([]int, 3, 5)
	fmt.Println("emptyint:", b)

	// make a slice of int of length 3 capacity 3
	c := make([]int, 3)
	fmt.Println("empstring:", c)

	//make a slice from existing array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	s[0] = 16
	fmt.Println(s)
	fmt.Println(primes)

}
