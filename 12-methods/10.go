package main

import "fmt"

type Person struct {
	name string
	age  int
}

//Method with Pointer receiver
func (p *Person) changeName(newName string) {
	p.name = newName
}

//Method with pointer receiver
func (p *Person) changeAge(newAge int) {
	p.age = newAge
}

type Employee struct {
	empId int
	Person
}

//Method with pointer receiver
func (p *Employee) changeEmpId(newId int) {
	p.empId = newId
}

func main() {

	var e Employee
	e.name = "Siddarth"
	e.age = 23
	e.empId = 4557
	e.changeAge(32)

	fmt.Printf("Employee age after change : %d \n", e.age)
}
