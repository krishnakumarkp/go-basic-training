package main

import "fmt"

func main() {
	//if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	num := 9

	// if, else if, else
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//if with a short statement
	//n := 10
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even \n", n)
	} else {
		fmt.Println(n)
	}

}
