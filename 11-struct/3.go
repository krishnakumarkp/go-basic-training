package main

import "fmt"

// define a type person
type Person struct {
	name            string
	age             int
	location, phone string
}

func main() {

	// Create Instance of struct type
	var p1 Person

	p1.name = "krishna"
	p1.age = 56
	p1.location = "Pune"
	p1.phone = "9008744556"

	fmt.Printf("p1 = %+v\n", p1)

	//Creating a Struct Instance Using a Struct Literal
	p2 := Person{
		name:     "Rohit",
		age:      45,
		location: "Gandhinagar",
		phone:    "94003372",
	}
	fmt.Printf("p2 = %+v\n", p2)

	//Struct Instantiation using new keyword
	//p3 is a pointer to an instance of person
	p3 := new(Person)
	fmt.Printf("p3 = %+v\n", *p3)
	p3.name = "Gaurav"
	p3.age = 22
	p3.location = "Pune"
	p3.phone = "90088765456"
	fmt.Printf("p3 = %+v\n", *p3)

	var p4 = Person{"Vishal", 33, "Hyderabad", "8087654878"} //cant skip any value
	fmt.Printf("p4 = %+v\n", p4)

	//Struct Instantiation Using Pointer Address Operator
	var p5 = &Person{"Siddarth", 22, "Pune", "7098708432"} //cant skip any value
	fmt.Printf("p5 = %+v\n", *p5)

}
