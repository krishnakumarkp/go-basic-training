package main

import (
	"fmt"
	"krishnakumarkp/srp/problem/employee"
)

func main() {

	emp := employee.Employee{"Alice", 40, 25}

	fmt.Println("Salary:", emp.CalculatePay())
	emp.ReportHours()
	emp.SaveEmployee()

}
