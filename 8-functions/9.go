package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer fmt.Println("four")
	fmt.Println("end")
}
