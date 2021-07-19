package main

import "fmt"

func main() {
	//passing slice by reference
	slice := []string{"a", "a", "a"}

	//changeSlice(slice)

	changeSlice2(&slice)

	fmt.Println("slice from main", slice)
}

func changeSlice(s []string) {
	s[0] = "b"
	s[1] = "b"
	s[2] = "b"
	s = append(s, "c")
	fmt.Println("slice from changeSlice", s)
}

func changeSlice2(s *[]string) {
	(*s)[0] = "b"
	(*s)[1] = "b"
	(*s)[2] = "b"
	*s = append(*s, "c")
	fmt.Println("slice from changeSlice", *s)
}

// Therefore make sure to keep in mind that you can pass a slice by value
// if you want to modify merely the values of the elements, not their number or position,
// otherwise weird bugs will arise from time to time
