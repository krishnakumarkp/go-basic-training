package main

import "fmt"

func main() {
	fmt.Println("main() started")
	//when a panic occurs, no further statements are executed
	foo()

	fmt.Println("main() done")
}

func foo() {
	fmt.Println("foo() started")
	panic("panic from foo()")
	fmt.Println("foo() done")
}
