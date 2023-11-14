package main

import "fmt"

type Person struct {
	name string
	age  int
}

//Method with value receiver
func (p Person) changeName(newName string) {
	p.name = newName
}

//Method with pointer receiver
func (p *Person) changeAge(newAge int) {
	p.age = newAge
}

func main() {
	p := Person{
		name: "Krishna",
		age:  50,
	}

	fmt.Printf("Person name before change: %s \n", p.name)
	p.changeName("Gaurav")
	fmt.Printf("Person name after change : %s \n", p.name)

	fmt.Printf("Person age before change: %d \n", p.age)
	//(&p).changeAge(60)
	p.changeAge(60)
	fmt.Printf("Person age after change : %d \n", p.age)

}

//create factory fucntions inside package to return structs to mimic constructors

// When to use pointer receiver and when to user value receiver
// Ponter receivers can be used when changes made to the receiver inside the method should
// be visible to the caller
// Pointer receivers can also be used in places where it's expensive to copy a data structure.
// In all other situations, value receivers can be used
