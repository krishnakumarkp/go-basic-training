package main

import (
	"fmt"
	"krishnakumarkp/srp/solution/internal/employee"
	"krishnakumarkp/srp/solution/internal/payroll"
	"krishnakumarkp/srp/solution/internal/persistence"
	"krishnakumarkp/srp/solution/internal/reporting"
)

func main() {
	emp := employee.Employee{"Alice", 40, 25}

	payCalc := payroll.PayCalculator{}
	reporter := reporting.Reporter{}
	saver := persistence.Saver{}

	facade := employee.NewEmployeeFacade(payCalc, reporter, saver)

	fmt.Println("Salary:", facade.CalculatePay(emp))
	facade.ReportHours(emp)
	facade.SaveEmployee(emp)

}
