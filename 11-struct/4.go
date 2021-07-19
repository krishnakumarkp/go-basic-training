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
	PersonalInfo  Person
	MonthlySalary []Salary
}

func main() {
	p1 := Person{
		name:     "Rohit",
		age:      45,
		location: "Gandhinagar",
		phone:    "94003372",
	}
	//Nested Struct Type
	e1 := Employee{
		PersonalInfo: p1,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Printf("e1 = %+v\n", e1)

	var e2 Employee

	p2 := Person{
		name:     "krishnakumar",
		age:      45,
		location: "Pune",
		phone:    "8007123343",
	}

	var janSal Salary
	janSal.Basic = 10000
	janSal.HRA = 5000
	janSal.TA = 4000

	var febSal Salary
	febSal.Basic = 10000
	febSal.HRA = 6000
	febSal.TA = 10000

	e2.PersonalInfo = p2
	//or
	e2.PersonalInfo.name = "krishnakumar"
	e2.PersonalInfo.age = 45
	e2.PersonalInfo.location = "Pune"
	e2.PersonalInfo.phone = "8007123343"

	e2.MonthlySalary = []Salary{janSal, febSal}
	//or
	e2.MonthlySalary = make([]Salary, 3)
	e2.MonthlySalary[0] = janSal
	e2.MonthlySalary[1] = febSal

	fmt.Printf("e2 = %+v\n", e2)

}
