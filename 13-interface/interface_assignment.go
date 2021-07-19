package main

import (
	"fmt"
)

// define  interface SalaryCalculator which has function CalculateSalary() which returns int
type SalaryCalculator interface {
	CalculateSalary() int
}

//create two structs Permanent and Contract
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

//implement CalculateSalary for permanent type
//salary of permanent employee is the sum of basic pay and pf

//your code goes here

//implement CalculateSalary for contract type
//salary of contract employee is the basic pay alone

//your code goes here

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees by calling method CalculateSalary
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	//your code goes here
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}

//now add new employee type
// type Freelancer struct {
// 	empId       int
// 	ratePerHour int
// 	totalHours  int
// }

//implement CalculateSalary for Freelancer type
//salary of freelancer  employee is ratePerHour * totalHours

//create an employee of freelancer type and add him to emplyoees slice
