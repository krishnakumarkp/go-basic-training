package main

import "fmt"

func main() {
	fmt.Println("Function Return Value:", fun1())
}

func fun1() int {
	i := 0
	defer func() {
		fmt.Println("1st", i)
		i++
	}()
	defer func() {
		fmt.Println("2nd", i)
		i++
	}()
	defer func() {
		fmt.Println("3rd", i)
		i++
	}()
	return i
}
