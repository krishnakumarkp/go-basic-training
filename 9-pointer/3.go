package main

import "fmt"

func main() {
	//passing pointer to a function
	number := 100
	changeValue(&number)

	var pnt *int
	pnt = &number
	changeValue(pnt)

	fmt.Printf("Pointer pnt of the type %T is having value %v \n", pnt, *pnt)

}

func changeValue(p *int) {
	*p = 2
}
