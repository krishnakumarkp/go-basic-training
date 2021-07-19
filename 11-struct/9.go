package main

import "fmt"

//Declaration of a struct type
type Person struct {
	name            string
	age             int
	location, phone string
}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	PersonalInfo Person
	Salary
}

func main() {

	p1 := Person{
		name:     "Rohit",
		age:      45,
		location: "Gandhinagar",
		phone:    "94003372",
	}
	var e1 Employee

	e1.PersonalInfo = p1
	//all the nested struct fields are automatically available on parent struct.
	// This is called field promotion.
	e1.Basic = 15000.00
	e1.HRA = 200.00
	e1.TA = 100

	fmt.Printf("e1 = %+v\n", e1)

}
