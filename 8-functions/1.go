package main

import (
	"fmt"
	"math/rand"
	"time"
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

//returns two ints and returns their sum as int
func add(a int, b int) int {
	return a + b
}

//multiple consecutive prameter of the same type
func addAdd(a, b, c int) int {
	return a + b + c
}

//function named return varibales
func multiply(a, b int) int {
	m := a * b
	return m
}

//go functions allow to return multiple values
func threeRandom() (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(200)
	y := rand.Intn(200)
	z := rand.Intn(200)

	return x, y, z
}

func threeRandom() (x, y, z int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(200)
	y = rand.Intn(200)
	z = rand.Intn(200)
	return
}
