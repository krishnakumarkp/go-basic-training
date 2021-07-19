package main

import "fmt"

//Function fields
// in go function is a type, so struct fields can also be functions.
type Person struct {
	FirstName string
	LastName  string
	FullName  func(string, string) string
}

func main() {

	p := Person{
		FirstName: "Krishna",
		LastName:  "Kumar",
		FullName: func(firstName string, lastName string) string {
			return firstName + " " + lastName
		},
	}
	fmt.Println(p.FullName(p.FirstName, p.LastName))
}
