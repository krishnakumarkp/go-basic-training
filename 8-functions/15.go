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
	//in order to recover a panicking program successfully
	//a deferred function with recover call must be executed right after
	//the panic
	defer defferingFunction1()
	panic("panic from panickingFunction()")
	// a := []int{1, 2, 3}
	// fmt.Println(a[3])
	defer defferingFunction2()
	fmt.Println("panickingFunction() done")
}

func defferingFunction1() {
	fmt.Println("defferingFunction1() started")
	if r := recover(); r != nil {
		fmt.Println("Recovered the Program that was panicking with value", r)
	}
	fmt.Println("defferingFunction1() done")
}

func defferingFunction2() {
	fmt.Println("defferingFunction2() executed")
}

func defferingFromMain() {
	fmt.Println("defferingFromMain() executed")
}
