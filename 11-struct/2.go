package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
	}

	var person1 person
	person1.name = "krishna"
	person1.age = 56

	fmt.Printf("krishna = %+v\n", person1)

}
