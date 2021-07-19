package main

import "fmt"

func main() {
	//declare a pointer variable
	var pnt *int
	fmt.Printf("Pointer pnt of the type %T is having value %v \n", pnt, pnt)

	//lets's crate a variable of type int and make pa point to it
	number := 100
	pnt = &number
	fmt.Printf("Pointer pnt of the type %T is having value %v \n", pnt, pnt)

	//In the shorthand foramt
	// pnt1 := &number
	// fmt.Printf("Pointer pnt of the type %T id having value %v \n", pnt1, pnt1)

	//dereferencing a Pointer
	fmt.Printf("Data at %v is %v \n", pnt, *pnt)

	//changing the variable value using a pointer
	*pnt = 123
	fmt.Printf("Data at %v is %v \n", pnt, *pnt)
	fmt.Printf("Value of varuable number is %v \n", number)

	//Nil pointer dereference: runtime panic
	var myPointerVar *int
	fmt.Println(*myPointerVar)

	// The new function
	// GO provides the new buit-in function which allocates memory and returns a pointer
	// to that memory.
	// This function will allocate some memory , write a zero-value of the type at that memory
	// location and return a pointer to that memory location.

	pa := new(int)
	fmt.Printf("Pointer pa of the type %T is having value %v \n", pa, *pa)

}

//https://golang.org/doc/faq#virus
