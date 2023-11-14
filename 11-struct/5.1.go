package main

import (
	"fmt"
	"reflect"
)

// type Person struct {
// 	Name          string `salary_slip:"Employee Name" `
// 	MonthlySalary int    `salary_slip:"Salary"`
// }

type Person struct {
	Name          string `salary_slip:"Full Name" profile:"Resource Name"`
	MonthlySalary int    `salary_slip:"Salary" profile:"ctc"`
}

func PrintSalarySlip(p Person) {
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag1Value := fieldType.Tag.Get("salary_slip")
		fmt.Printf("%s : %v\n", tag1Value, field.Interface())
	}
}

func PrintProfile(p Person) {
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag1Value := fieldType.Tag.Get("profile")
		fmt.Printf("%s : %v\n", tag1Value, field.Interface())
	}
}

func main() {
	person := Person{
		Name:          "Durvesh Patil ",
		MonthlySalary: 50000,
	}

	fmt.Println("===salary slip===")
	PrintSalarySlip(person)
	fmt.Println("")

	fmt.Println("===Emp  Profile===")
	PrintProfile(person)
}
