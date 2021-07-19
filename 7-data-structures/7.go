package main

import "fmt"

func main() {
	//make a slice of strings of length 3 capacity 3
	s := make([]string, 3)
	fmt.Println("empstring:", s)

	//set and get just like with array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//length of slice
	fmt.Println("len:", len(s))
	//capacity of slice
	fmt.Println("cap:", cap(s))

	//index out range eorror
	//s[3] = "d"

	//append elements in to a slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//Slice literal declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//append one slice to another slice

	//names := make([]string, 2)

	names := []string{"John", "Paul"}
	morenames := []string{"George", "Ringo", "Pete"}
	names = append(names, morenames...)
	fmt.Println("names:", names)

}
