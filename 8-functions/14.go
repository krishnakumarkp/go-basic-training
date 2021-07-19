package main

import "fmt"

func main() {
	//panic defer execution sequence
	fmt.Println("main() started")
	defer defferingFromMain()
	panickingFunction()
	fmt.Println("main() done")
}

func panickingFunction() {
	fmt.Println("panickingFunction() started")
	defer defferingFunction1()
	panic("panic from panickingFunction()")
	defer defferingFunction2()
	fmt.Println("panickingFunction() done")
}

func defferingFunction1() {
	fmt.Println("defferingFunction1() executed")
}

func defferingFunction2() {
	fmt.Println("defferingFunction2() executed")
}

func defferingFromMain() {
	fmt.Println("defferingFromMain() executed")
}
