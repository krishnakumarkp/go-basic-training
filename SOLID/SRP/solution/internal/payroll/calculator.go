package payroll

import "krishnakumarkp/srp/solution/internal/employee"

type PayCalculator struct{}

func (p PayCalculator) Calculate(emp employee.Employee) float64 {
	return emp.Hours * emp.Rate
}
