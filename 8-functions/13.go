package main

import "fmt"

func main() {
	//a deferred function is a function that will be executed just after the surrounding
	//function hits returned statement or panics
	fmt.Println("main() started")
	panickingFunction()
	fmt.Println("main() done")
}

func panickingFunction() {
	defer defferingFunction()
	fmt.Println("panickingFunction() started")
	panic("panic from panickingFunction()")
	fmt.Println("panickingFunction() done")
}

func defferingFunction() {
	fmt.Println("defferingFunction() executed")
}
