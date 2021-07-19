package main

import "fmt"

func main() {
	//slice is also copy by value not reference but copy also points to underlying array
	s := []string{"a", "b", "c", "d", "e"}
	fmt.Println("slice s:", s)
	c := s
	fmt.Println("copy:", c)
	s[2] = "kk"
	fmt.Println("slice s:", s)
	fmt.Println("copy:", c)
	fmt.Println("len slice:", len(s))
	fmt.Println("len copy:", len(c))
	s = append(s, "z")
	fmt.Println("len slice:", len(s))
	fmt.Println("len copy:", len(c))
	fmt.Println("c:", c)
	fmt.Println("s:", s)

	//copy slice
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	slc2 := make([]int, len(slc1))

	copy(slc2, slc1)
	fmt.Println("\nSlice2:", slc2)
}
