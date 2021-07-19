package main

import "fmt"

type Rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {

	r1 := Rectangle{10, 20, "Green"}
	fmt.Println(r1)

	r2 := r1
	r2.color = "Pink"
	fmt.Println(r1)

	r3 := &r1
	(*r3).color = "Red"
	//We can access the fields of a struct pointer without dereferencing it first.
	//Go will take care of dereferencing a pointer under the hood.
	r3.color = "Red"
	fmt.Println(r1)

}
