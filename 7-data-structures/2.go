package main

import "fmt"

func main() {
	// 	Here we create an array a that will hold exactly 5 ints.
	// 	The type of elements and length are both part of the array’s type.
	// 	By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the array[index] = value syntax,
	// and get a value with array[index].
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//The builtin function  len returns the length of an array.
	fmt.Println("len:", len(a))

	//Use this syntax to declare and initialize an array in one line.
	//Go array literal
	b := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println("dcl:", b)

	//type of elements and length are both part of the array’s type.
	//cannot use b (type [6]int) as type [5]int in assignment
	a = b
}
