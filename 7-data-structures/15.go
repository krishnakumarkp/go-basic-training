package main

import "fmt"

func main() {
	//qUIZ: what is the output
	num := make([]int, 3, 5)

	num = append(num, 12)
	num = append(num, 15)

	fmt.Printf("num=%v\n", num)
}
