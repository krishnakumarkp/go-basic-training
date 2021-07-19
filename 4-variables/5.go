package main

import "fmt"
//variables excercise
func main() {
	var (
		name, location string
		age            int
	)

	name = "yourname"
	//age := 50
	location = "pune"

	fmt.Println(location)
	fmt.Println(name, age, location)
	fmt.Printf("My name is %s, i am %d years old, i live in %s \n", name, age, location)
}

// Try variable grouping
// Try Type inference
// Try short variable declaration
// Try assigning age to name  name = age
