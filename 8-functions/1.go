package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := add(1, 2)
	fmt.Println("1+2 = ", sum)

	product := multiply(3, 2)
	fmt.Println("3*2 = ", product)

	r1, r2, _ := threeRandom()
	fmt.Println("3 random numbers", r1, r2, r3)

	// //blank identifier
	// _, _, r4 := threeRandom()
	// fmt.Println("3rd random number", r4)

}

// returns two ints and returns their sum as int
func add(a int, b int) int {
	return a + b
}

// multiple consecutive prameter of the same type
func addAdd(a, b, c int) int {
	return a + b + c
}

// go functions allow to return multiple values
func threeRandom() (int, int, int) {
	x := rand.Intn(200)
	y := rand.Intn(200)
	z := rand.Intn(200)

	return x, y, z
}

//named return varibales

func addAdd(a, b, c int) (sum int) {
	sum = a + b + c
	return
}

// function named return varibales
func threeRandom() (x, y, z int) {
	x = rand.Intn(200)
	y = rand.Intn(200)
	z = rand.Intn(200)
	return
}
