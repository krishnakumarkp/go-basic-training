package main

import "fmt"

//slice referencing to the slice of an array
func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	num := numbers[2:4]

	fmt.Printf("num=%v\n", num)
	fmt.Printf("length=%d\n", len(num))
	fmt.Printf("capacity=%d\n", cap(num))

	num[1] = 10

	fmt.Printf("num=%v\n", num)
	fmt.Printf("numbers=%v\n", numbers)

}
