package main

import "fmt"

func main() {

	//var b func(int,int)string
	//assign a function to a variable
	var a func()

	//These kinds of functions are called anonymous functions since they do not have a name.
	a = func() {
		fmt.Println("Hello world first class function")
	}
	fmt.Printf("%T \n", a)
	a()

	//It is also possible to call an anonymous function without assigning it to a variable
	//go anonymous function also know as function litral
	sum := func(a, b, c int) int {
		return a + b + c
	}(3, 5, 7)

	fmt.Println("5+3+7 = ", sum)

	//user defined types
	type number int

	var num number
	num = 10
	fmt.Println(num)

	//user defined function types
	type addition func(int, int) int

	var add addition

	add = func(a int, b int) int {
		return a + b
	}

	s := add(5, 6)
	fmt.Println("Sum", s)
}
